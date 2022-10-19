package traefik_header_transform

import (
	"context"
	"net/http"
)

type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
}

type HeaderTransform struct {
	next   http.Handler
	headers  map[string]string
	name   string
}

func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HeaderTransform {
		headers:  config.Headers,
		next:   next,
		name:   name,
	}, nil
}

func (ht *HeaderTransform) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Origin") == "null" {
		req.Header.Set("Origin", req.Header.Get("Referrer"))
	}

	ht.next.ServeHTTP(rw, req)
}