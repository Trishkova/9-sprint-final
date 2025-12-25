package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name        string
		size        int
		expectedLen int
	}{
		{
			name:        "Correct size",
			size:        100,
			expectedLen: 100,
		},
		{
			name:        "Zero size",
			size:        0,
			expectedLen: 0,
		},
		{
			name:        "Negative size",
			size:        -5,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			assert.Len(t, data, tt.expectedLen, "get lenght %d should be %d", len(data), tt.expectedLen)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Positive numbers",
			input:    []int{1, 5, 2, 9, 3},
			expected: 9,
		},
		{
			name:     "Mixed numbers",
			input:    []int{0, 1},
			expected: 1,
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Identical elements",
			input:    []int{7, 7, 7, 7},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			assert.Equal(t, tt.expected, result, "maximum(%v) should be %d", tt.input, tt.expected)
		})
	}
}
