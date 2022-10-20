package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func main() {
	f, e := os.Open("ex2.html")
	if e != nil {
		log.Fatal("Cannot open file")
	}

	doc, err := html.Parse(f)
	if err != nil {
		log.Fatal("Cannot open file")
	}

	linkChan := make(chan link)
	doneChan := make(chan bool)

	var wg sync.WaitGroup

	go runConcurently(doc, linkChan, &wg, doneChan)

	for {
		select {
		case l := <-linkChan:
			fmt.Printf("%+v \n", l)
		case <-doneChan:
			os.Exit(0)
		}
	}

}

func runConcurently(node *html.Node, linkChan chan link, wg *sync.WaitGroup, doneCHan chan bool) {
	wg.Add(1)
	go searchForLink(node, linkChan, wg)
	wg.Wait()
	doneCHan <- true

}

type link struct {
	url  string
	text string
}

func searchForLink(node *html.Node, linkChan chan link, wg *sync.WaitGroup) {

	switch node.Type {
	case html.ErrorNode:
		wg.Done()
		return
	case html.ElementNode:
		if strings.EqualFold(node.Data, "a") {
			l := link{
				//todo make func to filter only this attr val
				url:  node.Attr[0].Val,
				text: strings.TrimSpace(node.FirstChild.Data),
			}
			linkChan <- l
		}
	}
	if node.FirstChild != nil {
		wg.Add(1)
		go searchForLink(node.FirstChild, linkChan, wg)
	}
	if node.NextSibling != nil {
		wg.Add(1)
		go searchForLink(node.NextSibling, linkChan, wg)
	}
	wg.Done()
}
