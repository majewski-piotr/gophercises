package urlshort

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		destination, present := pathsToUrls[path]
		if present {
			http.Redirect(w, r, destination, http.StatusSeeOther)
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	redirectData := make([]RedirectPair, 0)
	err := yaml.Unmarshal(yml, &redirectData)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	return handleRedirectData(redirectData, fallback), err
}

// Handler for json data, json has to be a list of objects
// with parameters Path and Url.
// Redirects calls according to json, if path is not found then the
// fallback http.Handler will be called instead
func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var redirectData []RedirectPair
	err := json.Unmarshal(jsonData, &redirectData)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	return handleRedirectData(redirectData, fallback), err

}

type RedirectPair struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// helps transfer other input types to map handler
func handleRedirectData(redirectData []RedirectPair, fallback http.Handler) http.HandlerFunc {
	urlMap := make(map[string]string, len(redirectData))
	for _, pair := range redirectData {
		urlMap[pair.Path] = pair.Url
	}

	return MapHandler(urlMap, fallback)
}
