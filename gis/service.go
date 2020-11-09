package main

import (
	"encoding/json"
	"fmt"
	// expand "github.com/openvenues/gopostal/expand"
	// parser "github.com/openvenues/gopostal/parser"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	// "strings"
)

var nominatimSearchEP string = "https://nominatim.openstreetmap.org/search?q="
var nomipostSearch string = "&format=josn&limit=1"
var nominatimReverseEP string = "https://nominatim.openstreetmap.org/reverse?"

type GeoService interface {
	GeoCoding(address string) (float64, float64, error)
	ReverseGeoCoding(long, lat float64) (string, DetailedAddress, error)
	Autocomplete(partial_address string) ([]nominatimGeoResponse, error)
}

type geoService struct{}

type nominatimGeoResponse struct {
	DetailedAddress `json:"address"`
	BoundingBox     [4]string `json:"boundingbox"`
	Class           string    `json:"class"`
	DisplayName     string    `json:"display_name"`
	Icon            string    `json:"icon"`
	Importance      float64   `json:"importance"`
	Lat             string    `json:"lat"`
	License         string    `json:"licence"`
	Lon             string    `json:"lon"`
	OsmID           float64   `json:"osm_id"`
	OsmType         string    `json:"osd_type"`
	PlaceID         float64   `json:"place_id"`
	Type            string    `json:"type"`
}

type nominatimReverseResponse struct {
	PlaceID         int64   `json:"place_id"`
	License         string  `json:"license"`
	OSMType         string  `json:"osm_type"`
	OSMID           int64   `json:"osm_id"`
	Lat             string  `json:"lat"`
	Lon             string  `json:"lon"`
	PlaceRank       int32   `json:"place_rank"`
	Category        string  `json:"category"`
	Type            string  `json:"type"`
	Importance      float64 `json:"importance"`
	AddressType     string  `json:"addresstype"`
	DisplayName     string  `json:"display_name"`
	Name            string  `json:"name"`
	DetailedAddress `json:"address"`
	BoundingBox     [4]string `json:"bounding_box"`
}

func makeGeoSearchQuery(address string, limit uint8) (out string) {
	URL := nominatimSearchEP + address + "&format=json&limit=" + strconv.FormatUint(uint64(limit), 10)
	fmt.Println(URL)
	return URL
}

func (geoService) GeoCoding(address string) (float64, float64, error) {
	URL := makeGeoSearchQuery(address, 1)

	resp, err := http.Get(URL)
	handleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	var nmResponse []nominatimGeoResponse
	if err := json.Unmarshal(body, &nmResponse); err != nil {
		fmt.Println(err.Error())
		return 0, 0, err
	}
	fmt.Println(nmResponse[0].Lat, nmResponse[0].Lon)
	lon, err := strconv.ParseFloat(nmResponse[0].Lon, 32)
	lat, err := strconv.ParseFloat(nmResponse[0].Lat, 32)
	return lon, lat, nil
}

func PopulateReverseGeodingResponse(response nominatimReverseResponse) reversegeocodingResponse {
	var result reversegeocodingResponse
	return result
}

func (geoService) ReverseGeoCoding(lon, lat float64) (string, DetailedAddress, error) {
	lonStr := strconv.FormatFloat(lon, 'f', -1, 64)
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	URL := "https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=" + latStr + "&lon=" + lonStr

	resp, err := http.Get(URL)
	handleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	var nmResponse nominatimReverseResponse
	if err := json.Unmarshal(body, &nmResponse); err != nil {
		fmt.Println(err.Error())
		return "", DetailedAddress{}, err
	}
	return nmResponse.DisplayName, nmResponse.DetailedAddress, nil
}

func (geoService) Autocomplete(partialAddress string) ([]nominatimGeoResponse, error) {
	URL := makeGeoSearchQuery(partialAddress, 10)

	resp, err := http.Get(URL)
	handleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	var nmResponse []nominatimGeoResponse
	if err := json.Unmarshal(body, &nmResponse); err != nil {
		fmt.Println(err.Error())
		return []nominatimGeoResponse{}, err
	}
	fmt.Println(nmResponse)
	return nmResponse, nil
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Fatal Error ", err.Error())
		os.Exit(1)
	}
}
