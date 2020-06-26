package main

// Metar format with single fields
type Metar struct {
	time        string
	wind        string
	visibilty   string
	events      string
	clouds      string
	temperature string
	qnh         string
}
