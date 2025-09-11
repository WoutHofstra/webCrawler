package main

import (
        "fmt"
        "strings"
        "github.com/PuerkitoBio/goquery"
	"net/url"
)


func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

	UrlList := []string{}

        reader := strings.NewReader(htmlBody)
        doc, err := goquery.NewDocumentFromReader(reader)
        if err != nil {
                fmt.Println(err)
                return nil, err
        }

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			return
		}

		u, err := url.Parse(href)
		if err != nil {
			return
		}

		abs := baseURL.ResolveReference(u).String()
		UrlList = append(UrlList, abs)
	})

        return UrlList, nil
}
