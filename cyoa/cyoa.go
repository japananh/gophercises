package cyoa

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type storyMap map[string]Story

func (s *storyMap) UnMarshal(fileContent []byte) error {
	return json.Unmarshal(fileContent, &s)
}

type State struct {
	PageTitle string
	Content   Story
}

func RunServer(htmlPath string, jsonPath string) (err error) {
	html, err := template.ParseFiles(htmlPath)
	if err != nil {
		return err
	}

	tmpl := template.Must(html, err)

	fileContent, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	var data storyMap
	err = data.UnMarshal(fileContent)
	if err != nil {
		return err
	}

	PageTitle := "Choose Your Own Adventure"
	state := State{PageTitle: PageTitle, Content: data["intro"]}
	handler("/cyoa", state, tmpl)

	for key, element := range data {
		state := State{PageTitle: PageTitle, Content: element}
		go handler("/cyoa/"+key, state, tmpl)
	}

	if err = http.ListenAndServe(":80", nil); err != nil {
		return err
	}

	return
}

func handler(path string, data State, tmpl *template.Template) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.Execute(w, data)
	})
}

func RunCLI(jsonPath string) (err error) {
	fileContent, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	var data storyMap
	err = data.UnMarshal(fileContent)
	if err != nil {
		return err
	}

	quit := make(chan bool)
	arcChan := make(chan string, 1)

	arc := flag.String("arc", "intro", "story arc name")
	flag.Parse()
	arcChan <- *arc

	fmt.Println("---------------- Welcome to Choose Your Own Adventure! --------------------")

	go func() {
		for {
			select {
			case <-quit:
				return
			case arc := <-arcChan:
				if story, ok := data[arc]; ok {
					readStory(story, arcChan, quit)
				} else {
					quit <- true
				}
			}
		}
	}()

	<-quit
	fmt.Println("----------------------- The End! --------------------------")

	return
}

func readStory(story Story, arcChan chan<- string, quit chan bool) {
	fmt.Printf("\n--- %s ---\n\n", story.Title)
	for _, para := range story.Story {
		fmt.Println("    " + para)
	}
	fmt.Println()
	for i, option := range story.Options {
		fmt.Printf("--- Option %d: %s\n", i+1, option.Text)
	}
	if len(story.Options) == 0 {
		quit <- true
		return
	}

	fmt.Println("\nPlease enter a number to choose your option")
	fmt.Print("-> ")

	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		text = strings.TrimSpace(text)
		text = strings.Replace(text, "\n", "", -1) // convert CRLF to LF

		if selection, _ := strconv.Atoi(text); selection >= 1 && selection <= len(story.Options) {
			arcChan <- story.Options[selection-1].Arc
			break
		}

		fmt.Print("Oops. Invalid option. Please try again.\n-> ")
	}
}
