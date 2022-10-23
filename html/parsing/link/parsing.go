// Package link implemets tools to parse links from html pages
package link

import (
	"gophercises/utils/slices"
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
func GetLinks(root *html.Node) []Link {
	links := []Link{}
	if root == nil {
		return links
	}

	if isLink(root) {
		l := Link{
			Url:  root.Attr[0].Val,
			Text: getNestedText(root.FirstChild),
		}
		links = append(links, l)
	} else {
		links = slices.ConcatCopyPreAllocate(links, GetLinks(root.FirstChild))
	}

	if root.NextSibling != nil {
		links = slices.ConcatCopyPreAllocate(links, GetLinks(root.NextSibling))
	}
	return links

}

func isLink(n *html.Node) bool {
	return n != nil && n.Type == html.ElementNode && n.Data == "a"
}

// returns concatenated text from current tree
// level and successor nodes, recursively
// add one space between
func getNestedText(n *html.Node) string {
	var sb string
	if n == nil {
		return ""
	}

	if n.FirstChild != nil {
		sb += getNestedText(n.FirstChild)
	}

	if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
		sb += strings.TrimSpace(n.Data)
	}

	if n.NextSibling != nil {
		sb += getNestedText(n.NextSibling)
	}

	return sb
}
