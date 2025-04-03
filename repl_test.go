package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual (%d) does not match expected (%d)", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual (%s) does not match expected (%s)", word, expectedWord)
			}
		}
	}
}
