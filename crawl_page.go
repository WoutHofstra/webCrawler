package main

import (
	"fmt"
	"net/url"
	"strings"
	"os"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()

        cfg.mu.Lock()
	if cfg.maxPages <= len(cfg.pages) {
		os.Exit(0)
	}
        cfg.mu.Unlock()

	cfg.concurrencyControl <- struct{}{}
	defer func() { <-cfg.concurrencyControl }()

	parsedCurrent, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	domain1 := strings.TrimPrefix(cfg.baseURL.Hostname(), "www.")
	domain2 := strings.TrimPrefix(parsedCurrent.Hostname(), "www.")

	if domain1 != domain2 {
		fmt.Println("URL outside domain")
		return
	}

	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !cfg.addPageVisit(normalizedCurrent) {
		return
	}

	fmt.Println("Getting html from: ", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL)

	for _, u := range urls {

		fmt.Println(u)
		normalizedChild, err := normalizeURL(u)
		if err != nil {
			continue
		}

		if cfg.addPageVisit(normalizedChild) {
			cfg.wg.Add(1)
			go cfg.crawlPage(u)
		}
	}

	cfg.mu.Lock()
	for page, count := range cfg.pages {

		fmt.Println(page, count)

	}
	cfg.mu.Unlock()
}


func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {

	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.pages[normalizedURL] += 1
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}
