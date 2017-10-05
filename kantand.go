package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
)

func main() {
	// options with their defaults
	options := struct{
		host string
		port uint
		wwwRoot string
	}{
		host: "",
		port: 8000,
		wwwRoot: ".",
	}

	// parse command line options
	flag.StringVar(&options.host, "host", options.host, "Hostname or IP to bind to (empty string for all IPs on this host)")
	flag.UintVar(&options.port, "p", options.port, "Port to bind to")
	flag.StringVar(&options.wwwRoot, "www", options.wwwRoot, "Directory to serve")
	showHelp := flag.Bool("h", false, "Display help text")

	flag.Parse()

	if(*showHelp) {
		flag.Usage()
		return
	}

	bindTo := fmt.Sprintf("%v:%d", options.host, options.port)

	// serve the selected directory
	fmt.Printf("Serving directory '%v' via HTTP on '%v'.\n", options.wwwRoot, bindTo)
	log.Fatal(http.ListenAndServe(bindTo, http.FileServer(http.Dir(options.wwwRoot))))
}