package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"io/ioutil"
	"net/http"
)

func makeGeoCodingEndpoint(svc GeoService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(geocodingRequest)
		long, lat, err := svc.GeoCoding(req.Address)
		if err != nil {
			return geocodingResponse{long, lat, err.Error()}, nil
		}
		return geocodingResponse{long, lat, ""}, nil
	}
}

func makeReversegeoCodingEndpoint(svc GeoService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(reverseGeocodingRequest)
		address, err := svc.ReverseGeoCoding(req.Long, req.Lat)
		if err != nil {
			return reversegeocodingResponse{address, err.Error()}, nil
		}
		return reversegeocodingResponse{address, ""}, nil
	}
}

type geocodingRequest struct {
	Address string `json:"address"`
}

type reverseGeocodingRequest struct {
	Long float32 `json:"long"`
	Lat  float32 `json:"lat"`
}

type geocodingResponse struct {
	Long float32 `json:"long"`
	Lat  float32 `json:"lat"`
	Err  string  `json:"err,omitempty"`
}

type reversegeocodingResponse struct {
	Address string `json:"address"`
	Err     string `json:"err,omitempty"`
}

func decodeGeocodingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request geocodingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeReverseGeocodingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request reverseGeocodingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeGeocodingResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeGeocingRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeReversegeocodingResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeReversegeocodingRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
