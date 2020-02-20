package main

import (
	"testing"
)

func TestIncrementByOne(t *testing.T) {
	if IncrementByOne(2) != 3 {
		t.Error("expected 2+1 = 3")
	}
}

func TestTableIncrementByOnde(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 3},
		{-1, 0},
		{4, 5},
	}

	for _, test := range tests {
		if output := IncrementByOne(test.input); output != test.expected {
			t.Error("test failed for", test.input, test.expected, output)
		}

	}
}
