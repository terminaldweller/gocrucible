package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

func makeGeoCodingEndpoint(svc GeoService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(geocodingRequest)
		long, lat, err := svc.GeoCoding(req.Address, req.Lon, req.Lat)
		if err != nil {
			return geocodingResponse{long, lat, err.Error()}, nil
		}
		return geocodingResponse{long, lat, ""}, nil
	}
}

func makeReversegeoCodingEndpoint(svc GeoService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(reverseGeocodingRequest)
		address, detailedAddress, bbox, err := svc.ReverseGeoCoding(req.Long, req.Lat)
		if err != nil {
			return reversegeocodingResponse{"", DetailedAddress{}, bBox{}, err.Error()}, nil
		}
		return reversegeocodingResponse{address, detailedAddress, bbox, ""}, nil
	}
}

func makeAutocompleteEndpoint(svc GeoService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(autocompleteRequest)
		geoResponse, err := svc.Autocomplete(req.PartialString, req.Lon, req.Lat)
		if err != nil {
			return autocompleteResponse{geoResponse, err.Error()}, nil
		}
		return autocompleteResponse{geoResponse, ""}, nil
	}
}

type DetailedAddress struct {
	City            string `json:"city"`
	CityDistrict    string `json:"city_district"`
	Construction    string `json:"construction"`
	Continent       string `json:"continent"`
	Country         string `json:"country"`
	CountryCode     string `json:"country_code"`
	HouseNumber     string `json:"house_number"`
	Neighbourhood   string `json:"neighbourhood"`
	PostCode        string `json:"postcode"`
	Public_Building string `json:"public_building"`
	State           string `json:"state"`
	Suburb          string `json:"suburb"`
}

type geocodingRequest struct {
	Address string  `json:"address"`
	Lon     float64 `json:"lon"`
	Lat     float64 `json:"lat"`
}

type autocompleteRequest struct {
	PartialString string  `json:"autocomp"`
	Lon           float64 `json:"lon"`
	Lat           float64 `json:"lat"`
}

type reverseGeocodingRequest struct {
	Long float64 `json:"lon"`
	Lat  float64 `json:"lat"`
}

type bBox struct {
	Left   float64
	Bottom float64
	Right  float64
	Top    float64
}

type geocodingResponse struct {
	Long float64 `json:"lon"`
	Lat  float64 `json:"lat"`
	Err  string  `json:"err,omitempty"`
}

type reversegeocodingResponse struct {
	DisplayName     string `json:"display_name"`
	DetailedAddress `json:"address"`
	bBox            `json:"bbox"`
	Err             string `json:"err,omitempty"`
}

type autocompleteResponse struct {
	Addresses []nominatimGeoResponse `json:"detailed_address"`
	Err       string                 `json:"err,omitempty"`
}

func decodeGeocodingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request geocodingRequest
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	for k, v := range params {
		switch k {
		case "address":
			request.Address = v[0]
		case "lon":
			request.Lon, _ = strconv.ParseFloat(v[0], 64)
		case "lat":
			request.Lat, _ = strconv.ParseFloat(v[0], 64)
		default:
		}
	}

	fmt.Println(request)
	return request, nil
}

func decodeReverseGeocodingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request reverseGeocodingRequest
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	for k, v := range params {
		switch k {
		case "lon":
			request.Long, _ = strconv.ParseFloat(v[0], 64)
		case "lat":
			request.Lat, _ = strconv.ParseFloat(v[0], 64)
		default:
		}
	}

	fmt.Println(request)
	return request, nil
}

func decodeAutocompleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request autocompleteRequest
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	for k, v := range params {
		switch k {
		case "partial_address":
			request.PartialString = v[0]
		case "lat":
			request.Lat, _ = strconv.ParseFloat(v[0], 64)
		case "lon":
			request.Lon, _ = strconv.ParseFloat(v[0], 64)
		default:
		}
	}

	fmt.Println(request)
	return request, nil
}

func encodeGeocodingResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
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

func encodeAutocompleteResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func encodeReversegeocodingResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
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
