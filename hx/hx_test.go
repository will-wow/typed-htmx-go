package hx_test

import (
	"testing"
	"time"

	"github.com/will-wow/typed-htmx-go/hx"
	"github.com/will-wow/typed-htmx-go/hx/swap"
	"github.com/will-wow/typed-htmx-go/hx/trigger"
)

func TestHX(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		want  string
		attrs *hx.HX
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
					Strategy(swap.OuterHTML).
					Settle(time.Second).
					ShowElement("#example", swap.Top),
			),
			want: `hx-swap='outerHTML settle:1s show:#example:top'`,
		},
		{
			name:  "SwapOOB",
			attrs: hx.New().SwapOOB(),
			want:  `hx-swap-oob='true'`,
		},
		{
			name:  "SwapOOBWithStrategy",
			attrs: hx.New().SwapOOBWithStrategy(swap.AfterBegin),
			want:  `hx-swap-oob='afterbegin'`,
		},
		{
			name:  "SwapOOBSelector",
			attrs: hx.New().SwapOOBSelector(swap.AfterBegin, "#example"),
			want:  `hx-swap-oob='afterbegin:#example'`,
		},
		{
			name:  "Target",
			attrs: hx.New().Target("#example"),
			want:  `hx-target='#example'`,
		},
		{
			name:  "Target non-standard",
			attrs: hx.New().Target(hx.TargetThis),
			want:  `hx-target='this'`,
		},
		{
			name: "TargetSelector",
			attrs: hx.New().Target(
				hx.TargetRelative(hx.Closest, "#example"),
			),
			want: `hx-target='closest #example'`,
		},
		{
			name:  "Trigger",
			attrs: hx.New().Trigger("click"),
			want:  `hx-trigger='click'`,
		},
		{
			name: "TriggerExtended",
			attrs: hx.New().TriggerExtended(
				trigger.NewEvent("click").Filter("ctrlKey").Target("#element"),
				trigger.NewPoll(time.Second),
				trigger.NewIntersectEvent().Root("#element").Threshold(0.2),
			),
			want: `hx-trigger='click[ctrlKey] target:#element, every 1s, intersect root:#element threshold:0.2'`,
		},
		{
			name:  "Vals",
			attrs: hx.New().Vals(map[string]int{"one": 1, "two": 2}),
			want:  `hx-vals='{"one":1,"two":2}'`,
		},
		{
			name:  "ValsJS",
			attrs: hx.New().ValsJS(map[string]string{"lastKey": "event.key"}),
			want:  `hx-vals='js:{lastKey:event.key}'`,
		},
		{
			name:  "ValsJS with invalid identifier",
			attrs: hx.New().ValsJS(map[string]string{"last-key": "event.key"}),
			want:  `hx-vals='js:{"last-key":event.key}'`,
		},
		// Additional Attributes
		{
			name:  "Confirm",
			attrs: hx.New().Confirm("Are you sure?"),
			want:  `hx-confirm='Are you sure?'`,
		},
		{
			name:  "Delete",
			attrs: hx.New().Delete("/example"),
			want:  `hx-delete='/example'`,
		},
		{
			name:  "Disable",
			attrs: hx.New().Disable(),
			want:  `hx-disable`,
		},
		{
			name:  "DisabledElt",
			attrs: hx.New().DisabledElt("#example"),
			want:  `hx-disabled-elt='#example'`,
		},
		{
			name: "DisabledElt closest",
			attrs: hx.New().DisabledElt(
				hx.DisabledEltRelative(hx.DisabledEltClosest, "#example"),
			),
			want: `hx-disabled-elt='closest #example'`,
		},
		{
			name:  "DisabledElt this",
			attrs: hx.New().DisabledElt(hx.DisabledEltThis),
			want:  `hx-disabled-elt='this'`,
		},
		{
			name:  "Disinherit",
			attrs: hx.New().Disinherit(hx.Get, hx.Boost),
			want:  `hx-disinherit='hx-get hx-boost'`,
		},
		{
			name:  "DisinheritAll",
			attrs: hx.New().DisinheritAll(),
			want:  `hx-disinherit='*'`,
		},
		{
			name:  "Encoding",
			attrs: hx.New().Encoding(hx.EncodingMultipart),
			want:  `hx-encoding='multipart/form-data'`,
		},
		{
			name:  "Ext",
			attrs: hx.New().Ext("example-extension"),
			want:  `hx-ext='example-extension'`,
		},
		{
			name:  "ExtIgnore",
			attrs: hx.New().ExtIgnore("example-extension"),
			want:  `hx-ext='ignore:example-extension'`,
		},
		{
			name:  "Headers",
			attrs: hx.New().Headers(map[string]string{"Content-Type": "application/json"}),
			want:  `hx-headers='{"Content-Type":"application/json"}'`,
		},
		{
			name:  "HeadersJS",
			attrs: hx.New().HeadersJS(map[string]string{"Content-Type": "getContentType()"}),
			want:  `hx-headers='js:{"Content-Type":getContentType()}'`,
		},
		{
			name:  "History",
			attrs: hx.New().History(true),
			want:  `hx-history='true'`,
		},
		{
			name:  "History off",
			attrs: hx.New().History(false),
			want:  `hx-history='false'`,
		},
		{
			name:  "HistoryElt",
			attrs: hx.New().HistoryElt(),
			want:  `hx-history-elt`,
		},
		{
			name:  "Include",
			attrs: hx.New().Include("#example"),
			want:  `hx-include='#example'`,
		},
		{
			name:  "Include this",
			attrs: hx.New().Include(hx.IncludeThis),
			want:  `hx-include='this'`,
		},
		{
			name: "Include relative",
			attrs: hx.New().Include(
				hx.IncludeRelative(hx.Closest, "#example"),
			),
			want: `hx-include='closest #example'`,
		},
		{
			name:  "Indicator",
			attrs: hx.New().Indicator("#example"),
			want:  `hx-indicator='#example'`,
		},
		{
			name: "Indicator relative",
			attrs: hx.New().Indicator(
				hx.IndicatorRelative(hx.IndicatorClosest, "#example"),
			),
			want: `hx-indicator='closest #example'`,
		},
		{
			name:  "ParamsAll",
			attrs: hx.New().ParamsAll(),
			want:  `hx-params='*'`,
		},
		{
			name:  "ParamsNone",
			attrs: hx.New().ParamsNone(),
			want:  `hx-params='none'`,
		},
		{
			name:  "Params",
			attrs: hx.New().Params("one", "two"),
			want:  `hx-params='one,two'`,
		},
		{
			name:  "ParamsNot",
			attrs: hx.New().ParamsNot("one", "two"),
			want:  `hx-params='not one,two'`,
		},
		{
			name:  "Patch",
			attrs: hx.New().Patch("/example"),
			want:  `hx-patch='/example'`,
		},
		{
			name:  "Preserve",
			attrs: hx.New().Preserve(),
			want:  `hx-preserve`,
		},
		{
			name:  "Prompt",
			attrs: hx.New().Prompt("Enter a value"),
			want:  `hx-prompt='Enter a value'`,
		},
		{
			name:  "Put",
			attrs: hx.New().Put("/example"),
			want:  `hx-put='/example'`,
		},
		{
			name:  "ReplaceURL",
			attrs: hx.New().ReplaceURL(true),
			want:  `hx-replace-url='true'`,
		},
		{
			name:  "ReplaceURLWith",
			attrs: hx.New().ReplaceURLWith("/example"),
			want:  `hx-replace-url='/example'`,
		},
		{
			name:  "Sync",
			attrs: hx.New().Sync(hx.SyncThis),
			want:  `hx-sync='this'`,
		},
		{
			name:  "SyncStrategy",
			attrs: hx.New().SyncStrategy(hx.SyncThis, hx.SyncDrop),
			want:  `hx-sync='this:drop'`,
		},
		{
			name: "SyncStrategy relative",
			attrs: hx.New().SyncStrategy(
				hx.SyncRelative(hx.Closest, "#example"),
				hx.SyncDrop,
			),
			want: `hx-sync='closest #example:drop'`,
		},
		{
			name:  "Validate",
			attrs: hx.New().Validate(true),
			want:  `hx-validate='true'`,
		},
		{
			name:  "More",
			attrs: hx.New().More(map[string]any{"method": "GET", "action": "/page"}),
			want:  `action='/page' method='GET'`,
		},
		{
			name:  "Unset",
			attrs: hx.New().Unset(hx.Boost, hx.Get),
			want:  `hx-boost='unset' hx-get='unset'`,
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
