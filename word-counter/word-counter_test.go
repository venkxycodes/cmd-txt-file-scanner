package word_counter

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    " 123 Go!! Gophers",
			expected: []string{"go", "gophers"},
		},
		{
			input:    "No-symbols-here",
			expected: []string{"no", "symbols", "here"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			"$$$ %%%",
			[]string{},
		},
	}
	for _, tt := range tests {
		got := tokenize(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("tokenize(%q) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
