package main

import (
	"github.com/japananh/gophercises/cyoa"
	"github.com/japananh/gophercises/link"
	"github.com/japananh/gophercises/quiz"
	"github.com/japananh/gophercises/sitemap"
	"github.com/japananh/gophercises/urlshort"
)

func main() {
	quiz.RunQuiz()
	urlshort.RunUrlshort()
	cyoa.RunServer("./cyoa/layout.html", "./cyoa/story.json")
	cyoa.RunCLI("./cyoa/story.json")
	link.Runner("./link/ex4.html")
	sitemap.Crawl()
}
