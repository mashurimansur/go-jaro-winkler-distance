package jaro

import (
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		s1, s2        string
		caseSensitive bool
		expected      float64
	}{
		{"hello", "hello", true, 1.0},
		{"hello", "Hello", true, 0.825000},
		{"hello", "Hello", false, 1.0},
		{"hello", "world", true, 0.0},
		{"hello", "helo", true, 0.953333},
		{"abcd", "bcda", true, 0.833333},
		{"test", "tset", true, 1.0},
		{"", "test", true, 0.0},
		{"test", "", true, 0.0},
	}

	for _, tt := range tests {
		result := Distance(tt.s1, tt.s2, tt.caseSensitive)
		if (result - tt.expected) > 0.01 {
			t.Errorf("Distance(%q, %q, %v) = %f; want %f", tt.s1, tt.s2, tt.caseSensitive, result, tt.expected)
		}
	}
}
