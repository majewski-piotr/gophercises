package main

import (
	"database/sql"
	"flag"
	"fmt"
	"gophercises/urlshort"
	"log"
	"net/http"
	"os"
)

func main() {
	//read flag for bonus 1
	yamlPath := flag.String("yaml", "redirects.yaml", "path of the yaml file with redirects")
	jsonPath := flag.String("json", "redirects.json", "path of the json file with redirects")
	flag.Parse()

	getConnectionPostgres()

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
	_ = jsonHandler

	//Bonus3
	postgresConnection, err := getConnectionPostgres()
	postgresHandler, err := urlshort.PostgresHandler(postgresConnection, mapHandler)

	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8082")
	http.ListenAndServe(":8082", postgresHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func getConnectionPostgres() (*sql.DB, error) {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "notBestPracticeToPutItHere"
	dbname := "redirects"
	psqlConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	result, err := sql.Open("pgx", psqlConnectionString)
	err = result.Ping()
	if err != nil {
		log.Fatalf("Error connecting to db : %s", err)
	}
	return result, err
}
