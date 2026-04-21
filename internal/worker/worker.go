package worker

import (
	"net/http"
	"sync"

	"github.com/rranand/URL-Checker/internal/model"
	"github.com/rranand/URL-Checker/internal/urlchecker"
)

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

func workerPool(client *http.Client, wg *sync.WaitGroup, urlChan <-chan string, resChan chan<- model.ResultModel) {
	defer func() {
		wg.Done()
	}()

	for url := range urlChan {
		resChan <- urlchecker.URLChecker(client, url)
	}

}
