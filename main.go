package main


import (
	"fmt"
	"os"
	"net/url"
	"sync"
	"strconv"
)

type config struct {
	pages			map[string]int
	baseURL			*url.URL
	mu			*sync.Mutex
	concurrencyControl	chan struct{}
	wg			*sync.WaitGroup
	maxPages		int
}

func main() {

	args := os.Args
	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

    	maxConcurrency, err := strconv.Atoi(args[2])
    	if err != nil {
        	panic("maxConcurrency must be an integer")
    	}

    	maxPages, err := strconv.Atoi(args[3])
    	if err != nil {
        	panic("maxPages must be an integer")
    	}

        cfg := &config{
                pages:                  make(map[string]int),
                mu:                     &sync.Mutex{},
                concurrencyControl:     make(chan struct{}, maxConcurrency),
                wg:                     &sync.WaitGroup{},
		maxPages:		maxPages,
        }

	fmt.Printf("starting crawl of: %s\n", args[1])

        parsedBase, err := url.Parse(args[1])
        if err != nil {
                fmt.Println(err)
                return
        }
	cfg.baseURL = parsedBase

	cfg.crawlPage(args[1])
}
