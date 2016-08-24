package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080")
	//resp, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	status := resp.Status
	fmt.Printf("Status %s\n", status)
	fmt.Printf("%s\n", robots)
}
