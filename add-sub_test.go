package main

import (
	"testing"
)

func testadd(t *testing.T) {
	tests := []struct {
		num1     int
		num2     int
		expected int
	}{
		{1, 5, 6},
		{2, 15, 17},
		{22, 30, 52},
		{23, 5, 28},
		{2, 191, 193},
		{1, 0, 1},
	}
	for _, tt := range tests {
		result := add(tt.num1, tt.num2)
		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}

}
