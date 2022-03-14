package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type configs struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if v, found := pathsToUrls[r.URL.Path]; found {
			http.Redirect(w, r, v, http.StatusSeeOther)
		}
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
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(data)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(data []byte) ([]configs, error) {
	m := make([]configs, 0)

	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func buildMap(parsedYaml []configs) map[string]string {
	result := make(map[string]string)
	for _, config := range parsedYaml {
		result[config.Path] = config.Url
	}
	return result
}
