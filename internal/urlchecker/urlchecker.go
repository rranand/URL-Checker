package urlchecker

import (
	"net/http"
	"time"

	"github.com/rranand/URL-Checker/internal/model"
)

func URLChecker(url string) model.ResultModel {
	client := http.Client{
		Timeout: timeoutDuration,
		Transport: &http.Transport{
			MaxIdleConns:        maxIdleConns,
			MaxIdleConnsPerHost: maxIdleConnsPerHost,
			IdleConnTimeout:     idleConnTimeout,
		},
	}

	res := model.ResultModel{
		URL: url,
	}

	start := time.Now()
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		res.Err = err
	} else {
		defer resp.Body.Close()
		res.StatusCode = resp.StatusCode
	}

	res.Duration = duration

	return res
}
