package main

import (
	"github.com/go-kit/kit/endpoint"
)

func makeGeoCodingProxyEndpoint(svc GeoService) endpoint.Endpoint {
}

type geocodingRequest struct {
	address string `json:"address"`
}

type reverseGeocodingRequest struct {
	long float32 `json:"long"`
	lat  float32 `json:"lat"`
}

type geocodingResponse struct {
	long float32 `json:"long"`
	lat  float32 `json:"lat"`
	err  string  `json:"err, omitempty"`
}

type reversegeocodingResponse struct {
	address string `json:"address"`
	err     string `json:"err, omitempty"`
}
