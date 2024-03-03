package fetch

import (
	"net/http"
	"syscall/js"
)

// Client is an HTTP client.
type Client struct {
	// namespace - Objects that Fetch API belongs to. Default is Global
	namespace js.Value
	userAgent string
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
			userAgent: c.userAgent,
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

// WithUserAgent sets the user agent on requests.
func WithUserAgent(ua string) ClientOption {
	return func(c *Client) {
		c.userAgent = ua
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
