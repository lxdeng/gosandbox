package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {

	// same response for all requests
	// curl http://localhost:8080/x/y/z
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
      <html>
        <head>
          <title>Chat</title>
        </head>
        <body>
          Let's chat!
        </body>
      </html>
    `))
	})

	http.HandleFunc("/count", handler)

	fmt.Println("ListenAndServe: 8080")
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var mu sync.Mutex
var count int64

func handler(w http.ResponseWriter, r *http.Request) {
	var c int64

	mu.Lock()
	count++
	c = count
	mu.Unlock()

	fmt.Fprintf(w, "URL.Path=%q counter=%d\n", r.URL.Path, c)
}
