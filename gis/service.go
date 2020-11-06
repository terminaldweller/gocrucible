package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var nominatimSearchEP string = "https://nominatim.openstreetmap.org/search"
var nominatimReverseEP string = "https://nominatim.openstreetmap.org/reverse"

type GeoService interface {
	GeoCoding(address string) (float32, float32, error)
	ReverseGeoCoding(long, lat float32) (string, error)
}

type geoService struct{}

func (geoService) GeoCoding(address string) (float32, float32, error) {
	searchQuery := nominatimSearchEP + address

	resp, err := http.Get(searchQuery)
	handleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	fmt.Println(string(body))

	var nomiReq geocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&nomiReq); err == nil {
		return 0, 0, err
	}
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

func handleError(err error) {
	if err != nil {
		fmt.Println("Fatal Error ", err.Error())
		os.Exit(1)
	}
}
