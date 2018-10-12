package main

import (
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"time"
)

func test() {
	getJokes()
}

func getJokes() {
	proxy, err := url.Parse("http://127.0.0.1:1080")
	if err != nil {
		log.Fatalln(err)
	}

	tr := &http.Transport{
		Proxy: http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout: time.Second * 5,
	}

	res, err := client.Get("https://www.pornhub.com/video/search?search=japanese")
	//res, err := client.Get("https://www.qiushibaike.com/")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d, %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(doc.Find(".footer").Text())
	doc.Find("#networkbar_items_wrap > li").Each(func(i int, s *goquery.Selection) {

	})

	//doc.Find(".content").Each(func(i int, s *goquery.Selection) {
	//	fmt.Println(s.Text())
	//})
}


