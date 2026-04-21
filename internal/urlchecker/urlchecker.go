// Package urlchecker provides tools to montior health of an URL.
package urlchecker

import (
	"net/http"
	"time"

	"github.com/rranand/URL-Checker/internal/model"
)

// URLChecker accepts http client and URL.
// It generate GET http request to check health of given URL.
// This is a pretty straight forward, recording time taken to get the health check.
// Store error and status code received from response.
func URLChecker(client *http.Client, url string) model.ResultModel {

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
