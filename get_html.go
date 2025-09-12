package main

import (
	"io"
	"net/http"
	"fmt"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "BootCrawler/1.0")

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("got network error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %s", res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if i := strings.IndexByte(contentType, ';'); i != -1 {
    		contentType = contentType[:i]
	}

	contentType = strings.TrimSpace(strings.ToLower(contentType))
	if contentType != "text/html" {
    		return "", fmt.Errorf("got non-HTML response: %s", res.Header.Get("Content-Type"))
	}

	htmlBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}

	htmlBody := string(htmlBodyBytes)
	return htmlBody, nil
}





