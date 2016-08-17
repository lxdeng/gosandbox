package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	url := flag.String("url", "", "url to connect")
	thread := flag.Int("thread", 1, "number of threads")
	minutes := flag.Int("time", 1, "number of minutes to run")

	flag.Parse()

	fmt.Println("url=", *url)
	fmt.Println("thread=", *thread)
	fmt.Println("time=", *minutes)

	for i := 0; i < *thread; i++ {
		go httpGet(i, *url)
	}

	delayMinute(time.Duration(*minutes))
}

func httpGet(id int, url string) {
	for {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		robots, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		status := res.Status
		fmt.Printf("%d: %s, %d\n", id, status, len(robots))
	}
}

func delayMinute(n time.Duration) {
	time.Sleep(n * time.Minute)
}
