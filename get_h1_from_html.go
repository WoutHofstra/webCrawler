package main

import (
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
)


func getH1FromHTML(html string) string {

	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	h1List := doc.Find("h1")
	if h1List.Size() <= 0 {
		return ""
	}
	h1 := h1List.Text()

	return h1

}
