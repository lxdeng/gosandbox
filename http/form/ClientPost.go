package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	hc := http.Client{}

	apiUrl := "http://localhost:9090"

	form := url.Values{}
	form.Add("key1", "value1")
	form.Add("key2", "value2")

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(form.Encode()))

	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := hc.Do(req)

	fmt.Printf("resp=%v\n", resp)
	fmt.Printf("err=%v\n", err)
}
