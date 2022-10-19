package traefik_header_transform

import (
	"context"
	"net/http"
)

type Config struct {}

// KuzzleAuth a plugin to use Kuzzle as authentication provider for Basic Auth Traefik middleware.
type HeaderTransform struct {
	next   http.Handler
	name   string
}

func CreateConfig() *Config {
	return &Config {
		// ...
	}
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HeaderTransform{
		next:   next,
		name:   name,
	}, nil
}

func (ht *HeaderTransform) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var origin string = rw.Header().Get("Origin")
	var referrer string = rw.Header().Get("Referrer")

	if origin == "null" {
		rw.Header().Set("Origin", referrer)
	}

	ht.next.ServeHTTP(rw, req)
}