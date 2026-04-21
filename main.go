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

	// Creating waitgroup to track all request completion, as we are creating
	// threads (workers) to monitor URL health, it may happen URL health is still
	// being checked and main thread exit from the program.
	var wg sync.WaitGroup

	// Defining number of workers which will montior the health of URLs.
	workerPoolSize := 10

	// Defining size of buffer, it is related with workerPoolSize.
	buffer := workerPoolSize * 2

	// Assume we have 25 URLs, 10 workers, 20 buffer for channels.

	// Case 1 - What if buffer size >= workerpoolsize
	// url channel will start accepting urls from list, but health check
	// will take little time to complete. So, until that point, url channel
	// will be blocked as there is no room for new urls until few health results
	// are generated. So, if some urls checks are done, health check will initate for next urls

	// Case 2 - What if buffer size < workerpoolsize
	// url channel will start accepting urls from list, but it will consume urls less than
	// workerPoolSize, so few workers will always be free. We are not consuming all worker pools.
	// Remaining worker pool will waste.

	urlChan := make(chan string, buffer)
	resChan := make(chan model.ResultModel, buffer)

	// Initiating worker pools
	worker.InitiatateWorkerPool(workerPoolSize, &wg, urlChan, resChan)

	// ----------------------------------------------------------------------------------------------

	// We have created a thread here which will receive result from
	// result channel from workers. Later, we started sending urls
	// to url channel which will be consumed by the workers. Once all urls are sent,
	// we are closing the channel. And later we are waiting from wg to check if all request are
	// processed or not. As health check will take time, if we don't use wg.Wait() main thread will
	// exit the program once url channel is closed.

	// wg.Wait() is not sufficent here, we need different waitgroup to track if result channel is closed or not.
	// We can only return from the program only if result channel is closed. But why, wg.Wait() should be sufficent as it will
	// wait until counter is zero means all health checks are done of urls, but there is a race condition between printing of remaining health result
	// and program exiting. Because after wg.Wait() still result channel is opened, still it is showing results of health.
	// So we can only make sure by checking if result are processed or not by checking result channel is closed or not.

	// go func() {
	// 	for res := range resChan {
	// 		fmt.Println(res)
	// 	}
	// }()

	// for _, url := range urls {
	// 	urlChan <- url
	// }

	// close(urlChan)

	// wg.Wait()

	// ----------------------------------------------------------------------------------------------

	// We have created a thread here which will send url to
	// url channel from urls list. Later, we are started listening
	// from result channel for the health reports.

	// In this case, we are only exiting from the program only if all results are processed, and url channel
	// will be closed once all urls are sent to url channel.

	go func() {
		for i := range urls {
			urlChan <- urls[i]
		}
		close(urlChan)
	}()

	for res := range resChan {
		fmt.Println(res.Status())
	}

}
