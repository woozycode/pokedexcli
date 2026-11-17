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
			expected: []string{"Charmander", "Bulbasaur", "PIKACHU"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Printf("Actual word: %v -> expected word: %v", actual, c.expected)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual and length of expected do not match")
			fmt.Printf("Length of actual: %d -> length of expected: %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word and expectedWord do not match")
			}
		}
	}

	return
}
