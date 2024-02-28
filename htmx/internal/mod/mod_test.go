package mod_test

import (
	"testing"

	"github.com/will-wow/typed-htmx-go/htmx/internal/mod"
)

type StringAlias string

const (
	One StringAlias = "one"
	Two StringAlias = "two"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		name      string
		main      string
		modifiers map[StringAlias]string
		want      string
	}{
		{
			name:      "just a name",
			main:      "main",
			modifiers: map[StringAlias]string{},
			want:      "main",
		},
		{
			name: "one modifier",
			main: "main",
			modifiers: map[StringAlias]string{
				One: "one",
			},
			want: "main one:one",
		},
		{
			name: "many modifiers are sorted",
			main: "main",
			modifiers: map[StringAlias]string{
				Two: "two",
				One: "one",
			},
			want: "main one:one two:two",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mod.Join(tt.main, tt.modifiers)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
