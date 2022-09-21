package link

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func Runner(filename string) {
	s, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	links, err := Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range links {
		fmt.Println("Href: ", i.Href)
		fmt.Println("Text: ", i.Text)
	}
}

// Parse parses the HTML file and returns Links
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return getAllLinks(doc), nil
}

func getAllLinks(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				txt := extractText(n)
				links = append(links, Link{a.Val, txt})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		exLinks := getAllLinks(c)
		links = append(links, exLinks...)
	}
	return links
}

func extractText(n *html.Node) string {
	var text string
	if n.Type != html.ElementNode && n.Data != "a" && n.Type != html.CommentNode {
		text = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	return strings.Trim(text, "\n")
}
