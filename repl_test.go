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
			input:    "",
			expected: []string{},
		},
		{
			input:    "SingleWord",
			expected: []string{"singleword"},
		},
		{
			input:    "A SentANCE WITH mixed CASE",
			expected: []string{"a", "sentance", "with", "mixed", "case"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Len of actual doesn't match Expected")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%v does not match %v", word, expectedWord)
			}
		}
	}
}
