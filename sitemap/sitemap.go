package sitemap

import (
	"encoding/xml"
	"flag"
	"github.com/japananh/gophercises/link"
	"golang.org/x/exp/slices"
	"io"
	"net/http"
	urlLib "net/url"
	"os"
	"strings"
)

type input struct {
	url   string
	depth int
}

type urlset struct {
	XmlNS string `xml:"xmlns,attr"`
	Urls  []Url  `xml:"url"`
}

type Url struct {
	Loc string `xml:"loc"`
}

func Crawl() (err error) {
	input := readFlags("https://www.calhoun.io/", 1)
	var visited []string
	linkList, err := iterateUrlListToExtractLink(input.url, visited, input.depth)
	if err != nil {
		return err
	}

	// create a xml file
	filePath := "./sitemap/sitemap.xml"
	xmlFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	// write the header to xml file
	xmlWriter := io.Writer(xmlFile)
	if _, err := xmlFile.WriteString(xml.Header); err != nil {
		return err
	}

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("", "    ")

	// write xml content to the created file
	content := new(urlset)
	content.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"
	for _, item := range linkList {
		content.Urls = append(content.Urls, Url{Loc: item})
	}
	if err := enc.Encode(&content); err != nil {
		return err
	}

	return nil
}

func iterateUrlListToExtractLink(domain string, visited []string, depth int) ([]string, error) {
	if depth == 0 {
		return visited, nil
	}

	result, err := extractLinkFromUrl(domain)
	if err != nil {
		return nil, err
	}

	depth -= 1
	for _, item := range result {
		if slices.Contains(visited, item.Href) || item.Href == "" {
			continue
		}
		visited = append(visited, item.Href)
		if _, err := iterateUrlListToExtractLink(item.Href, visited, depth); err != nil {
			continue
		}
	}

	return visited, nil
}

func extractLinkFromUrl(url string) ([]link.Link, error) {
	// parse domain to get HTML
	body, err := crawlHTML(url)
	if err != nil {
		return nil, err
	}

	// parse HTML to get <a href="...">...</a>
	result, err := link.Parse(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	return remapURL(result, url)
}

// Remap href with the default domain if needed
// Ex: domain: https://domain.com and href: /about -> https://domain.com/about
func remapURL(list []link.Link, url string) ([]link.Link, error) {
	// get domain from url
	urlInfo, err := urlLib.Parse(url)
	if err != nil {
		return nil, err
	}

	var result []link.Link
	for _, item := range list {
		if item.Href == "" {
			continue
		}

		if item.Href[0:1] == "/" {
			result = append(result, link.Link{Href: urlInfo.Scheme + "://" + urlInfo.Host + item.Href, Text: item.Text})
			continue
		}

		result = append(result, item)
	}

	return result, nil
}

// get HTML from an url
func crawlHTML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func readFlags(defaultURL string, defaultDepth int) *input {
	url := flag.String("url", defaultURL, "an url")
	depth := flag.Int("depth", defaultDepth, "maximum number of links to follow when building a sitemap")
	flag.Parse()
	return &input{url: *url, depth: *depth}
}
