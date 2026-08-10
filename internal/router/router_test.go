package router

import (
	"strings"
	"testing"
)

func TestChoose(t *testing.T) {
	tests := []struct{ prompt, want string }{
		{"hello", "local"},
		{"please translate this", "cheap"},
		{"debug this stack trace", "normal"},
		{"review this architecture and migration plan", "strong"},
		{strings.Repeat("x", 3001), "strong"},
	}
	for _, tt := range tests {
		if got := Choose(tt.prompt).Route; got != tt.want {
			t.Errorf("Choose(%q)=%q, want %q", tt.prompt, got, tt.want)
		}
	}
}
