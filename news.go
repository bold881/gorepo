package main

import (
	"fmt"
	"time"
)

func save(ch2Save chan PageItem, pIs PageItems) {
	for {
		pageItem := <-ch2Save
		//pIs.Add(pageItem.url, pageItem.content)
		fmt.Println(pageItem.url, pageItem.content)
	}
}

func add2crawl(ch2Crawl chan string, chPI chan PageItem) {
	for {
		url := <-ch2Crawl
		go CrawlGoQuery(url, chPI, ch2Crawl)
	}
}

func main() {
	seedUrls := []string{
		"http://www.cq.xinhuanet.com/",
	}

	var crawedItems PageItems
	crawedItems.Init()

	//var crawledUrls CrawledURLs
	crwedUrls.Init()

	ch2Crawl := make(chan string)
	chPageItem := make(chan PageItem)

	for _, url := range seedUrls {
		go CrawlGoQuery(url, chPageItem, ch2Crawl)
	}

	go save(chPageItem, crawedItems)

	//go add2crawl(ch2Crawl, chPageItem, crawledUrls)

	for {
		url := <-ch2Crawl
		go CrawlGoQuery(url, chPageItem, ch2Crawl)
		time.Sleep(10 * time.Millisecond)
	}

	//close(ch2Crawl)

	//close(chPageItem)
}
