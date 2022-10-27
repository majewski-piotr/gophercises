// Package for building map of sites from single domain
package sitemap

import (
	"fmt"
	"gophercises/html/parsing/link"
	"log"
	"net/http"
	"strings"
	"sync"

	h "golang.org/x/net/html"
)

// type representing whole domain
type Domain struct {
	Name  link.Link
	Sites []link.Link
}

func (d *Domain) Fill() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go d.feedLinks(d.Name, wg)
	wg.Wait()
}

// todo: concurrent structure for links

// todo: 'recursive' concurrent method of downloading pages
// works only if their addres is from domain and not in structure yet
func (d *Domain) feedLinks(myLink link.Link, wg *sync.WaitGroup) {
	if myLink.Url == "" {
		fmt.Println("got an empty link")
		wg.Done()
		return
	}
	//todo: handle these errors
	rsp, err := http.Get(myLink.Url)
	if err != nil {
		log.Fatal(err)
	}
	doc, _ := h.Parse(rsp.Body)
	links := link.GetLinks(doc)
	for _, l := range links {
		fmt.Println("found", l.Url)
		if strings.HasPrefix(l.Url, "./") {
			l.Url = d.Name.Url + l.Url[2:]
			fmt.Println("now its", l.Url)
		}
		if d.belongs(l) && !d.contains(l) {
			fmt.Println("adding", l.Url)
			d.Sites = append(d.Sites, l)
			wg.Add(1)
			go d.feedLinks(l, wg)
		}
	}
	wg.Done()
}

// todo : this needs to be concurrent safe!!!
func (d Domain) contains(l link.Link) bool {
	for _, s := range d.Sites {
		if s.Url == l.Url {
			return true
		}
	}
	return false
}

// checks if given string is from the domain
func (d Domain) belongs(l link.Link) bool {

	fmt.Println(strings.HasPrefix(l.Url, d.Name.Url))
	return strings.HasPrefix(l.Url, d.Name.Url)
}
