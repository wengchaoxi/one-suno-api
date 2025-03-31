package provider

import (
	"net/http"
	"time"
)

func NewHTTPClient(headers map[string]string) *http.Client {
	tr := &http.Transport{
		MaxIdleConns:        256,
		IdleConnTimeout:     90 * time.Second,
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 100,
	}

	return &http.Client{
		Transport: &headerTransport{
			Transport: tr,
			headers:   headers,
		},
		Timeout: 360 * time.Second,
	}
}

type headerTransport struct {
	headers map[string]string
	*http.Transport
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		req.Header.Add(k, v)
	}
	return t.Transport.RoundTrip(req)
}
