// Package for building map of sites from single domain
package sitemap

import (
	"gophercises/html/parsing/link"
	"log"
	"net/url"
	"strings"
)

// type representing whole domain
type Domain struct {
	Name  string
	Sites []link.Link
}

// cheks if given string is from the domain
func (d Domain) contains(s string) bool {

	if strings.HasPrefix(s, "/") {
		return true
	}

	url, err := url.Parse(s)
	if err != nil {
		log.Printf("during parsing url of %v occured error:%v", s, err)
		return false
	}

	domain := strings.TrimPrefix(url.Hostname(), "www.")

	return d.Name == domain
}
