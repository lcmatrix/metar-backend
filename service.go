package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v\n", os.Args)
	params := RequestParams{}
	params.icao = []string{os.Args[1]}
	FetchMetar(params)
}
