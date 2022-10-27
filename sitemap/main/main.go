package main

import (
	"fmt"
	"gophercises/html/parsing/link"
	"gophercises/sitemap"
	"sync"
)

func main() {
	d := sitemap.Domain{
		Name:  link.Link{Text: "", Url: "http://warmachina.com.pl/"},
		Sites: make([]link.Link, 0),
	}

	wg := new(sync.WaitGroup)

	wg.Add(1)
	d.Fill()
	fmt.Printf("%v", len(d.Sites))
}
