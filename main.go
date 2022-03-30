package main

import (
	"gophercises/cyoa"
	"gophercises/quiz"
	"gophercises/urlshort"
)

func main() {
	quiz.RunQuiz()
	urlshort.RunUrlshort()
	cyoa.RunServer("./cyoa/layout.html", "./cyoa/story.json")
}
