package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("CorrectSize", func(t *testing.T) {
		size := 100
		data := generateRandomElements(size)
		if len(data) != size {
			t.Errorf("Expected size %d, get size %d", size, len(data))
		}
	})

	t.Run("ZeroSize", func(t *testing.T) {
		data := generateRandomElements(0)
		if len(data) != 0 {
			t.Errorf("Expected empty slice, get size %d", len(data))
		}
	})

	t.Run("NegativeSize", func(t *testing.T) {
		data := generateRandomElements(-5)
		if len(data) != 0 {
			t.Errorf("Expected empty slice, get size %d", len(data))
		}
	})

	t.Run("CheckValues", func(t *testing.T) {
		size := 10
		data := generateRandomElements(size)

		allZeros := true
		for _, v := range data {
			if v != 0 {
				allZeros = false
				break
			}
		}

		if allZeros {
			t.Log("Warning: all of numbers equal zero")
		}
	})
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
			name:     "Negative numbers",
			input:    []int{-10, -5, -20, -2},
			expected: -2,
		},
		{
			name:     "Mixed numbers",
			input:    []int{-1, 0, 1},
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
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
