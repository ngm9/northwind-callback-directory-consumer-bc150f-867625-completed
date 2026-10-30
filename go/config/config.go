package config

import "time"

type Config struct {
	UpstreamBaseURL string
	RequestTimeout  time.Duration
}

func Load() Config {
	return Config{
		UpstreamBaseURL: "https://jsonmock.hackerrank.com/api/countries",
		RequestTimeout:  5 * time.Second,
	}
}
