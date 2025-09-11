package main

import (
        "fmt"
        "strings"
        "github.com/PuerkitoBio/goquery"
        "net/url"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

        imgList := []string{}

        reader := strings.NewReader(htmlBody)
        doc, err := goquery.NewDocumentFromReader(reader)
        if err != nil {
                fmt.Println(err)
                return nil, err
        }

        doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
                href, ok := s.Attr("src")
                if !ok {
                        return
                }

                u, err := url.Parse(href)
                if err != nil {
                        return
                }

                abs := baseURL.ResolveReference(u).String()
                imgList = append(imgList, abs)
        })

        return imgList, nil
}
