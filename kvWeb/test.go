package main

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type myElement struct {
	Name    string
	Surname string
	ID      string
}

var DATA = make(map[string]myElement)
var DATAFILE = "./tmp/datafile.gob"

func save() error {
	fmt.Println("saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	encoder := gob.NewEncoder(f)
	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("encode data failed:", err)
		return err
	}
	return nil
}
func load() error {
	fmt.Println("loading:", DATAFILE)
	f, _ := os.Open(DATAFILE)
	defer f.Close()
	decoder := gob.NewDecoder(f)
	decoder.Decode(&DATA)
	return nil
}
func ADD(k string, element myElement) bool {
	if k == "" {
		return false
	}
	if LOOKUP(k) == nil {
		DATA[k] = element
		return true
	}
	return false
}
func LOOKUP(k string) *myElement {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	}
	return nil
}
func DELETE(k string) bool {
	if LOOKUP(k) != nil {
		delete(DATA, k)
		return true
	}
	return false
}
func CHANGE(k string, element myElement) bool {
	DATA[k] = element
	return true
}
func PRINT() {
	for s, element := range DATA {
		fmt.Printf("key:%s,value:%v\n", s, element)
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving", r.Host, "for", r.URL.Path)
	tp := template.Must(template.ParseGlob("home.gohtml"))
	tp.ExecuteTemplate(w, "home.gohtml", nil)
}
func listAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listing the contents of the KV store")
	fmt.Fprintf(w, "<a href=\"/\",style=\"margin-right:20px;\">Home sweet home!</a>")
	fmt.Fprintf(w, "<a href=\"/\",style=\"margin-right:20px;\">List all elments!</a>")
	fmt.Fprintf(w, "<a href=\"/\",style=\"margin-right:20px;\">Change an elment!</a>")
	fmt.Fprintf(w, "<a href=\"/\",style=\"margin-right:20px;\">Insert new elment!</a>")
	fmt.Fprintf(w, "<h1>The contens of the KV store are:</h1>")
	fmt.Fprintf(w, "<ul>")
	for s, element := range DATA {
		fmt.Fprintf(w, "<li>")
		fmt.Fprintf(w, "<strong>%s</strong>with value:%v\n", s, element)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</ul>")
}
func changeElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("change an elment of the kv store")
	tmpl := template.Must(template.ParseFiles("update.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	key := r.FormValue("key")
	n := myElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		ID:      r.FormValue("id"),
	}
	if !CHANGE(key, n) {
		fmt.Println("update opration failed")
	} else {
		err := save()
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, struct {
			success bool
		}{true})
	}
}
func insertElment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("insert element to the kv store")
	tmpl := template.Must(template.ParseFiles("insert.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	key := r.FormValue("key")
	n := myElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		ID:      r.FormValue("id"),
	}
	if !ADD(key, n) {
		fmt.Println("insert opration failed")
	} else {
		err := save()
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, struct {
			success bool
		}{true})
	}
}
func main() {
	err := load()
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", homePage)
	http.HandleFunc("/change", changeElement)
	http.HandleFunc("/insert", insertElment)
	http.HandleFunc("/list", listAll)
	err = http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
