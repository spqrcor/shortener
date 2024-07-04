package main

import "flag"

var flagRunAddr string
var flagBaseUrl string

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&flagBaseUrl, "b", "http://localhost:8080", "base url")
	flag.Parse()
}
