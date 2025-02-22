package main

import (
	"slices"
	"testing"
)

func TestMainLogic(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Positive numbers",
			input:    []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "Negative numbers",
			input:    []int{-1, -2, -3, -4, -5},
			expected: -1,
		},
		{
			name:     "Mixed numbers",
			input:    []int{-1, 2, 0, -4, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maxVal := slices.Max(tt.input)
			if maxVal != tt.expected {
				t.Errorf("slices.Max() = %v, expected %v", maxVal, tt.expected)
			}
		})
	}
}
