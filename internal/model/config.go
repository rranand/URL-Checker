package model

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// Config Model contains few parameters that will be needed
// for configuration of server
type Config struct {
	TimeoutDuration         time.Duration `json:"timeout_duration"`
	MaxIdleConns            int           `json:"max_idle_conn"`
	MaxIdleConnsPerHost     int           `json:"max_idle_conn_per_host"`
	IdleConnTimeoutDuration time.Duration `json:"idle_conn_timeout_duration"`
	NoOfWorkers             int           `json:"no_of_workers"`
}

// GetHTTPTransport return *http.Transport from config struct
func (c *Config) GetHTTPTransport() *http.Transport {
	return &http.Transport{
		MaxIdleConns:        c.MaxIdleConns,
		MaxIdleConnsPerHost: c.MaxIdleConnsPerHost,
		IdleConnTimeout:     c.IdleConnTimeoutDuration,
	}
}

// DefaultConfig return default config struct
func DefaultConfig() Config {
	return Config{
		TimeoutDuration:         3 * time.Second,
		MaxIdleConns:            100,
		MaxIdleConnsPerHost:     10,
		IdleConnTimeoutDuration: 30 * time.Second,
		NoOfWorkers:             10,
	}
}

// ParseConfig return parsed config from filename
// If filename is blank, application will proceed with default configuration
func ParseConfig(filename string) (Config, error) {
	dConfig := DefaultConfig()

	if len(filename) == 0 {
		return dConfig, nil
	}

	var conf Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(data, &conf)

	if err == nil {
		// From JSON, time duration related parameters are in ns, so here ns to seconds conversion is done
		conf.TimeoutDuration *= time.Second
		conf.IdleConnTimeoutDuration *= time.Second
	}

	return conf, err
}
