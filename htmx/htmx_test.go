package htmx_test

import (
	"fmt"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/on"
	"github.com/will-wow/typed-htmx-go/htmx/swap"
	"github.com/will-wow/typed-htmx-go/htmx/trigger"
)

var hx = htmx.NewStringAttrs()

func ExampleHX_Boost() {
	fmt.Println(hx.Boost(true))
	// Output: hx-boost='true'
}

func ExampleHX_Get() {
	fmt.Println(hx.Get("/example"))
	// Output: hx-get='/example'
}

func ExampleHX_Get_format() {
	fmt.Println(hx.Get("/example/%d", 1))
	// Output: hx-get='/example/1'
}

func ExampleHX_Post() {
	fmt.Println(hx.Post("/example"))
	// Output: hx-post='/example'
}

func ExampleHX_Post_format() {
	fmt.Println(hx.Post("/example/%d", 1))
	// Output: hx-post='/example/1'
}

func ExampleHX_On() {
	fmt.Println(hx.On("click", `alert("clicked")`))
	// Output: hx-on:click='alert("clicked")'
}

func ExampleHX_On_htmxEvent() {
	fmt.Println(hx.On(on.BeforeRequest, `alert("before")`))
	// Output: hx-on:htmx:before-request='alert("before")'
}

func ExampleHX_PushURL() {
	fmt.Println(hx.PushURL(true))
	// Output: hx-push-url='true'
}

func ExampleHX_PushURLPath() {
	fmt.Println(hx.PushURLPath("/example"))
	// Output: hx-push-url='/example'
}

func ExampleHX_PushURLPath_format() {
	fmt.Println(hx.PushURLPath("/example/%d", 1))
	// Output: hx-push-url='/example/1'
}

func ExampleHX_Select() {
	fmt.Println(hx.Select("#example"))
	// Output: hx-select='#example'
}

func ExampleHX_SelectOOB() {
	fmt.Println(hx.SelectOOB("#info-details", "#other-details"))
	// Output: hx-select-oob='#info-details,#other-details'
}

func ExampleHX_SelectOOBWithStrategy() {
	fmt.Println(hx.SelectOOBWithStrategy(
		htmx.SelectOOBStrategy{Selector: "#info-details", Strategy: swap.AfterBegin},
		htmx.SelectOOBStrategy{Selector: "#other-details", Strategy: ""},
	))
	// Output: hx-select-oob='#info-details:afterbegin,#other-details'
}

func ExampleHX_Swap() {
	fmt.Println(hx.Swap(swap.OuterHTML))
	// Output: hx-swap='outerHTML'
}

func ExampleHX_SwapExtended() {
	fmt.Println(hx.SwapExtended(
		swap.New().
			Strategy(swap.OuterHTML).
			Settle(time.Second).
			ShowElement("#example", swap.Top),
	))
	// Output: hx-swap='outerHTML settle:1s show:#example:top'
}

func ExampleHX_SwapOOB() {
	fmt.Println(hx.SwapOOB())
	// Output: hx-swap-oob='true'
}

func ExampleHX_SwapOOBWithStrategy() {
	fmt.Println(hx.SwapOOBWithStrategy(swap.AfterBegin))
	// Output: hx-swap-oob='afterbegin'
}

func ExampleHX_SwapOOBSelector() {
	fmt.Println(hx.SwapOOBSelector(swap.AfterBegin, "#example"))
	// Output: hx-swap-oob='afterbegin:#example'
}

func ExampleHX_Target() {
	fmt.Println(hx.Target("#example"))
	// Output: hx-target='#example'
}

func ExampleHX_Target_nonStandard() {
	fmt.Println(hx.Target(htmx.TargetThis))
	// Output: hx-target='this'
}

func ExampleHX_Target_relativeSelector() {
	fmt.Println(hx.Target(
		htmx.TargetRelative(htmx.Closest, "#example"),
	))
	// Output: hx-target='closest #example'
}

func ExampleHX_Trigger() {
	fmt.Println(hx.Trigger("click"))
	// Output: hx-trigger='click'
}

func ExampleHX_Trigger_nonStandard() {
	fmt.Println(hx.Trigger(trigger.Load))
	// Output: hx-trigger='load'
}

