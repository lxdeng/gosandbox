package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	path := "/tmp/uhttpd.sock"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fmt.Println("Unix HTTP server: " + path)

	root := "."
	if len(os.Args) > 2 {
		root = os.Args[2]
	}

	server := http.Server{
		Handler: http.FileServer(http.Dir(root)),
	}

	unixListener, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
	server.Serve(unixListener)
}
