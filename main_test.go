package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input string
		want  string
		err   bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"a0", "", false},
		{"f1b01", "", true},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"", "", false},
		{"a2", "aa", false},
		{"ор5", "оррррр", false},
	}

	for _, tt := range tests {
		got, err := unpackString(tt.input)
		if (err != nil) != tt.err {
			t.Errorf("unpackString(%q) error = %v, wantErr %v", tt.input, err, tt.err)
		}
		if got != tt.want {
			t.Errorf("unpackString(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
