package main

import (
        "testing"
)

func TestP1FromHTML(t *testing.T) {
        tests := []struct {
                name          string
                inputHTML     string
                expected      string
        }{
                {
                        name:     "get P1",
                        inputHTML: "<html><body><h1>Welcome to Boot.dev</h1><main><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></main></body></html>",
                        expected: "Learn to code by building real projects.",
                },
                {
                        name:   "no P1 found",
                        inputHTML: "<html><html>",
                        expected: "",
                },
                {
                        name:   "invalid html document",
                        inputHTML: "this is the invalid document",
                        expected: "",
                },
        }

        for i, tc := range tests {
                t.Run(tc.name, func(t *testing.T) {
                        actual := getFirstParagraphFromHTML(tc.inputHTML)
                        if actual != tc.expected {
                                t.Errorf("Test %v - %s FAIL: expected output: %v, actual: %v", i, tc.name, tc.expected, actual)
                        }
                })
        }
}
