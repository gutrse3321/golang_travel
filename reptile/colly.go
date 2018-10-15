package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"log"
	"strconv"
)

func main() {
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

	flag, baseBox := choeeiBaseData(c)
	if !flag {
		log.Fatalln("Fatal Fail!")
	}

	for _, item := range baseBox {
		tempVar := item.(map[string]interface{})
		videoDetailPage(tempVar["linkUrl"].(string))
	}
}

func choeeiBaseData(c *colly.Collector) (flag bool, baseBox []interface{}) {
	var count int64 = 0
	for {
		count++
		if count > 1 {
			return true, baseBox
		}

		c.OnHTML("a.img", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			title := e.Attr("title")
			imgSrc := e.ChildAttr("img", "data-thumb_url")
			//videoUrl := videoDetailPage("https://www.pornhub.com"+link, c)
			fmt.Printf("Video found: %q, Link: %s\n", title, link)
			//fmt.Printf("Link Video: %s\n", videoUrl)
			fmt.Printf("Link image: %s\n\n", imgSrc)
			avbox := make(map[string]interface{})
			avbox["title"] = title
			avbox["imgUrl"] = imgSrc
			avbox["linkUrl"] = link
			baseBox = append(baseBox, avbox)
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

func videoDetailPage(url string) {
	var buffer bytes.Buffer
	//https://www.pornhub.com/view_video.php?viewkey=ph5a34470258c5e
	buffer.WriteString("https://www.pornhub.com")
	buffer.WriteString(url)

	var err error
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	run, err := runner.New(runner.Flag("headless", true))
	runner.URL(buffer.String())
	if err != nil {
		log.Fatal(err)
	}

	err = run.Start(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	c, err := chromedp.New(ctxt, chromedp.WithRunner(run))
	if err != nil {
		log.Fatal(err)
	}

	var attr string
	var ok bool
	err = c.Run(ctxt, getVideoRealPath(buffer.String(), &attr, &ok))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("attr: ", attr)

	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func getVideoRealPath(s string, attr *string, ok *bool) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(s),
		chromedp.WaitReady(`source`),
		chromedp.AttributeValue(`source`, "src", attr, ok, chromedp.ByQuery),
	}
}
