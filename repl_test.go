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
			input:    " hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " THIS IS         A TEST STRING",
			expected: []string{"this", "is", "a", "test", "string"},
		},
		{
			input:    "               test",
			expected: []string{"test"},
		},
		{
			input:    "testing is always appropriate",
			expected: []string{"testing", "is", "always", "appropriate"},
		},
		{
			input:    "   Help me    I'm under the water      ",
			expected: []string{"help", "me", "i'm", "under", "the", "water"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual: %v not equal to Expected: %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Word: %v not equal to Expected Word: %v", word, expectedWord)
			}
		}
	}
}
