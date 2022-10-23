// Package link implemets tools to parse links from html pages
package link

import (
	"gophercises/copy/slice"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Url  string
	Text string
}

// Traverses html.Node recursively and returns slice of links
// Links are discovereds by node.Data == "a"
// and all nested text inside them goes to Link.Text field
func GetLinks(n *html.Node) []Link {
	links := []Link{}
	if n == nil {
		return links
	}

	for n = n.FirstChild; n != nil; n = n.NextSibling {
		if isLink(n) {
			l := Link{
				Url:  n.Attr[0].Val,
				Text: getNestedText(n),
			}
			links = append(links, l)
		} else {
			links = slice.ConcatCopyPreAllocate(links, GetLinks(n))
		}
	}
	return links
}

func isLink(n *html.Node) bool {
	return n != nil && n.Type == html.ElementNode && n.Data == "a"
}

// returns concatenated text from childrens of
// a given node, recursively, adds one space between
func getNestedText(n *html.Node) string {
	var sb string
	if n == nil {
		return ""
	}

	for n = n.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
			sb += strings.TrimSpace(n.Data)
		}
		sb += getNestedText(n)
	}
	return sb
}
