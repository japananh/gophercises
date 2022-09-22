package main

import (
	"flag"
	"log"

	"github.com/japananh/gophercises/cyoa"
	"github.com/japananh/gophercises/link"
	"github.com/japananh/gophercises/quiz"
	"github.com/japananh/gophercises/sitemap"
	"github.com/japananh/gophercises/task"
	"github.com/japananh/gophercises/urlshort"
)

const (
	Cyoa     = "cyoa"
	Link     = "link"
	Quiz     = "quiz"
	Sitemap  = "sitemap"
	UrlShort = "urlshort"
	Task     = "task"
)

func main() {
	defaultExercise := "task"
	exName := flag.String("ex-name", defaultExercise, "exercise name")

	switch *exName {
	case Cyoa:
		if err := cyoa.RunServer("./cyoa/layout.html", "./cyoa/story.json"); err != nil {
			log.Fatal(err)
		}
		if err := cyoa.RunCLI("./cyoa/story.json"); err != nil {
			log.Fatal(err)
		}
	case Link:
		if err := link.Runner("./link/ex1.html"); err != nil {
			log.Fatal(err)
		}
	case Quiz:
		if err := quiz.RunQuiz(); err != nil {
			log.Fatal(err)
		}
	case Sitemap:
		if err := sitemap.Crawl(); err != nil {
			log.Fatal(err)
		}
	case UrlShort:
		if err := urlshort.RunUrlshort(); err != nil {
			log.Fatal(err)
		}
	case Task:
		if err := task.InitDatabase(); err != nil {
			log.Fatal(err)
		}
		task.Execute()
	}
}
