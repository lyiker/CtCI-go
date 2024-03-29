package stringandarray

import (
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"abcd", "abcd", true},
		{"abcd", "abdc", true},
		{"abcc", "ccbb", false},
		{"abcc", "abcc ", false},
		{" ", " ", true},
		{"", "", true},
	}
	for _, c := range cases {
		actual := CheckPermutation(c.input1, c.input2)
		if actual != c.expected {
			t.Fatalf("Inputs %s, %s. Expected: %t, actual: %t\n",
				c.input1, c.input2, c.expected, actual)
		}
	}
}
