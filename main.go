package main

import (
	"gophercises/cyoa"
	"gophercises/link"
	"gophercises/quiz"
	"gophercises/sitemap"
	"gophercises/urlshort"
)

func main() {
	quiz.RunQuiz()
	urlshort.RunUrlshort()
	cyoa.RunServer("./cyoa/layout.html", "./cyoa/story.json")
	cyoa.RunCLI("./cyoa/story.json")
	link.Runner("./link/ex4.html")
	sitemap.Crawl()
}
