package main

import (
	"flag"
	"log"

	"github.com/japananh/gophercises/cyoa"
	"github.com/japananh/gophercises/link"
	"github.com/japananh/gophercises/phone"
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
	Phone    = "phone"
)

func main() {
	defaultExercise := "task"
	exName := flag.String("ex-name", defaultExercise, "exercise name")

	switch *exName {
	case Cyoa:
		if err := cyoa.RunServer("./cyoa/layout.html", "./cyoa/story.json"); err != nil {
			log.Fatalln(err)
		}
		if err := cyoa.RunCLI("./cyoa/story.json"); err != nil {
			log.Fatalln(err)
		}
	case Link:
		if err := link.Runner("./link/ex1.html"); err != nil {
			log.Fatalln(err)
		}
	case Quiz:
		if err := quiz.RunQuiz(); err != nil {
			log.Fatalln(err)
		}
	case Sitemap:
		if err := sitemap.Crawl(); err != nil {
			log.Fatalln(err)
		}
	case UrlShort:
		if err := urlshort.RunUrlshort(); err != nil {
			log.Fatalln(err)
		}
	case Task:
		if err := task.InitDatabase(); err != nil {
			log.Fatalln(err)
		}
		if err := task.Execute(); err != nil {
			log.Fatalln(err)
		}
	case Phone:
		if err := phone.Start(); err != nil {
			log.Fatalln(err)
		}
	}
}
