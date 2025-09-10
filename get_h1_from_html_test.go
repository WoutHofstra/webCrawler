package main

import (
        "testing"
)

func TestH1FromHTML(t *testing.T) {
        tests := []struct {
                name          string
                inputHTML     string
                expected      string
        }{
                {
                        name:     "get H1",
                        inputHTML: "<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>",
                        expected: "Welcome to Boot.dev",
                },
		{
			name:	"no H1 found",
			inputHTML: "<html><html>",
			expected: "",
		},
        }

        for i, tc := range tests {
                t.Run(tc.name, func(t *testing.T) {
                        actual, err := getH1FromHTML(tc.inputHTML)
                        if err != nil {
                                t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
                                return
                        }
                        if actual != tc.expected {
                                t.Errorf("Test %v - %s FAIL: expected output: %v, actual: %v", i, tc.name, tc.expected, actual)
                        }
                })
        }
}