func ExampleHX_TriggerExtended() {
	fmt.Println(hx.TriggerExtended(
		trigger.On("click").When("ctrlKey").Target("#element"),
		trigger.Every(time.Second),
		trigger.Intersect().Root("#element").Threshold(0.2),
	))
	// Output: hx-trigger='click[ctrlKey] target:#element, every 1s, intersect root:#element threshold:0.2'
}

func ExampleHX_Vals() {
	fmt.Println(hx.Vals(map[string]int{"one": 1, "two": 2}))
	// Output: hx-vals='{"one":1,"two":2}'
}

func ExampleHX_Vals_error() {
	fmt.Println(hx.Vals(func() {}))
	// Output: hx-vals='{}'
}

func ExampleHX_Vals_invalid() {
	// You would expect this to be an error, but `Vals` doesn't check the type of the value for performance reasons.
	fmt.Println(hx.Vals(0))
	// Output: hx-vals='0'
}

func ExampleHX_ValsJS() {
	fmt.Println(hx.ValsJS(map[string]string{"lastKey": "event.key"}))
	// Output: hx-vals='js:{lastKey:event.key}'
}

func ExampleHX_ValsJS_withInvalidIdentifier() {
	fmt.Println(hx.ValsJS(map[string]string{"last-key": "event.key"}))
	// Output: hx-vals='js:{"last-key":event.key}'
}

func ExampleHX_Confirm() {
	fmt.Println(hx.Confirm("Are you sure?"))
	// Output: hx-confirm='Are you sure?'
}

func ExampleHX_Delete() {
	fmt.Println(hx.Delete("/example"))
	// Output: hx-delete='/example'
}

func ExampleHX_Delete_format() {
	fmt.Println(hx.Delete("/example/%d", 1))
	// Output: hx-delete='/example/1'
}

func ExampleHX_Disable() {
	fmt.Println(hx.Disable())
	// Output: hx-disable
}

func ExampleHX_DisabledElt() {
	fmt.Println(hx.DisabledElt("#example"))
	// Output: hx-disabled-elt='#example'
}

func ExampleHX_DisabledElt_relative() {
	fmt.Println(hx.DisabledElt(
		htmx.DisabledEltRelative(htmx.DisabledEltClosest, "#example"),
	))
	// Output: hx-disabled-elt='closest #example'
}

func ExampleHX_DisabledElt_this() {
	fmt.Println(hx.DisabledElt(htmx.DisabledEltThis))
	// Output: hx-disabled-elt='this'
}

func ExampleHX_Disinherit() {
	fmt.Println(hx.Disinherit(htmx.Get, htmx.Boost))
	// Output: hx-disinherit='hx-get hx-boost'
}

func ExampleHX_DisinheritAll() {
	fmt.Println(hx.DisinheritAll())
	// Output: hx-disinherit='*'
}

func ExampleHX_Encoding() {
	fmt.Println(hx.Encoding(htmx.EncodingMultipart))
	// Output: hx-encoding='multipart/form-data'
}

func ExampleHX_Ext() {
	fmt.Println(hx.Ext("example-extension"))
	// Output: hx-ext='example-extension'
}

func ExampleHX_ExtIgnore() {
	fmt.Println(hx.ExtIgnore("example-extension"))
	// Output: hx-ext='ignore:example-extension'
}

func ExampleHX_Headers() {
	fmt.Println(hx.Headers(map[string]string{"Content-Type": "application/json"}))
	// Output: hx-headers='{"Content-Type":"application/json"}'
}

func ExampleHX_Headers_error() {
	fmt.Println(hx.Headers(func() {}))
	// Output: hx-headers='{}'
}

func ExampleHX_HeadersJS() {
	fmt.Println(hx.HeadersJS(map[string]string{"Content-Type": "getContentType()"}))
	// Output: hx-headers='js:{"Content-Type":getContentType()}'
}

func ExampleHX_History() {
	fmt.Println(hx.History(true))
	// Output: hx-history='true'
}

func ExampleHX_History_off() {
	fmt.Println(hx.History(false))
	// Output: hx-history='false'
}

func ExampleHX_HistoryElt() {
	fmt.Println(hx.HistoryElt())
	// Output: hx-history-elt
}

