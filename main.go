package main

import (
	"sync"
	"web-crawler/crawler"
	"web-crawler/downloader"
	"web-crawler/fifo_queue"
	"web-crawler/htmlparser"
	links2 "web-crawler/links"
)

func main() {
	var close sync.WaitGroup
	fifoQueue := fifo_queue.New()
	parser := htmlparser.New()
	links := links2.New()
	htmlDownloader := downloader.New()
	webCrawler := crawler.New(fifoQueue, links, htmlDownloader, parser, "abc")
	for i := 1; i <= 5; i++ {
		close.Add(1)
		go webCrawler.Crawl(&close)
	}
	close.Wait()
}
