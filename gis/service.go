package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var nominatimSearchEP string = "https://nominatim.openstreetmap.org/search?q="
var searchQueryParams string = "&format=json&countrycodes=ir&dedupe=1&addressdetails=1&bounded=1"
var nominatimReverseEP string = "https://nominatim.openstreetmap.org/reverse?"

type GeoService interface {
	GeoCoding(address string, lon float64, lat float64) ([]geocodingResponseElement, error)
	ReverseGeoCoding(long, lat float64) (string, DetailedAddress, error)
	Autocomplete(partial_address string, lon float64, lat float64) ([]nominatimGeoResponse, error)
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
	BoundingBox     [4]string `json:"boundingbox"`
}

//FIXME
func getBBoxFromLocation(lon float64, lat float64) bBox {
	return bBox{51.247101, 35.614884, 51.564331, 35.775486}
}

func makeGeoSearchQuery(address string, limit uint8, tehranBBox bBox) (out string) {
	URL := nominatimSearchEP + address + "&limit=" +
		strconv.FormatUint(uint64(limit), 10) + "&viewbox=" +
		strconv.FormatFloat(tehranBBox.Left, 'f', 6, 64) + "," +
		strconv.FormatFloat(tehranBBox.Bottom, 'f', 6, 64) + "," +
		strconv.FormatFloat(tehranBBox.Right, 'f', 6, 64) + "," +
		strconv.FormatFloat(tehranBBox.Top, 'f', 6, 64) + searchQueryParams
	fmt.Println(URL)
	return URL
}

//FIXME
func buildAddress(nmResponse nominatimGeoResponse) string {
	var address string
	if nmResponse.DetailedAddress.Suburb != "" {
		address += nmResponse.DetailedAddress.Suburb + ","
	}
	if nmResponse.DetailedAddress.Neighbourhood != "" {
		address += nmResponse.DetailedAddress.Neighbourhood + ","
	}
	if nmResponse.DetailedAddress.HouseNumber != "" {
		address += nmResponse.DetailedAddress.HouseNumber
	}

	return address
}

func (geoService) GeoCoding(searchTerm string, lon float64, lat float64) ([]geocodingResponseElement, error) {
	fmt.Println("geocode")
	defaultBbox := getBBoxFromLocation(lon, lat)
	if lon == 0 && lat == 0 {
		defaultBbox = bBox{51.247101, 35.614884, 51.564331, 35.775486}
	}
	URL := makeGeoSearchQuery(searchTerm, 10, defaultBbox)

	fmt.Println(URL)
	resp, err := http.Get(URL)
	handleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	var nmResponse []nominatimGeoResponse
	var result []geocodingResponseElement
	if err := json.Unmarshal(body, &nmResponse); err != nil {
		fmt.Println(err.Error())
		return []geocodingResponseElement{}, err
	}

	fmt.Println(nmResponse)
	for i := 0; i < len(nmResponse); i++ {
		fmt.Println("begin")
		result[i].Long, _ = strconv.ParseFloat(nmResponse[i].Lon, 32)
		result[i].Lat, _ = strconv.ParseFloat(nmResponse[i].Lat, 32)

		result[i].bBox.Bottom, _ = strconv.ParseFloat(nmResponse[i].BoundingBox[0], 64)
		result[i].bBox.Left, _ = strconv.ParseFloat(nmResponse[i].BoundingBox[1], 64)
		result[i].bBox.Right, _ = strconv.ParseFloat(nmResponse[i].BoundingBox[2], 64)
		result[i].bBox.Top, _ = strconv.ParseFloat(nmResponse[i].BoundingBox[3], 64)

		result[i].Address = buildAddress(nmResponse[i])

		result[i].DetailedAddress = nmResponse[i].DetailedAddress
		fmt.Println("end")
	}

	return result, nil
}

func PopulateReverseGeodingResponse(response nominatimReverseResponse) reversegeocodingResponse {
	var result reversegeocodingResponse
	return result
}

func getNominatimReverseGeoResponse(latStr, lonStr string, zoomLevel uint8) (string, error) {
	URL := "https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=" +
		latStr + "&lon=" + lonStr + "&zoom=" + strconv.FormatUint(uint64(zoomLevel), 8) + "14"
	resp, err := http.Get(URL)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	handleError(err)

	var nmResponse nominatimReverseResponse
	if err := json.Unmarshal(body, &nmResponse); err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return nmResponse.Name, nil
}

func makeAddress(lonStr, latStr string) string {
	var result string

	suburbName, _ := getNominatimReverseGeoResponse(latStr, lonStr, 14)
	majorStreetName, _ := getNominatimReverseGeoResponse(latStr, lonStr, 16)
	majorAndMinorStreetName, _ := getNominatimReverseGeoResponse(latStr, lonStr, 17)
	buildingName, _ := getNominatimReverseGeoResponse(latStr, lonStr, 18)

	if suburbName != "" {
		result += suburbName
	}

	if majorStreetName != "" && majorAndMinorStreetName != "" {
		if majorStreetName != majorAndMinorStreetName {
			result += majorStreetName + majorAndMinorStreetName
		} else {
			result += majorStreetName
		}
	}

	if buildingName != "" {
		result += buildingName
	}

	fmt.Println("result:" + result)
	return result
}

func (geoService) ReverseGeoCoding(lon, lat float64) (string, DetailedAddress, error) {
	fmt.Println("reverse")
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

	var bb bBox
	bb.Bottom, _ = strconv.ParseFloat(nmResponse.BoundingBox[0], 64)
	bb.Left, _ = strconv.ParseFloat(nmResponse.BoundingBox[1], 64)
	bb.Right, _ = strconv.ParseFloat(nmResponse.BoundingBox[2], 64)
	bb.Top, _ = strconv.ParseFloat(nmResponse.BoundingBox[3], 64)

	address := makeAddress(nmResponse.Lon, nmResponse.Lat)

	return address, nmResponse.DetailedAddress, nil
}

// FIXME-need a ranking function
func rankCompletions(nmResponse []nominatimGeoResponse) []nominatimGeoResponse {
	return nmResponse
}

func (geoService) Autocomplete(partialAddress string, lon float64, lat float64) ([]nominatimGeoResponse, error) {
	fmt.Println("autocomp")
	defaultBbox := getBBoxFromLocation(lon, lat)
	if lon == 0 && lat == 0 {
		defaultBbox = bBox{51.247101, 35.614884, 51.564331, 35.775486}
	}
	URL := makeGeoSearchQuery(partialAddress, 10, defaultBbox)

	fmt.Println(URL)
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
	nmResponse = rankCompletions(nmResponse)
	return nmResponse, nil
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Fatal Error ", err.Error())
		os.Exit(1)
	}
}
