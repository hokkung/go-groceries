package cnt

import (
	"crypto/tls"
	"time"
)

type ClientOptions struct {
	Name    string
	BaseURL string

	Timeout       time.Duration
	TLSConfig     *tls.Config
	PingOptions   *PingOptions
	APIKeyOptions *APIKeyOptions
}

type PingOptions struct {
	Endpoint string
}

type APIKeyOptions struct {
	Key   string
	Value string
}
