package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func echoServer(c net.Conn) {
	buf := make([]byte, 1024)

	for {
		count, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[:count]
		println("Server got:", string(data))
		_, err = c.Write(data)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

//var connections = make([]net.Conn, 0)

func main() {
	fmt.Println("server started:...")

	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("got signal")
		fmt.Println(sig)
		//closeConnections()

		fmt.Println("Closing listener")
		l.Close()
		os.Exit(1)
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		//connections = append(connections, conn)
		go echoServer(conn)
	}
}

/*
func closeConnections() {
	fmt.Printf("Closing connections: %d connections\n", len(connections))
	for _, conn := range connections {
		conn.Close()
	}
}
*/
