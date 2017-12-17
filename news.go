package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

func save(ch2Save chan PageItem, pIs PageItems, s *mgo.Session) {
	for {
		pageItem := <-ch2Save
		if !MgoSave(s, pageItem) {
			fmt.Println(pageItem.title, pageItem.meta)
		}
		//Mgotest()
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

	session, err := mgo.Dial("101.200.47.113")
	//session, err := mgo.Dial("10.115.0.29")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	EnsureIndex(session)

	// Get all scraped URLs
	urls := GetCrawledUrls(session)

	for _, szurl := range urls {
		crwedUrls.Add(szurl.Url)
	}

	ch2Crawl := make(chan string)
	chPageItem := make(chan PageItem)

	for _, url := range seedUrls {
		go CrawlGoQuery(url, chPageItem, ch2Crawl)
	}

	go save(chPageItem, crawedItems, session)

	//go add2crawl(ch2Crawl, chPageItem, crawledUrls)

	for {
		url := <-ch2Crawl
		go CrawlGoQuery(url, chPageItem, ch2Crawl)
		time.Sleep(10 * time.Millisecond)
	}

	//close(ch2Crawl)

	//close(chPageItem)
}
