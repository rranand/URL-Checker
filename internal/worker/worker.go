// Package worker provide functions to create worker pools.
package worker

import (
	"net/http"
	"sync"

	"github.com/rranand/URL-Checker/internal/model"
	"github.com/rranand/URL-Checker/internal/urlchecker"
)

// InitiatateWorkerPool create light weight workers.
// It accepts workerPoolSize, waitgroup, sender url channel, receiver result channel.
// Here are initializing client as same client config will used to process request for different urls.
// If we create multiple clients for every urls, it will take little time everytime just to do redudant task
// that is initializing same client every time.
// Based on workerPoolSize, number of workers will created. Channels are restricted to send and recieve only mode.
// At last, we are waiting until wait group counter reset to zero which represents all health checks of given urls
// are completed, we are closing the result channel as no more data will given to the channel and in main program, channel
// will process remaining data in buffer and exit from the loop once all result are processed.
func InitiatateWorkerPool(workerPoolSize int, wg *sync.WaitGroup, urlChan <-chan string, resChan chan<- model.ResultModel) {
	client := http.Client{
		Timeout: urlchecker.TimeoutDuration,
		Transport: &http.Transport{
			MaxIdleConns:        urlchecker.MaxIdleConns,
			MaxIdleConnsPerHost: urlchecker.MaxIdleConnsPerHost,
			IdleConnTimeout:     urlchecker.IdleConnTimeout,
		},
	}

	for range workerPoolSize {
		wg.Add(1)
		go workerPool(&client, wg, urlChan, resChan)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()
}

// workerPool is one of the worker from worker pool.
// It receive url to be checked from urlChan and after getting health
// of URL, it return health data in result channel.
// This worker pool process the request, in simple words, it is taking one url
// at a time and get result from the url checker and add that result in result channel.
// It will listen until url channel is closed, once closed wait group will me marked as done.
func workerPool(client *http.Client, wg *sync.WaitGroup, urlChan <-chan string, resChan chan<- model.ResultModel) {
	defer func() {
		wg.Done()
	}()

	for url := range urlChan {
		resChan <- urlchecker.URLChecker(client, url)
	}

}
