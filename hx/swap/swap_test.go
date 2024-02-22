package swap_test

import (
	"testing"
	"time"

	"github.com/will-wow/typed-htmx-go/hx/swap"
)

func TestSwap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		builder *swap.Builder
		want    string
	}{
		{
			name:    "Default",
			builder: swap.New(),
			want:    "innerHTML",
		},
		{
			name:    "Strategy",
			builder: swap.New().Strategy(swap.OuterHTML),
			want:    "outerHTML",
		},
		{
			name:    "Transition",
			builder: swap.New().Transition(),
			want:    "innerHTML transition:true",
		},
		{
			name:    "SwapTiming",
			builder: swap.New().Swap(500 * time.Millisecond),
			want:    "innerHTML swap:500ms",
		},
		{
			name:    "SettleTiming",
			builder: swap.New().Settle(500 * time.Millisecond),
			want:    "innerHTML settle:500ms",
		},
		{
			name:    "Clear",
			builder: swap.New().Transition().Clear(swap.Transition),
			want:    "innerHTML",
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
