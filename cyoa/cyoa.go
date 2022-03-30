package cyoa

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type StoryMap struct {
	Intro  Story            `json:"intro"`
	Others map[string]Story `json:"-"`
}

func (s *StoryMap) UnMarshal(fileContent []byte) error {
	err := json.Unmarshal(fileContent, &s.Others)

	if err != nil {
		return err
	}

	if n, ok := s.Others["intro"]; ok {
		s.Intro = n
		delete(s.Others, "intro")
	}

	return nil
}

type State struct {
	PageTitle string
	Content   Story
}

func RunServer(htmlPath string, jsonPath string) {
	html, err := template.ParseFiles(htmlPath)
	checkErr(err)
	tmpl := template.Must(html, err)

	fileContent, err := ioutil.ReadFile(jsonPath)
	checkErr(err)

	var data StoryMap
	err = data.UnMarshal(fileContent)
	checkErr(err)

	pageTitle := "Choose Your Own Adventure"
	state := State{PageTitle: pageTitle, Content: data.Intro}
	handler("/cyoa", state, tmpl)

	for key, element := range data.Others {
		state := State{PageTitle: pageTitle, Content: element}
		go handler("/cyoa/"+key, state, tmpl)
	}

	err = http.ListenAndServe(":80", nil)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("Program exited due to error: ", err)
	}
}

func handler(path string, state State, tmpl *template.Template) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, state)
	})
}
