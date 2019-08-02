package urlshort

import (
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if newURL, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, newURL, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]RedirectObject, error) {
	var redirectArray []RedirectObject

	err := yaml.Unmarshal(yml, &redirectArray)
	if err != nil {
		return nil, err
	}
	return redirectArray, nil
}

func buildMap(yaml []RedirectObject) map[string]string {
	ret := make(map[string]string)
	for _, obj := range yaml {
		ret[obj.Path] = obj.URL
	}
	return ret
}

type RedirectObject struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLFileHandler will parse the provided YAML file and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
func YAMLFileHandler(filename string, fallback http.Handler) (http.HandlerFunc, error) {
	yaml, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error loading file: %s", err)
	}
	return YAMLHandler(yaml, fallback)
}
