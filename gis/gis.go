package main

import (
	// "fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	// "log"
	"net/http"
)

// func sendSearch(searchQuery string) {
// 	resp, err := http.Get(searchQuery)
// 	handleError(err)
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	handleError(err)

// 	fmt.Println(string(body))
// }

// func sendReverse(searchQuery string) {
// 	resp, err := http.Get(searchQuery)
// 	handleError(err)
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	handleError(err)

// 	fmt.Println(string(body))
// }

func main() {
	// sendSearch("17, Strada Pictor Alexandru Romano, Bukarest, Bucharest, Sector 2, Bucharest, 023964, Romania")
	// sendReverse("https://nominatim.openstreetmap.org/reverse?format=geojson&lat=44.50155&lon=11.33989")
	// fileserver := http.FileServer(http.Dir("/search"))

	// sendSearch("https://nominatim.openstreetmap.org/search?q=135+pilkington+avenue,+birmingham&format=xml&polygon_geojson=1&addressdetails=1")

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
