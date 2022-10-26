// Package implementig 'choose your own adventure' game type
package cyoa

import (
	"fmt"
	"io"
	"net/http"
)

// Represents single scene
type Adventure struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Represents options from scene
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// Represents map of scenes forming a full story
// Holds reference to template used to present
// Adventures to players
type AdventureHandler struct {
	AM       map[string]Adventure
	Template Executable
}

// interface to match both http and text template
type Executable interface {
	Execute(wr io.Writer, data any) error
}

// Handles requests from html pages generated from
// adventures structs. Trims backslash from path to
// mach keys in a given map
func (ah AdventureHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	ah.Template.Execute(rw, ah.AM[path])
}

// Runs adventures through cli, outputs through os.Strout,
// reads input from the console via fmt.Scanln
func (ah AdventureHandler) RunCli(in io.Reader, out io.Writer) {
	scene := "intro"
	for {
		a := ah.AM[scene]
		ah.Template.Execute(out, a)

		if len(a.Options) == 0 {
			return
		}

		optionNumber := -1
		fmt.Fscanln(in, &optionNumber)
		if optionNumber >= 0 && optionNumber < len(a.Options) {
			scene = a.Options[optionNumber].Arc
		} else {
			fmt.Fprintf(out, "Option %v is invalid, repeating scene\n", optionNumber)
		}
	}
}
