package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	var i int
	for {
		i++
		log.Println("count: ", i)
		c := colly.NewCollector(
			colly.AllowedDomains("sikii.cc"),
			colly.MaxDepth(1000),
			colly.Async(true),
		)
		c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1000})
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			fmt.Printf("route: %q -> %s\n", e.Text, link)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})

		c.Visit("http://sikii.cc/")
		if i == 100000 {
			break
		}
	}
}
