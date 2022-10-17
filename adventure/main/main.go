package main

import (
	"adventure"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	jsonPtr := flag.String("json", "gopher.json", "path to sory file in json format")
	templatePtr := flag.String("template", "adventure.gohtml", "path to html template file")
	flag.Parse()

	tpl, tplErr := template.ParseFiles(*templatePtr)
	if tplErr != nil {
		panic(tplErr)
	}
	jsonBytes, jsonErr := os.ReadFile(*jsonPtr)
	if jsonErr != nil {
		panic(jsonErr)
	}
	var am map[string]adventure.Adventure
	json.Unmarshal(jsonBytes, &am)

	ah := adventure.AdventureHandler{
		AM: am,
		T:  tpl,
	}

	fmt.Println("Starting the server on :8082")
	http.ListenAndServe(":8082", ah)

}
