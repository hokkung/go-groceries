package cnt

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/go-resty/resty/v2"
	"time"
)

var DefaultTimeout = 5 * time.Second

type Client struct {
	name string

	engine        *resty.Client
	pingOptions   *PingOptions
	apiKeyOptions *APIKeyOptions
}

// Name returns name of destination service
func (c Client) Name() string {
	return c.name
}

// Ping pings to destination service
func (c Client) Ping() error {
	ctx := context.Background()
	res, err := c.Req(ctx).Get(c.pingOptions.Endpoint)
	if err != nil {
		return err
	}

	if !res.IsSuccess() {
		return errors.New("ping response is not ok")
	}

	return nil
}

// Req creates request with api key header
func (c Client) Req(ctx context.Context) *resty.Request {
	if c.apiKeyOptions != nil {
		return c.engine.R().
			SetContext(ctx).
			SetHeader(c.apiKeyOptions.Key, c.apiKeyOptions.Value)
	}
	return c.engine.R()
}

// NewClient returns new instance
func NewClient(ops *ClientOptions) (*Client, error) {
	r := resty.New().
		SetTimeout(DefaultTimeout).
		SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		})

	if ops.BaseURL == "" {
		return nil, errors.New("invalid base url")
	}

	r.SetBaseURL(ops.BaseURL)

	if ops.Timeout != 0 {
		r.SetTimeout(ops.Timeout)
	}
	if ops.TLSConfig != nil {
		r.SetTLSClientConfig(ops.TLSConfig)
	}

	c := &Client{
		engine: r,
		name:   ops.Name,
	}

	if ops.APIKeyOptions != nil {
		c.apiKeyOptions = ops.APIKeyOptions
	}

	if ops.PingOptions != nil {
		c.pingOptions = ops.PingOptions
		err := c.Ping()
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// ProvideClient provides client for di
func ProvideClient(ops *ClientOptions) (*Client, error) {
	return NewClient(ops)
}
