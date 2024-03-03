package fetch

import (
	"net/http"
	"syscall/js"
)

// Client is an HTTP client.
type Client struct {
	// namespace - Objects that Fetch API belongs to. Default is Global
	namespace js.Value
	header    http.Header
}

// applyOptions applies client options.
func (c *Client) applyOptions(opts []ClientOption) {
	for _, opt := range opts {
		opt(c)
	}
}

// HTTPClient returns *http.Client.
func (c *Client) HTTPClient(redirect RedirectMode) *http.Client {
	return &http.Client{
		Transport: &transport{
			namespace: c.namespace,
			redirect:  redirect,
			header:    c.header,
		},
	}
}

// ClientOption is a type that represents an optional function.
type ClientOption func(*Client)

// WithBinding changes the objects that Fetch API belongs to.
// This is useful for service bindings, mTLS, etc.
func WithBinding(bind js.Value) ClientOption {
	return func(c *Client) {
		c.namespace = bind
	}
}

// WithHeader sets one or more custom headers.
func WithHeader(header http.Header) ClientOption {
	return func(c *Client) {
		c.header = header
	}
}

// NewClient returns new Client
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		namespace: js.Global(),
	}
	c.applyOptions(opts)

	return c
}
