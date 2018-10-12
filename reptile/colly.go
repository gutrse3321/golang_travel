package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"log"
	"strconv"
)

func main() {
	if choeeiBaseData() {
		log.Fatalln("ok")
	}
}

func choeeiBaseData() (flag bool) {
	var count int64 = 0
	for {
		count++
		if count > 72 {
			return true
		}

		var c = colly.NewCollector(
			colly.AllowedDomains("www.pornhub.com"),
			colly.Async(true),
		)

		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
		rp, err := proxy.RoundRobinProxySwitcher("http://127.0.0.1:1080", "https://127.0.0.1:1080")
		if err != nil {
			log.Fatalln(err)
		}
		c.SetProxyFunc(rp)
		//c.Async = true
		c.OnHTML("a.img", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			title := e.Attr("title")
			imgSrc := e.ChildAttr("img", "data-thumb_url")
			//videoUrl := videoDetailPage("https://www.pornhub.com"+link, c)
			fmt.Printf("Video found: %q, Link: %s\n", title, link)
			//fmt.Printf("Link Video: %s\n", videoUrl)
			fmt.Printf("Link image: %s\n\n", imgSrc)
		})

		c.OnRequest(func(r *colly.Request) {
			log.Println("Visiting", r.URL.String())
		})

		baseURL := "https://www.pornhub.com/video/search?search=japanese+%E7%84%A1+%E4%BF%AE%E6%AD%A3&page="
		var buffer bytes.Buffer
		strNum := strconv.FormatInt(count, 10)
		buffer.WriteString(baseURL)
		buffer.WriteString(strNum)
		c.Visit(buffer.String())
		c.Wait()
	}
}

//func videoDetailPage(url string, ctx *colly.Collector) (videoUrl string) {
//	ctx.OnHTML("head", func(e *colly.HTMLElement) {
//		videoUrl = e.ChildAttr("meta[name='twitter:player']", "content")
//		//source, exist := e.DOM.Children().Find("video").Children().Find("source").Attr("src")
//		//if !exist {
//		//	source = "视频不存在"
//		//}
//	})
//	ctx.Visit(url)
//	return
//}
