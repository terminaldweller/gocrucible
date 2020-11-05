package main

import (
	"encoding/json"
	"github.com/awslabs/smithy-go/transport/http"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
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

func decodeGeocodingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request geocodingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
