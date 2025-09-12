# webCrawler

![Digging through websites since 2025](crawly.png)


A lightweight concurrent web crawler written in Go.  
It crawls a website starting from a given URL, collects links and images, and ensures concurrency is managed with goroutines, mutexes, and wait groups.

Features

🌐 Crawl a website: Starts from a base URL and recursively crawls internal links.  
🔗 Track links: Keeps track of how many times each URL has been discovered.  
🖼 Extract images: Collects image URLs from each page.  
⚡ Concurrent crawling: Uses goroutines to crawl multiple pages at once.  
🔒 Thread-safe storage: Shared state is protected with a sync.Mutex.  
⏳ Concurrency control: Limits the number of simultaneous HTTP requests using a buffered channel.  
🛑 Configurable max pages: Stops crawling after reaching a given maximum number of pages.  

Usage

go run . <startURL> <maxConcurrency> <maxPages>

Example:

go run . "https://www.wagslane.dev/
" 3 10

Arguments:

startURL → The URL to begin crawling.

maxConcurrency → The maximum number of concurrent requests (for example: 3).

maxPages → The maximum number of pages to crawl (for example: 10).

How it works

A config struct holds:  
pages → a map[string]int that tracks visited URLs and their counts  
baseURL → the original domain, used to avoid crawling external sites  
mu → a mutex for safe concurrent access to pages  
concurrencyControl → a buffered channel to control goroutine concurrency  
wg → a wait group to wait for all goroutines to finish  
maxPages → a limit on how many pages to crawl  

The crawler:  
Normalizes and checks URLs against the base domain.  
Downloads the HTML of each page.  
Extracts links (a[href]) and images (img[src]).  
Adds new URLs to the crawl queue if they haven’t been seen.  
Stops once the maxPages limit is reached.  
Requirements  
Go 1.20+

Internet access to crawl live sites

Notes

To debug, you can set maxConcurrency to 1 to see crawling happen sequentially.  
Crawling large websites can be slow or blocked by rate limiting.  
Only internal links (same domain as startURL) are crawled.
