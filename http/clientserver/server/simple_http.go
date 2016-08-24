package main

import (
	"fmt"
	"log"
	"net/http"
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

	fmt.Println("ListenAndServe: 8080")
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
