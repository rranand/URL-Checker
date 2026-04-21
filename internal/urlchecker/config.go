package urlchecker

import "time"

var (
	timeoutDuration     = 3 * time.Second
	maxIdleConns        = 100
	maxIdleConnsPerHost = 10
	idleConnTimeout     = 30 * time.Second
)
