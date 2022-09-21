package urlshort

import (
	"fmt"
	"net/http"
	"os"
)

func RunUrlshort() (err error) {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	// 	yaml := `
	// - path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution
	// `

	yaml, err := readFile("./urlshort/paths.json")
	if err != nil {
		return err
	}

	yamlHandler, err := YAMLHandler(yaml, mapHandler)
	if err != nil {
		return err
	}

	fmt.Println("Starting the server on :8080")
	// http.ListenAndServe(":8080", mapHandler)

	if err := http.ListenAndServe(":8080", yamlHandler); err != nil {
		return err
	}

	return
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello, world!")
}

func readFile(path string) (yamlFile []byte, err error) {
	yamlFile, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return
}
