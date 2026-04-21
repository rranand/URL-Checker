package worker

import (
	"fmt"
	"sync"

	"github.com/rranand/URL-Checker/internal/model"
	urlchecker "github.com/rranand/URL-Checker/internal/urlChecker"
)

func InitiatateWorkerPool(workerPoolSize int, wg *sync.WaitGroup, urlChan <-chan string, resChan chan<- model.ResultModel) {
	for i := range workerPoolSize {
		wg.Add(1)
		go workerPool(i, wg, urlChan, resChan)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()
}

func workerPool(workerPoolId int, wg *sync.WaitGroup, urlChan <-chan string, resChan chan<- model.ResultModel) {
	defer func() {
		fmt.Println("worker pool closed, id:", workerPoolId)
		wg.Done()
	}()

	fmt.Println("worker pool initialized, id:", workerPoolId)

	for url := range urlChan {
		resChan <- urlchecker.URLChecker(url)
	}

}
