package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

type List struct {
	Title    string
	Subtitle string
	Contents []string
}

// func renderNode(n *html.Node) string {
// 	var buf bytes.Buffer
// 	w := io.Writer(&buf)
// 	html.Render(w, n)
// 	return buf.String()
// }

/*
	Crawling Example useing Colly
	http://go-colly.org/
	for React Version page
	https://github.com/facebook/react/blob/main/CHANGELOG.md
*/
func Crawler() {

	c := colly.NewCollector(
		// Visit only domains: github.com
		colly.AllowedDomains("github.com"),
	)
	l := make([]List, 0)
	var b List

	// <ul> 태그 일 경우 > <li> 태그 안의 값들을 정리하는 함수

	var organize func(*html.Node)
	organize = func(n *html.Node) {

		if n.Type == 1 && n.Data != "a" {
			fmt.Printf(n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			organize(c)
		}
	}

	//e.Name : 해당 태그
	//e.Text : 태그안의 값

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("article", func(_ int, el *colly.HTMLElement) {
			for index, value := range el.DOM.Children().Nodes {
				switch value.Data {
				case "h2":
					l = append(l, b)
					b.Title = value.LastChild.Data
					// fmt.Println("h2", value.LastChild.Data)
				case "h3":
					b.Subtitle = value.LastChild.Data
					// fmt.Println("h3", value.LastChild.Data)
				// case "h4":
				// case "p":
				case "ul":
					fmt.Println(index)
					organize(value)
					// fmt.Println(value.LastChild.Data)
				default:
					// fallthrough
					// fmt.Println("-")
				}

			}
			fmt.Println(l)
		})

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://github.com/facebook/react/blob/main/CHANGELOG.md")

}
