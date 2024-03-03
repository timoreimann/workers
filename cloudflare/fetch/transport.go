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
	userAgent string
}

// RoundTrip replaces http.DefaultTransport.RoundTrip to use cloudflare fetch
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.userAgent != "" {
		req.Header.Add("User-Agent", t.userAgent)
	}

	return fetch(t.namespace, req, &RequestInit{
		Redirect: t.redirect,
	})
}
