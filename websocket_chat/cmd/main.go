package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	addr   = ":2023"
	banner = ` 
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |

   ChatRoom，start on%s`
)

func main() {
	fmt.Printf(banner+"\n", addr)
	server.RegisterHandle()
	log.Fatal(http.ListenAndServe(addr, nil))
}
