package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"log"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.pornhub.com"),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
	rp, err := proxy.RoundRobinProxySwitcher("http://127.0.0.1:1080", "https://127.0.0.1:1080")
	if err != nil {
		log.Fatalln(err)
	}
	c.SetProxyFunc(rp)

	c.OnHTML("head", func(e *colly.HTMLElement) {
		videoLink := e.ChildAttr("meta[name='twitter:player']", "content")
		fmt.Println("Link Video: ", videoLink)
	})

	//c.OnHTML(".mhp1138_videoWrapper", func(e *colly.HTMLElement) {
	//	source, exist := e.DOM.Children().Find("video").Children().Find("source").Attr("src")
	//	if !exist {
	//		source = "视频不存在"
	//	}
	//	fmt.Println("Link Video: ", source)
	//})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	baseURL := "https://www.pornhub.com/view_video.php?viewkey=ph5b7da169af7cd"
	c.Visit(baseURL)
}
