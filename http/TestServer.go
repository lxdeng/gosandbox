package main

import (
	"fmt"
	"net/http"
	"log"
)

type MyHandler struct {
	name string
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%s name=%s\n", r.URL.Path, h.name)
}

func main() {

	a := [...]string {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	fmt.Println(a)
	
	http.HandleFunc("/foo", handler)

	h1 := new(MyHandler)
	h1.name = "h1"
	http.Handle("/bar", h1)

	//http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%s counter=%d\n", r.URL.Path, 1)
}