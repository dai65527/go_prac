package main

import "testing"

func TestIsEven(t *testing.T) {
	cases := []struct {
		name     string
		input    int
		expected bool
	}{
		{name: "+odd", input: 21, expected: false},
		{name: "+even", input: 42, expected: true},
		{name: "-odd", input: -21, expected: false},
		{name: "-even", input: -42, expected: true},
		{name: "zero", input: 0, expected: true},
		{name: "intmin", input: -2147483648, expected: true},
		{name: "intmax", input: 2147483647, expected: false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := IsEven(c.input); c.expected != actual {
				t.Errorf("want IsOdd(%d) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}
