package hx_test

import (
	"testing"

	"github.com/will-wow/typed-htmx-go/hx"
	"github.com/will-wow/typed-htmx-go/hx/swap"
)

func TestHX(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		want  string
		attrs hx.HX
	}{
		{
			name:  "Boost",
			attrs: hx.New().Boost(true),
			want:  `hx-boost='true'`,
		},
		{
			name:  "Get",
			attrs: hx.New().Get("/example"),
			want:  `hx-get='/example'`,
		},
		{
			name:  "Post",
			attrs: hx.New().Post("/example"),
			want:  `hx-post='/example'`,
		},
		{
			name:  "On",
			attrs: hx.New().On("click", `alert("clicked")`),
			want:  `hx-on:click='alert("clicked")'`,
		},
		{
			name:  "OnHTMX",
			attrs: hx.New().OnHTMX("before-request", `alert("before")`),
			want:  `hx-on::before-request='alert("before")'`,
		},
		{
			name:  "PushURL",
			attrs: hx.New().PushURL(true),
			want:  `hx-push-url='true'`,
		},
		{
			name:  "PushURLPath",
			attrs: hx.New().PushURLPath("/example"),
			want:  `hx-push-url='/example'`,
		},
		{
			name:  "Select",
			attrs: hx.New().Select("#example"),
			want:  `hx-select='#example'`,
		},
		{
			name:  "SelectOOB",
			attrs: hx.New().SelectOOB("#info-details", "#other-details"),
			want:  `hx-select-oob='#info-details,#other-details'`,
		},
		{
			name: "SelectOOBWithStrategy",
			attrs: hx.New().SelectOOBWithStrategy(
				hx.SelectOOBStrategy{Selector: "#info-details", Strategy: swap.AfterBegin},
				hx.SelectOOBStrategy{Selector: "#other-details", Strategy: ""},
			),
			want: `hx-select-oob='#info-details:afterbegin,#other-details'`,
		},
		{
			name:  "Swap",
			attrs: hx.New().Swap(swap.OuterHTML),
			want:  `hx-swap='outerHTML'`,
		},
		{
			name: "SwapExtended",
			attrs: hx.New().SwapExtended(
				swap.New().
					Style(swap.OuterHTML).
					SettleTiming("1s").
					ShowElement("#example", swap.Top),
			),
			want: `hx-swap='outerHTML settle:1s show:#example:top'`,
		},
		{
			name:  "Target",
			attrs: hx.New().Target("#example"),
			want:  `hx-target='#example'`,
		},
		{
			name:  "TargetSpecial",
			attrs: hx.New().TargetSpecial(hx.TargetThis),
			want:  `hx-target='this'`,
		},
		{
			name:  "TargetSelector",
			attrs: hx.New().TargetRelative(hx.TargetSelectorClosest, "#example"),
			want:  `hx-target='closest #example'`,
		},
		{
			name:  "Vals",
			attrs: hx.New().Vals(map[string]int{"one": 1, "two": 2}),
			want:  `hx-vals='{"one":1,"two":2}'`,
		},
		{
			name:  "ValsJS",
			attrs: hx.New().ValsJS(map[string]int{"one": 1, "two": 2}),
			want:  `hx-vals='js:{"one":1,"two":2}'`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.attrs.String()
			if got != tt.want {
				t.Errorf("got: %s, want: %s", got, tt.want)
			}
		})
	}
}
