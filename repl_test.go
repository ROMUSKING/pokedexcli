package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Return to your    browser to continue with the next lesson",
			expected: []string{"return", "to", "your", "browser", "to", "continue", "with", "the", "next", "lesson"},
		},
		{
			input:    "Error: expected stdout to contain  Hello, World!",
			expected: []string{"error:", "expected", "stdout", "to", "contain", "hello,", "world!"},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		// add more casees here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Expected %s but got %s", expectedWord, word)
			}
		}
	}
}
