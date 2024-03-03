package fetch

import (
	"net/http"
	"syscall/js"
)

// transport is an implementation of http.RoundTripper
type transport struct {
	// namespace - Objects that Fetch API belongs to. Default is Global
	namespace js.Value
	redirect  RedirectMode
	header    http.Header
}

// RoundTrip replaces http.DefaultTransport.RoundTrip to use cloudflare fetch
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, values := range t.header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	return fetch(t.namespace, req, &RequestInit{
		Redirect: t.redirect,
	})
}
