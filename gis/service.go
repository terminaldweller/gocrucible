package main

import (
	"encoding/json"
	"fmt"
	// expand "github.com/openvenues/gopostal/expand"
	parser "github.com/openvenues/gopostal/parser"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var nominatimSearchEP string = "https://nominatim.openstreetmap.org/?addressdetails=1&q="
var nomipostSearch string = "&format=josn&limit=1"
var nominatimReverseEP string = "https://nominatim.openstreetmap.org/reverse"

type GeoService interface {
	GeoCoding(address string) (float32, float32, error)
	ReverseGeoCoding(long, lat float32) (string, error)
	Autocomplete(partial_address string) (string, error)
}

type geoService struct{}

func (geoService) GeoCoding(address string) (float32, float32, error) {
	parsedAddress := parser.ParseAddress(address)
	var searchQuery string
	for key, value := range parsedAddress {
		fmt.Println(key, value)
		searchQuery += value.Value + "+"
	}
	searchQuery = strings.ReplaceAll(searchQuery, " ", "+")
	searchQuery += nomipostSearch
	URL := nominatimSearchEP + searchQuery
	URL = "https://nominatim.openstreetmap.org/?addressdetails=1&q=bakery+in+berlin+wedding&format=json&limit=1"
	fmt.Println(URL)

	resp, err := http.Get(URL)
	handleError(err)
	defer resp.Body.Close()
	// fmt.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	// fmt.Println(string(body))
	var result interface{}
	json.Unmarshal(body, &result)
	fmt.Println(result)
	// fmt.Println(result.lat)
	// fmt.Println(result.lon)

	var nomiReq geocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&nomiReq); err != nil {
		fmt.Println(err.Error())
		return 10, 10, err
	}
	// fmt.Println(nomiReq.Lat, nomiReq.Long)
	return nomiReq.Long, nomiReq.Lat, nil
}

func (geoService) ReverseGeoCoding(long, lat float32) (string, error) {
	searchQuery := nominatimReverseEP + strconv.FormatFloat(float64(long), 'f', 6, 64) + strconv.FormatFloat(float64(lat), 'f', 6, 64)

	resp, err := http.Get(searchQuery)
	handleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	fmt.Println(string(body))

	var nomiReq reversegeocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&nomiReq); err == nil {
		return "", err
	}
	return nomiReq.Address, nil
}

func (geoService) Autocomplete(partialAddress string) (string, error) {
	return partialAddress, nil
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Fatal Error ", err.Error())
		os.Exit(1)
	}
}
