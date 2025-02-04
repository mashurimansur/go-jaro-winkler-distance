package jaro

import (
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		s1, s2   string
		options  Options
		expected float64
	}{
		{"hello", "hello", Options{CaseSensitive: true}, 1.0},
		{"hello", "Hello", Options{CaseSensitive: true}, 0.825000},
		{"hello", "Hello", Options{CaseSensitive: false}, 1.0},
		{"hello", "world", Options{CaseSensitive: true}, 0.0},
		{"hello", "helo", Options{CaseSensitive: true}, 0.953333},
		{"abcd", "bcda", Options{CaseSensitive: true}, 0.833333},
		{"test", "tset", Options{CaseSensitive: true}, 1.0},
		{"", "test", Options{CaseSensitive: true}, 0.0},
		{"test", "", Options{CaseSensitive: true}, 0.0},
	}

	for _, tt := range tests {
		result := Distance(tt.s1, tt.s2, tt.options)
		if (result - tt.expected) > 0.01 {
			t.Errorf("Distance(%q, %q, %v) = %f; want %f", tt.s1, tt.s2, tt.options, result, tt.expected)
		}
	}
}
