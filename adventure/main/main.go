package main

import (
	"adventure"
	"encoding/json"
	"flag"
	"fmt"
	ht "html/template"
	"log"
	"net/http"
	"os"
	tt "text/template"
)

func main() {
	typePtr := flag.String("type", "web", "type of application, cli or web")
	jsonPtr := flag.String("json", "gopher.json", "path to sory file in json format")
	templatePtr := flag.String("template", "adventure.gohtml", "path to template file. use html for web and text for cli")
	flag.Parse()

	jsonBytes, jsonErr := os.ReadFile(*jsonPtr)
	if jsonErr != nil {
		panic(jsonErr)
	}
	var am map[string]adventure.Adventure
	json.Unmarshal(jsonBytes, &am)

	switch *typePtr {
	case "web":
		runWeb(templatePtr, am)
	case "cli":
		runCli(templatePtr, am)
	default:
		log.Fatal("Wrong type of app:", *typePtr)
	}

}
func runCli(templatePtr *string, am map[string]adventure.Adventure) {
	tpl, tplErr := tt.ParseFiles(*templatePtr)
	if tplErr != nil {
		panic(tplErr)
	}

	ah := adventure.AdventureHandler{
		AM:       am,
		Template: tpl,
	}

	ah.RunCli()
}

func runWeb(templatePtr *string, am map[string]adventure.Adventure) {
	tpl, tplErr := ht.ParseFiles(*templatePtr)
	if tplErr != nil {
		panic(tplErr)
	}

	ah := adventure.AdventureHandler{
		AM:       am,
		Template: tpl,
	}

	fmt.Println("Starting the server on :8082")
	http.ListenAndServe(":8082", ah)
}
