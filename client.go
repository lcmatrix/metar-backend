package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gitlab.com/lcmatrix/metar"
)

// URL to weather API
const URL = "https://api.met.no/weatherapi/tafmetar/1.0/metar.txt"

// FetchMetar calls weather service API and returns an array of metar informations.
func FetchMetar(params RequestParams) []metar.Metar {
	var urlBuilder strings.Builder
	urlBuilder.WriteString(URL)
	urlBuilder.WriteString("?icao=")
	urlBuilder.WriteString(strings.Join(params.icao, ","))
	if len(params.date) != 0 {
		urlBuilder.WriteString("&date=")
		urlBuilder.WriteString(params.date)
	}
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

	metars := convertResponseToMetar(b)
	if len(params.date) == 0 {
		return metars[len(metars)-1:]
	}

	return metars
}

func convertResponseToMetar(b []byte) []metar.Metar {
	var metarReturn []metar.Metar
	var builder strings.Builder
	builder.Write(b)
	s := builder.String()
	arr := strings.Split(s, "\n")
	for i := 0; i < len(arr); i++ {
		m, err := metar.ParseMetar(arr[i])
		if err == nil {
			metarReturn = append(metarReturn, *m)
		} else {
			log.Printf("Error in converting metar: %v", err)
		}
	}
	return metarReturn
}

// RequestParams is a struct of possible parameters
type RequestParams struct {
	icao           []string
	date           string
	timezoneOffset string
}
