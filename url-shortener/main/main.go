package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"url-shortener/urlshort"
)

func main() {
	//read flag for bonus 1
	yamlPath := flag.String("yaml", "redirects.yaml", "path of the yaml file with redirects")
	jsonPath := flag.String("json", "redirects.json", "path of the json file with redirects")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//read file for bonus 1
	yamlBytes, yamlErr := os.ReadFile(*yamlPath)
	if yamlErr != nil {
		panic(yamlErr)
	}

	//read file for bonus w
	jsonBytes, jsonErr := os.ReadFile(*jsonPath)
	if jsonErr != nil {
		panic(jsonErr)
	}

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := urlshort.YAMLHandler(yamlBytes, mapHandler)

	//Bonus2
	jsonHandler, err := urlshort.JSONHandler(jsonBytes, yamlHandler)

	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8082")
	http.ListenAndServe(":8082", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
