package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://www.google.com",
		"https://www.google.com",
		"https://www.google.com",
		"https://www.google.com",
	}
	for _, url := range urls {
		doHttp(url)
	}
}

func doHttp(url string) {
	t := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("kog")
	}
	defer resp.Body.Close()
	fmt.Printf("<%s> - Static[%d] - Latency %d ms \n", url, resp.StatusCode, time.Since(t).Microseconds())
}
