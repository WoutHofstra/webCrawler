package main


import (
	"fmt"
	"net/url"
	"strings"
)


func normalizeURL(rawUrl string) (string, error) {

	url, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	list := []string{url.Host, url.Path}

	newUrl := strings.Join(list, "")

	return newUrl, nil
}
