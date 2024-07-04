package main

import "flag"

var flagRunAddr string
var flagBaseURL string

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&flagBaseURL, "b", "http://localhost:8080", "base url")
	flag.Parse()
}
