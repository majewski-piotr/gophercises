// Package main builds executable to parse links
package main

import (
	"fmt"
	"gophercises/html/parsing/link"
	"log"
	"os"

	h "golang.org/x/net/html"
)

func main() {
	f, e := os.Open("ex4.html")
	if e != nil {
		log.Fatal("Cannot open file")
	}

	doc, err := h.Parse(f)
	if err != nil {
		log.Fatal("Cannot open file")
	}

	for _, l := range link.GetLinks(doc) {
		fmt.Println(l)
	}

}
