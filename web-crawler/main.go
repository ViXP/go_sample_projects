package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"golang.org/x/net/html"
)

var fetched map[string]bool

type result struct {
	url   string
	links []string
	err   error
	depth int16
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	fetched = make(map[string]bool)
	now := time.Now()
	Crawl("http://google.com", 2)
	fmt.Println("time taken: ", time.Since(now))
}

func Crawl(url string, depth int16) {
	channel := make(chan *result)

	defer close(channel)

	go fetch(url, depth, &channel)
	fetched[url] = true

	for fetching := 1; fetching > 0; fetching-- {
		res := <-channel

		if res.err != nil {
			continue
		}

		fmt.Printf("found url: %s\n", res.url)

		if res.depth > 0 {
			for _, u := range res.links {
				if !fetched[u] {
					fetching++
					go fetch(u, res.depth-1, &channel)
					fetched[u] = true
				}
			}
		}
	}
}

func fetch(url string, depth int16, ch *chan *result) {
	links, err := findLinks(url)
	*ch <- &result{url, links, err, depth}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return visit(nil, doc), nil
}

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = visit(links, child)
	}

	return links
}
