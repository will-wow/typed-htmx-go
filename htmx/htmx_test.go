package htmx_test

import (
	"testing"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/swap"
)

func TestHX(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		want  string
		attrs htmx.Builder
	}{
		{
			name:  "Boost",
			attrs: htmx.HX().Boost(true),
			want:  `hx-boost='true'`,
		},
		{
			name:  "Get",
			attrs: htmx.HX().Get("/example"),
			want:  `hx-get='/example'`,
		},
		{
			name:  "Post",
			attrs: htmx.HX().Post("/example"),
			want:  `hx-post='/example'`,
		},
		{
			name:  "On",
			attrs: htmx.HX().On("click", `alert("clicked")`),
			want:  `hx-on:click='alert("clicked")'`,
		},
		{
			name:  "OnHTMX",
			attrs: htmx.HX().OnHTMX("before-request", `alert("before")`),
			want:  `hx-on::before-request='alert("before")'`,
		},
		{
			name:  "PushURL",
			attrs: htmx.HX().PushURL(true),
			want:  `hx-push-url='true'`,
		},
		{
			name:  "PushURLPath",
			attrs: htmx.HX().PushURLPath("/example"),
			want:  `hx-push-url='/example'`,
		},
		{
			name:  "Select",
			attrs: htmx.HX().Select("#example"),
			want:  `hx-select='#example'`,
		},
		{
			name:  "SelectOOB",
			attrs: htmx.HX().SelectOOB("#info-details", "#other-details"),
			want:  `hx-select-oob='#info-details,#other-details'`,
		},
		{
			name: "SelectOOBWithStrategy",
			attrs: htmx.HX().SelectOOBWithStrategy(
				htmx.SelectOOBStrategy{Selector: "#info-details", Strategy: swap.AfterBegin},
				htmx.SelectOOBStrategy{Selector: "#other-details", Strategy: ""},
			),
			want: `hx-select-oob='#info-details:afterbegin,#other-details'`,
		},
		{
			name:  "Swap",
			attrs: htmx.HX().Swap(swap.OuterHTML),
			want:  `hx-swap='outerHTML'`,
		},
		{
			name: "SwapExtended",
			attrs: htmx.HX().SwapExtended(
				swap.New().
					Style(swap.OuterHTML).
					SettleTiming("1s").
					ShowElement("#example", swap.Top),
			),
			want: `hx-swap='outerHTML settle:1s show:#example:top'`,
		},
		{
			name:  "Target",
			attrs: htmx.HX().Target("#example"),
			want:  `hx-target='#example'`,
		},
		{
			name:  "TargetSpecial",
			attrs: htmx.HX().TargetSpecial(htmx.TargetThis),
			want:  `hx-target='this'`,
		},
		{
			name:  "TargetSelector",
			attrs: htmx.HX().TargetRelative(htmx.TargetSelectorClosest, "#example"),
			want:  `hx-target='closest #example'`,
		},
		{
			name:  "Vals",
			attrs: htmx.HX().Vals(map[string]int{"one": 1, "two": 2}),
			want:  `hx-vals='{"one":1,"two":2}'`,
		},
		{
			name:  "ValsJS",
			attrs: htmx.HX().ValsJS(map[string]int{"one": 1, "two": 2}),
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
