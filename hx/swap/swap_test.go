package swap_test

import (
	"testing"

	"github.com/will-wow/typed-htmx-go/hx/swap"
)

func TestSwap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		builder swap.Builder
		want    string
	}{
		{
			name:    "Default",
			builder: *swap.New(),
			want:    "innerHTML",
		},
		{
			name:    "Strategy",
			builder: *swap.New().Strategy(swap.OuterHTML),
			want:    "outerHTML",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.builder.String()

			if got != tt.want {
				t.Errorf("got: %s, want: %s", got, tt.want)
			}
		})
	}
}
