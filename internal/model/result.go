package model

import "time"

type ResultModel struct {
	URL        string
	StatusCode int
	Err        error
	Duration   time.Duration
}
