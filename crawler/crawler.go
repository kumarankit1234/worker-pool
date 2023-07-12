package crawler

import (
	"fmt"
	"sync"
	"sync/atomic"
	"web-crawler/downloader"
	"web-crawler/fifo_queue"
	"web-crawler/htmlparser"
	"web-crawler/links"
)

type Crawler interface {
	Crawl(wg *sync.WaitGroup)
}

type crawlerImpl struct {
	queue          fifo_queue.Queue
	links          links.Links
	downloader     downloader.Downloader
	parser         htmlparser.HtmlParser
	workInProgress int32
}

func New(queue fifo_queue.Queue, links links.Links, downloader downloader.Downloader, parser htmlparser.HtmlParser, startingUrl string) Crawler {
	queue.Add(startingUrl)
	return &crawlerImpl{
		queue:          queue,
		links:          links,
		downloader:     downloader,
		parser:         parser,
		workInProgress: 0,
	}
}

func (c *crawlerImpl) canExit() bool {
	t := c.workInProgress == 0 && c.queue.IsEmpty()
	fmt.Println("can exit is ", t)
	return t
}

func (c *crawlerImpl) process(url string) {
	if url == "" {
		return
	}
	if c.links.IsVisited(url) {
		return
	}

	newLinks := []string{}
	content := c.downloader.Download(url)
	links := c.parser.Parse(content)
	for _, link := range links {
		if c.links.AreSameDomain(url, link) && !c.links.IsVisited(link) {
			go func() { c.queue.Add(link) }()
			newLinks = append(newLinks, link)
		}
	}
	fmt.Printf("Url is %s and links are ", url)
	fmt.Println(newLinks)
}

func (c *crawlerImpl) Crawl(wg *sync.WaitGroup) {

	for !c.canExit() {
		atomic.AddInt32(&c.workInProgress, 1)
		link := c.queue.Get()
		c.process(link)
		atomic.AddInt32(&c.workInProgress, -1)
	}

	wg.Done()
}
