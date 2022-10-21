package main

import (
	"fmt"

	"log"
	"os"

	hlp "gophercises/html-link-parser"

	"golang.org/x/net/html"
)

func main() {
	f, e := os.Open("ex4.html")
	if e != nil {
		log.Fatal("Cannot open file")
	}

	doc, err := html.Parse(f)
	if err != nil {
		log.Fatal("Cannot open file")
	}

	for _, l := range hlp.GetLinks(doc) {
		fmt.Println(l)
	}

}
