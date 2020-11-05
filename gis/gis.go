package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	// "strings"
)

var nominatimSearchEP string = "https://nominatim.openstreetmap.org/search"
var nominatimReverseEP string = "https://nominatim.openstreetmap.org/reverse"

type GeoService interface {
	GeoCoding(long, lat float32) (string, error)
	ReverseGeoCoding(address string) (float32, float32, error)
}

type geoService struct{}

// func (geoService) GeoCoding(long, lat float32) (string, error) {
// 	sendReverse()
// }

// func (geoService) ReverseGeoCoding(address string) (float32, float32, error) {
// 	sendSearch(address)
// }

func handleError(err error) {
	if err != nil {
		fmt.Println("Fatal Error ", err.Error())
		os.Exit(1)
	}
}

func sendSearch(searchQuery string) {
	resp, err := http.Get(searchQuery)
	handleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	fmt.Println(string(body))
}

func sendReverse(searchQuery string) {
	resp, err := http.Get(searchQuery)
	handleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	fmt.Println(string(body))
}

func main() {
	// sendSearch("17, Strada Pictor Alexandru Romano, Bukarest, Bucharest, Sector 2, Bucharest, 023964, Romania")
	sendSearch("https://nominatim.openstreetmap.org/search?q=135+pilkington+avenue,+birmingham&format=xml&polygon_geojson=1&addressdetails=1")
	// sendReverse("https://nominatim.openstreetmap.org/reverse?format=geojson&lat=44.50155&lon=11.33989")
	// fileserver := http.FileServer(http.Dir("/search"))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {

}

func revsearchHandler(w http.ResponseWriter, r *http.Request) {

}
