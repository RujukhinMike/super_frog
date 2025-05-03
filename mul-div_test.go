package main

import (
	"testing"
)

func testmul(t *testing.T) {
	tests := []struct {
		num1     int
		num2     int
		expected int
	}{
		{1, 2, 2},
		{2, 5, 10},
		{5, 8, 40},
		{5, 0, 0},
	}
	for _, testiki := range tests {
		got := mul(testiki.num1, testiki.num2)
		if got != testiki.expected {
			t.Errorf("expected %d, got %d", testiki.expected, got)
		}
	}
}

func testdiv(t *testing.T) {
	tests := []struct {
		num1     int
		num2     int
		expected int
	}{
		{1, 2, 0},
		{2, 0, 0},
		{10, 2, 5},
		{5, 5, 1},
	}
	for _, testiki := range tests {
		got := div(testiki.num1, testiki.num2)
		if got != testiki.expected {
			t.Errorf("expected %d, got %d", testiki.expected, got)
		}
	}
}
