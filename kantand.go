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
		bindTo string
		wwwRoot string
	}{
		bindTo: ":8000",
		wwwRoot: ".",
	}

	// parse command line options
	flag.StringVar(&options.bindTo, "bind", options.bindTo, "Host and port to bind to")
	flag.StringVar(&options.wwwRoot, "www", options.wwwRoot, "Directory to serve")
	showHelp := flag.Bool("h", false, "Display help text")

	flag.Parse()

	if(*showHelp) {
		flag.Usage()
		return
	}

	// serve the selected directory
	fmt.Printf("Serving directory '%v' via HTTP on '%v'.\n", options.wwwRoot, options.bindTo)
	log.Fatal(http.ListenAndServe(options.bindTo, http.FileServer(http.Dir(options.wwwRoot))))
}