package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("%v\n", os.Args)

	http.HandleFunc("/api/metar/", handleRequest)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	icaoCode := strings.Replace(r.URL.Path, "/api/metar/", "", -1)
	match, err := regexp.MatchString("[A-Za-z]{4}", icaoCode)
	var body map[string]string
	if !match {
		body = map[string]string{
			"error": "Not a valid icao code",
		}
	} else if err != nil {
		body = map[string]string{
			"error": err.Error(),
		}
	}
	if len(body) > 0 {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			log.Printf("Error marshalling JSON %v", err)
			w.WriteHeader(500)
		}
		w.WriteHeader(400)
		w.Write(jsonBody)
		return
	}
	date := r.URL.Query().Get("date")
	params := RequestParams{icao: []string{icaoCode}}
	if date != "" {
		params.date = date
	}
	metars := FetchMetar(params)
	jsonEncoded, err := json.Marshal(metars)
	if err != nil {
		log.Printf("Error marschalling metar %v", err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Content-Length", fmt.Sprintf("%d", len(jsonEncoded)))
		w.WriteHeader(200)
		w.Write(jsonEncoded)
	}
}
