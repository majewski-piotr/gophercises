package adventure

import (
	"html/template"
	"net/http"
)

type Adventure struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type AdventureHandler struct {
	AM map[string]Adventure
	T  *template.Template
}

// Handles requests from html pages generated from
// adventures structs. Trims backslash from path to
// mach keys in a given map
func (ah AdventureHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	ah.T.Execute(rw, ah.AM[path])
}
