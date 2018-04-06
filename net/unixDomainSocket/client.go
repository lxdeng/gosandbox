package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		count, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("Client got:", string(buf[:count]))
	}
}

func main() {
	fmt.Println("client started:...")
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	fmt.Println("connected...")

	go reader(c)
	for i := 0; i < 5; i++ {
		fmt.Println("client said: hi")
		_, err := c.Write([]byte("hi"))
		if err != nil {
			log.Fatal("write error:", err)
			break
		}
		time.Sleep(1e9)
	}
}
