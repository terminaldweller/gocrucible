package main

import (
	"flag"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"strconv"
)

func main() {
	var svc GeoService
	svc = geoService{}

	var portFlag = flag.Int("port", 8181, "determines which port gis should listen on.")
	flag.Parse()

	geocodingHandler := httptransport.NewServer(makeGeoCodingEndpoint(svc), decodeGeocodingRequest, encodeGeocodingResponse)
	reversegeocodingHandler := httptransport.NewServer(makeReversegeoCodingEndpoint(svc), decodeReverseGeocodingRequest, encodeReversegeocodingResponse)
	autocompleteHandler := httptransport.NewServer(makeAutocompleteEndpoint(svc), decodeAutocompleteRequest, encodeAutocompleteResponse)

	http.Handle("/geocode", geocodingHandler)
	http.Handle("/reverse", reversegeocodingHandler)
	http.Handle("/autocomp", autocompleteHandler)

	serveAddress := ":" + strconv.Itoa(*portFlag)

	http.ListenAndServe(serveAddress, nil)
}
