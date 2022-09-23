package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var g_crawlMutex sync.Mutex
var g_mapProcessedUrls map[string]bool
var g_sync sync.WaitGroup

func CrawlNewInstance(url string, depth int, fetcher Fetcher) {
	defer g_sync.Done()
	Crawl(url, depth, fetcher)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	g_crawlMutex.Lock()
	if !g_mapProcessedUrls[url] {
		// Set url is processed.
		g_mapProcessedUrls[url] = true
	}
	g_crawlMutex.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q depth: %v\n", url, body, depth)
	for _, u := range urls {
		// Parallel and sync.
		g_crawlMutex.Lock()
		bIsUrlProcessed := g_mapProcessedUrls[u]
		g_crawlMutex.Unlock()

		if !bIsUrlProcessed {
			// Start processing in new Gorutine.
			g_sync.Add(1)
			go CrawlNewInstance(u, depth-1, fetcher)
		}
	}
	return
}

func main() {
	// Init map processed urls.
	g_mapProcessedUrls = make(map[string]bool)
	Crawl("https://golang.org/", 4, fetcher)
	g_sync.Wait()
	fmt.Println("All Done\ng_mapProcessedUrls: ", g_mapProcessedUrls)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
