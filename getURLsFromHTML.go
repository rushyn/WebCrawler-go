package main

import (
	"strings"

	"golang.org/x/net/html"
)


func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	reader := strings.NewReader(htmlBody)
	node, err := html.Parse(reader)
	if err != nil {
		return []string{}, err
	}

	links := findLinks(node, rawBaseURL)

    return links, err

}

func findLinks(n *html.Node, rawBaseURL string) []string {

	links := []string{}
	
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr{
			if a.Key == "href" {
				if strings.Contains(a.Val, rawBaseURL) || strings.Contains(a.Val, "http"){
					links = append(links, a.Val)
				}else{
					links = append(links, rawBaseURL + a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, findLinks(c, rawBaseURL)...)
	}

	return links
}