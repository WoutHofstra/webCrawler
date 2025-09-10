package main

import (
        "fmt"
        "strings"
        "github.com/PuerkitoBio/goquery"
)


func getFirstParagraphFromHTML(html string) string {

        reader := strings.NewReader(html)
        doc, err := goquery.NewDocumentFromReader(reader)
        if err != nil {
                fmt.Println(err)
                return ""
        }

        mainList := doc.Find("main")
        if mainList.Size() >= 1 {

        	p1List := mainList.Find("p")

        	if p1List.Size() >= 1 {

        		firstP1 := p1List.First()
        		p1 := firstP1.Text()

        		return p1
	        }
        }

	newPList := doc.Find("p")
	if newPList.Size() <= 0 {
		return ""
	}
	newFirstP1 := newPList.First()
	newP1 := newFirstP1.Text()
	return newP1
}
