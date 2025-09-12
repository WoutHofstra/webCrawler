package main


import (
	"fmt"
	"os"
	"net/url"
	"sync"
)

type config struct {
	pages			map[string]int
	baseURL			*url.URL
	mu			*sync.Mutex
	concurrencyControl	chan struct{}
	wg			*sync.WaitGroup
}

func main() {

	cfg := &config{
		pages:			make(map[string]int),
		mu:			&sync.Mutex{},
		concurrencyControl:	make(chan struct{}, 5),
		wg:			&sync.WaitGroup{},
	}

	args := os.Args
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
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
