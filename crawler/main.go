package main

import (
	"fmt"
	"os"
	"time"
)

const usage = `
usage:
	crawler <starting-url>
`

// linkq is the URL
// resultsq is hyperlinks we crawl
func worker(linkq chan string, resultsq chan []string) {
	// as the channel is closed, the loop shuts down.
	for link := range linkq {
		plinks, err := getPageLinks(link)
		if err != nil {
			fmt.Printf("error fetching %s: %v", link, err)
			//exit the current loop interation and start the next for iteration
			continue
		}
		fmt.Printf("%s (%d links\n", link, len(plinks.Links))
		time.Sleep(time.Millisecond * 500)
		// keep testing for whether there are links or not
		if len(plinks.Links) > 0 {
			// here, we dont want to get blocked
			// instead we delegate to a temp go routine
			// by making the write concurrent too, you are making sure that reads and writes can happen without
			//being sequential
			// But, would this guarantee accuracy?
			// here you are writing what you found to the channel.
			go func(links []string) {
				resultsq <- links
			}(plinks.Links)

		}
	}
}

//this method is wrong... why though
// all workers are trying to Read AND Write
// However, as long as somebody is writing, there is a block
// eventually we are running into the problem that everybody is reading and writing at the same time ---> DEADLOCK

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	nWorkers := 10
	// adding capacity to the chan is effectively adding more memory which allows workers to write more often (more space to do so)
	// So you improve speed by increasing number of workers and memory
	linkq := make(chan string, 1000)
	resultsq := make(chan []string, 1000)
	for i := 0; i < nWorkers; i++ {
		go worker(linkq, resultsq)
	}

	// links waiting to be fetched
	linkq <- os.Args[1]
	// map of string to boolean to track the URLS and whether we have visited it or not
	seen := map[string]bool{}

	// resultsq, the queue the worker writes URLs to
	for links := range resultsq {
		for _, link := range links {
			if !seen[link] {
				seen[link] = true
				linkq <- link
			}
		}
	}
}
