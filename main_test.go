package main

import "testing"

func TestConvertToUpperCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},
		{"world", "WORLD"},
		{"GoLang", "GOLANG"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ConvertToUpperCase(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
