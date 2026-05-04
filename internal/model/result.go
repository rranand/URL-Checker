// Package model contains models which contains metadata of URL.
package model

import (
	"time"
)

// ResultModel contains metadata of URL health.
type ResultModel struct {
	URL        string
	StatusCode int
	Err        error
	Duration   time.Duration
	WorkerID   int
}

// Status return status of URL's health as string
func (r *ResultModel) Status() string {
	status := ""

	if r.Err != nil {
		status = "DOWN"
	} else {

		if r.StatusCode >= 200 && r.StatusCode < 300 {
			status = "HEALTHY"
		} else if r.StatusCode >= 300 && r.StatusCode < 400 {
			status = "REDIRECT"
		} else if r.StatusCode >= 400 && r.StatusCode < 500 {
			status = "CLIENT_ERROR"
		} else {
			status = "SERVER_ERROR"
		}
	}

	return status
}
