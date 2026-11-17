package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "banana     apple strawberry",
			expected: []string{"banana", "apple", "strawberry"},
		},
		{
			input:    "hellorhighwater",
			expected: []string{"hellorhighwater"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " HelLO   WORLd",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Printf("Actual word slice: %v -> expected word slice: %v\n", actual, c.expected)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual slice %d and length of expected slice %d do not match\n", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v\n", c.input, actual, c.expected)
			}
		}
	}

	return
}
