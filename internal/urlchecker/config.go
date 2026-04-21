package urlchecker

import "time"

var (
	TimeoutDuration     = 3 * time.Second
	MaxIdleConns        = 100
	MaxIdleConnsPerHost = 10
	IdleConnTimeout     = 30 * time.Second
)
