package main

import (
	"fmt"
	"sync"

	"github.com/rranand/URL-Checker/internal/model"
	"github.com/rranand/URL-Checker/internal/worker"
)

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

	var wg sync.WaitGroup
	workerPoolSize := 10

	urlChan := make(chan string, workerPoolSize)
	resChan := make(chan model.ResultModel, workerPoolSize)

	worker.InitiatateWorkerPool(workerPoolSize, &wg, urlChan, resChan)

	for i := range urls {
		urlChan <- urls[i]
	}

	close(urlChan)

	for res := range resChan {
		fmt.Println(res)
	}

}