func ExampleHX_Include() {
	fmt.Println(hx.Include("#example"))
	// Output: hx-include='#example'
}

func ExampleHX_Include_this() {
	fmt.Println(hx.Include(htmx.IncludeThis))
	// Output: hx-include='this'
}

func ExampleHX_Include_relative() {
	fmt.Println(hx.Include(
		htmx.IncludeRelative(htmx.Closest, "#example"),
	))
	// Output: hx-include='closest #example'
}

func ExampleHX_Indicator() {
	fmt.Println(hx.Indicator("#example"))
	// Output: hx-indicator='#example'
}

func ExampleHX_Indicator_relative() {
	fmt.Println(hx.Indicator(
		htmx.IndicatorRelative(htmx.IndicatorClosest, "#example"),
	))
	// Output: hx-indicator='closest #example'
}

func ExampleHX_ParamsAll() {
	fmt.Println(hx.ParamsAll())
	// Output: hx-params='*'
}

func ExampleHX_ParamsNone() {
	fmt.Println(hx.ParamsNone())
	// Output: hx-params='none'
}

func ExampleHX_Params() {
	fmt.Println(hx.Params("one", "two"))
	// Output: hx-params='one,two'
}

func ExampleHX_ParamsNot() {
	fmt.Println(hx.ParamsNot("one", "two"))
	// Output: hx-params='not one,two'
}

func ExampleHX_Patch() {
	fmt.Println(hx.Patch("/example"))
	// Output: hx-patch='/example'
}

func ExampleHX_Patch_format() {
	fmt.Println(hx.Patch("/example/%d", 1))
	// Output: hx-patch='/example/1'
}

func ExampleHX_Preserve() {
	fmt.Println(hx.Preserve())
	// Output: hx-preserve
}

func ExampleHX_Prompt() {
	fmt.Println(hx.Prompt("Enter a value"))
	// Output: hx-prompt='Enter a value'
}

func ExampleHX_Put() {
	fmt.Println(hx.Put("/example"))
	// Output: hx-put='/example'
}

func ExampleHX_Put_format() {
	fmt.Println(hx.Put("/example/%d", 1))
	// Output: hx-put='/example/1'
}

func ExampleHX_ReplaceURL() {
	fmt.Println(hx.ReplaceURL(true))
	// Output: hx-replace-url='true'
}

func ExampleHX_ReplaceURLWith() {
	fmt.Println(hx.ReplaceURLWith("/example"))
	// Output: hx-replace-url='/example'
}

func ExampleHX_ReplaceURLWith_format() {
	fmt.Println(hx.ReplaceURLWith("/example/%d", 1))
	// Output: hx-replace-url='/example/1'
}

func ExampleHX_Request() {
	fmt.Println(hx.Request(htmx.RequestConfig{
		Timeout:     time.Second,
		Credentials: true,
		NoHeaders:   true,
	}))
	// Output: hx-request='"timeout":1000,"credentials":true,"noHeaders":true'
}

func ExampleHX_RequestJS() {
	fmt.Println(hx.RequestJS(htmx.RequestConfigJS{
		Timeout:     "getTimeoutSetting()",
		Credentials: "true",
		NoHeaders:   "noHeaders()",
	}))
	// Output: hx-request='js: timeout:getTimeoutSetting(),credentials:true,noHeaders:noHeaders()'
}

func ExampleHX_Sync() {
	fmt.Println(hx.Sync(htmx.SyncThis))
	// Output: hx-sync='this'
}

func ExampleHX_SyncStrategy() {
	fmt.Println(hx.SyncStrategy(htmx.SyncThis, htmx.SyncDrop))
	// Output: hx-sync='this:drop'
}

func ExampleHX_SyncStrategy_relative() {
	fmt.Println(hx.SyncStrategy(
		htmx.SyncRelative(htmx.Closest, "#example"),
		htmx.SyncDrop,
	))
	// Output: hx-sync='closest #example:drop'
}

func ExampleHX_Validate() {
	fmt.Println(hx.Validate(true))
	// Output: hx-validate='true'
}

func ExampleHX_Unset() {
	fmt.Println(hx.Unset(htmx.Boost))
	// Output: hx-boost='unset'
}
