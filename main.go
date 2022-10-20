package traefik_header_transform

import (
	"context"
	"log"
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
		if req.Header.Get("Referrer") != "" {
			req.Header.Set("Origin", req.Header.Get("Referrer"))
		} else {
			req.Header.Set("Origin", req.Header.Get("Referer"))
		}

	}
	
	rw.Header().Add("Access-Control-Allow-Origin", "")
	rw.Header().Add("Access-Control-Allow-Credentials", "true")
	rw.Header().Add("Vary", "Origin")

	log.Print("Response", rw.Header())
	log.Print("Request", req.Header)

	ht.next.ServeHTTP(rw, req)
}