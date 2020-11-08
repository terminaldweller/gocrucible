package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func main() {
	var svc GeoService
	svc = geoService{}

	geocodingHandler := httptransport.NewServer(makeGeoCodingEndpoint(svc), decodeGeocodingRequest, encodeGeocodingResponse)
	reversegeocodingHandler := httptransport.NewServer(makeReversegeoCodingEndpoint(svc), decodeReverseGeocodingRequest, encodeReversegeocodingResponse)
	autocompleteHandler := httptransport.NewServer(makeAutocompleteEndpoint(svc), decodeAutocompleteRequest, encodeAutocompleteResponse)

	http.Handle("/geocode", geocodingHandler)
	http.Handle("/reverse", reversegeocodingHandler)
	http.Handle("/autocomp", autocompleteHandler)
	http.ListenAndServe(":8080", nil)
}
