package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	readAll()
	copyToStdOut()
}

func readAll() {
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

func copyToStdOut() {
	log.Println("copyToStdOut:")
	url := "//localhost:8080/count"
	if !strings.HasPrefix(url, "http:") {
		url = "http:" + url
	}

	var resp *http.Response
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()
}
