package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// URL to weather API
const URL = "https://api.met.no/weatherapi/tafmetar/1.0/metar.txt"

// FetchMetar calls weather service API and returns an array of metar informations.
func FetchMetar(params RequestParams) []Metar {
	var urlBuilder strings.Builder
	urlBuilder.WriteString(URL)
	urlBuilder.WriteString("?icao=")
	urlBuilder.WriteString(strings.Join(params.icao, ","))
	res, err := http.Get(urlBuilder.String())
	if err != nil {
		log.Println(err)
	} else if !strings.HasPrefix(res.Status, "2") {
		log.Printf("Bad response code: %v\n", res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response: %v\n", err)
	}
	fmt.Printf("%s\n", b)
	var metar []Metar
	return metar
}

// RequestParams is a struct of possible parameters
type RequestParams struct {
	icao           []string
	date           time.Time
	timezoneOffset string
}
