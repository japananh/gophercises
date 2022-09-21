package main

import (
	"github.com/japananh/gophercises/cyoa"
	"github.com/japananh/gophercises/link"
	"github.com/japananh/gophercises/quiz"
	"github.com/japananh/gophercises/sitemap"
	"github.com/japananh/gophercises/urlshort"
	"log"
)

func main() {
	if err := quiz.RunQuiz(); err != nil {
		log.Fatal(err)
	}

	if err := urlshort.RunUrlshort(); err != nil {
		log.Fatal(err)
	}

	if err := cyoa.RunServer("./cyoa/layout.html", "./cyoa/story.json"); err != nil {
		log.Fatal(err)
	}

	if err := cyoa.RunCLI("./cyoa/story.json"); err != nil {
		log.Fatal(err)
	}

	if err := link.Runner("./link/ex1.html"); err != nil {
		log.Fatal(err)
	}

	if err := sitemap.Crawl(); err != nil {
		log.Fatal(err)
	}
}
