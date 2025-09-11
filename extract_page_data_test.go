package main

import (
	"testing"
)

func TestExtractPageData(t *testing.T) {
	html := `
		<html>
			<head><title>Test Page</title></head>
			<body>
				<h1>Hello World</h1>
				<p>This is the first paragraph.</p>
				<a href="/about">About</a>
				<img src="/logo.png" />
			</body>
		</html>`

	pageURL := "http://example.com/"

	data, err := extractPageData(html, pageURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data.H1 != "Hello World" {
		t.Errorf("expected H1 = 'Hello World', got %q", data.H1)
	}

	if data.FirstParagraph != "This is the first paragraph." {
		t.Errorf("expected first paragraph, got %q", data.FirstParagraph)
	}

	if len(data.OutgoingLinks) == 0 || data.OutgoingLinks[0] != "http://example.com/about" {
		t.Errorf("expected outgoing link to resolve, got %+v", data.OutgoingLinks)
	}

	if len(data.ImageURLs) == 0 || data.ImageURLs[0] != "http://example.com/logo.png" {
		t.Errorf("expected image URL to resolve, got %+v", data.ImageURLs)
	}
}
