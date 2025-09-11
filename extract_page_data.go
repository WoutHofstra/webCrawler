package main

import (
	"fmt"
	"net/url"
)

type PageData struct {
    URL            string
    H1             string
    FirstParagraph string
    OutgoingLinks  []string
    ImageURLs      []string
}


func extractPageData(html, pageURL string) (PageData, error) {

	var data PageData

	baseURL, err := url.Parse(pageURL)
	if err != nil {
		fmt.Println(err)
	}

	data.URL = pageURL
	data.H1 = getH1FromHTML(html)
	data.FirstParagraph = getFirstParagraphFromHTML(html)

	data.OutgoingLinks, err = getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Println(err)
	}

	data.ImageURLs, err = getImagesFromHTML(html, baseURL)
        if err != nil {
                fmt.Println(err)
        }

	return data, nil
}
