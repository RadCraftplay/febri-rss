package controllers

import (
	"febri-rss/common"
	"net/http"
)

var FebriApiClient *http.Client

type withHeader struct {
	http.Header
	rt http.RoundTripper
}

func WithHeader(rt http.RoundTripper) withHeader {
	if rt == nil {
		rt = http.DefaultTransport
	}

	return withHeader{Header: make(http.Header), rt: rt}
}

func (h withHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(h.Header) == 0 {
		return h.rt.RoundTrip(req)
	}

	req = req.Clone(req.Context())
	for k, v := range h.Header {
		req.Header[k] = v
	}

	return h.rt.RoundTrip(req)
}

func SetupHttpClients(configuration common.FebriRssConfiguration) {
	FebriApiClient = http.DefaultClient
	rt := WithHeader(FebriApiClient.Transport)
	rt.Set("XAppKey", configuration.FebriApi.AppKey)
	rt.Set("XAppSecret", configuration.FebriApi.AppSecret)
	FebriApiClient.Transport = rt
}
