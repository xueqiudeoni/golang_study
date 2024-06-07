package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

const maxMessageSize = 512

var (
	newline = []byte("\n")
	space   = []byte(" ")
	HUB     *Hub
)
var upgrader = websocket.Upgrader{WriteBufferSize: 1024, ReadBufferSize: 1024}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
	uid  string
}
type Hub struct {
	clients     map[*Client]bool
	broadcast   chan []byte
	register    chan *Client
	unregister  chan *Client
	userClients map[string]*Client // 记录uid和client的对应关系
	sync.RWMutex
}

func newClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
}
func NewHub() *Hub {
	return &Hub{
		clients:     make(map[*Client]bool),
		broadcast:   make(chan []byte),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		userClients: make(map[string]*Client, 128),
		RWMutex:     sync.RWMutex{},
	}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client] = true
			h.Unlock()
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				h.Lock()
				delete(h.clients, client)
				delete(h.userClients, client.uid)
				h.Unlock()
				close(client.send)
			}
			//case message := <-h.broadcast:
			//	for client := range h.clients {
			//		select {
			//		case client.send <- message:
			//		default:
			//			close(client.send)
			//			delete(h.clients, client)
			//		}
			//	}
		}
	}
}
func (c *Client) readPump() {
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Time{}) //永不超时
	for {
		_, p, e := c.conn.ReadMessage()
		if e != nil {
			fmt.Println(e)
			break
		}
		//message := bytes.TrimSpace(bytes.Replace(p, newline, space, -1))
		//c.hub.broadcast <- message
		var data map[string]interface{}
		err := json.Unmarshal(p, &data)
		if err != nil {
			break
		}
		if uid, ok := data["uid"]; ok {
			c.uid = uid.(string)
			c.hub.Lock()
			c.hub.userClients[c.uid] = c
			c.hub.Unlock()
		}
	}
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
}

// write 把消息发送给websocket客户端
func (c *Client) write() {
	defer c.conn.Close()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(time.Second))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, e := c.conn.NextWriter(websocket.TextMessage)
			if e != nil {
				return
			}
			w.Write(message)
			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := newClient(hub, conn)
	client.hub.register <- client
	go client.write()
	go client.readPump()
}
func Send(hub *Hub, w http.ResponseWriter, r *http.Request) {

	uid := r.FormValue("uid")
	if uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hub.Lock()
	client, ok := hub.userClients[uid]
	hub.Unlock()
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	message := r.FormValue("message")
	client.send <- []byte(message)
}
