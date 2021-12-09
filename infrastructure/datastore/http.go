package datastore

import "net/http"

// HTTPClient is an interface made for the client of http package.
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// ClientWrapper Wraps the Client from http in order abstract implementations
//with HTTPClient interface.
type ClientWrapper struct {
	client  *http.Client
	gateway string
}

func NewHTTPClient(gateway string) *ClientWrapper {
	return &ClientWrapper{client: http.DefaultClient, gateway: gateway}
}

// Get makes http.Client get request.
func (w *ClientWrapper) Get(url string) (*http.Response, error) {
	return w.client.Get(w.gateway + url)
}
