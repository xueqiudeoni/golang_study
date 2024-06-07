package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "serving:%s\n", r.URL.Path)
	fmt.Printf("served time for :%s\n ", r.Host)
}
func timeHandler(response http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "the current time is: "
	fmt.Fprintf(response, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(response, "<h1 align=\"center\">%s</h1>", t)
	fmt.Fprintf(response, "serving:%s\n", r.URL.Path)
	fmt.Printf("served time for :%s\n ", r.Host)
}
func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", myHandler)
	r.HandleFunc("/time", timeHandler)
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe(":1234", r)
	if err != nil {
		fmt.Println(err)
		return
	}
}
