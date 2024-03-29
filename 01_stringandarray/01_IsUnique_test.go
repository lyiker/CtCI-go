package stringandarray

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abcc", false},
		{" ", true},
		{"", true},
	}
	for _, c := range cases {
		actual := IsUnique(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %t, actual: %t\n", c.input, c.expected, actual)
		}
	}
}
