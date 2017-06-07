package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//
// curl -X GET "http://localhost:9090/?abc=111&url_long=222"
//
// curl -X POST -H 'Content-Type:application/x-www-form-urlencoded' -F 'abc=111' -F 'url_long=222' http://localhost:9090
// The difference between curl POST and the ClientPost.go

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Fprintf(w, "<p>%v</p>", r.Form)
	fmt.Println("path=", r.URL.Path)
	fmt.Println("scheme=", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "\nHello astaxie!\n") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName) // set router
	port := 9090
	log.Println("Listen on " + strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
