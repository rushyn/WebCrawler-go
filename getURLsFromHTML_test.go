package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="https://blog.boot.dev"><span>Go to Boot.dev, you React Andy</span></a>
				</body>
			</html>
			`,
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name:     "absolute",
			inputURL: "https://test.boot.log",
			inputBody: `
			<html>
				<body>
					<p> this is a log </p>
					<a href="/log/yay">
						<span>loging</span>
					</a>
					<a href="https://log.com/path/inf">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://test.boot.log/log/yay", "https://log.com/path/inf"},
		},
	}
	

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if reflect.DeepEqual(actual, tc.expected) != true {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}