package main

import (
	"log"
)

func main() {
	a := App{}
	a.Initialize("product.db")
	log.Println("port 8080")
	a.Run(":8080")
	log.Println("exit")
}
