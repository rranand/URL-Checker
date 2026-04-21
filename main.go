package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	url        string
	statusCode int
	err        error
	duration   time.Duration
}

func checkHealth(url string) Result {
	client := http.Client{
		Timeout: TimeoutDuration,
	}

	res := Result{
		url: url,
	}

	start := time.Now()
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		res.err = err
	} else {
		defer resp.Body.Close()
		res.statusCode = resp.StatusCode
	}

	res.duration = duration

	return res
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.cloudflare.com",
		"https://api.github.com",
		"https://httpbin.org/get",

		"http://github.com",
		"http://google.com",
		"http://httpbin.org/redirect/1",
		"http://httpbin.org/redirect/3",
		"http://httpbin.org/redirect-to?url=https://www.google.com",

		"https://httpbin.org/status/400",
		"https://httpbin.org/status/401",
		"https://httpbin.org/status/403",
		"https://httpbin.org/status/404",

		"https://httpbin.org/status/500",
		"https://httpbin.org/status/502",
		"https://httpbin.org/status/503",

		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/3",
		"https://deelay.me/5000/https://www.google.com",

		"https://expired.badssl.com/",
		"https://self-signed.badssl.com/",
		"https://wrong.host.badssl.com/",

		"http://localhost:9999",
		"http://invalid.url.test",
	}

	for i := range urls {
		res := checkHealth(urls[i])

		if res.err != nil {
			fmt.Println(res.err)
		}

	}
}
