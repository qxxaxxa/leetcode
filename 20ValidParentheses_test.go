package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {

	cases := []struct {
		argu     string
		excepted bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"", true},
		{"{[]}", true},
		{"[", false},
		{"}{", false},
		{"][", false},
	}
	for _, c := range cases {
		result := isValid(c.argu)
		if result != c.excepted {

			t.Fatalf("add function failed ,arg: %s,excepted:%t", c.argu, c.excepted)
		}
	}

}

