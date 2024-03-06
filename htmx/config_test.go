package htmx_test

import (
	"fmt"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx/hxconfig"
	"github.com/will-wow/typed-htmx-go/htmx/swap"
)

// ExampleHX_Config shows the default values, which can be omitted in normal use.
func ExampleHX_Config() {
	config := hx.Config(hxconfig.New().
		HistoryEnabled(true).
		HistoryCacheSize(10).
		RefreshOnHistoryMiss(false).
		DefaultSwapStyle(swap.InnerHTML).
		DefaultSwapDelay(0).
		DefaultSettleDelay(20 * time.Millisecond).
		IncludeIndicatorStyles(true).
		IndicatorClass("htmx-indicator").
		RequestClass("htmx-request").
		AddedClass("htmx-added").
		SettlingClass("htmx-settling").
		SwappingClass("htmx-swapping").
		AllowEval(true).
		AllowScriptTags(true).
		InlineScriptNonce("nonce").
		AttributesToSettle([]string{"foo"}).
		UseTemplateFragments(false).
		WSReconnectDelay("full-jitter").
		WSBinaryType("blob").
		DisableSelector("[hx-disable], [data-hx-disable]").
		WithCredentials(false).
		Timeout(time.Second). // Default is 0, this is an example
		ScrollBehavior(hxconfig.ScrollBehaviorSmooth).
		DefaultFocusScroll(false).
		GetCacheBusterParam(false).
		GlobalViewTransitions(false).
		MethodsThatUseUrlParams([]hxconfig.HTTPMethod{hxconfig.MethodGet}).
		SelfRequestsOnly(false).
		IgnoreTitle(false).
		ScrollIntoViewOnBoost(true).
		TriggerSpecsCache("cacheObject"), // Default is nil, this is an example
	)

	fmt.Println(config.String())

	// output: content='{"addedClass":"htmx-added","allowEval":true,"allowScriptTags":true,"attributesToSettle":["foo"],"defaultFocusScroll":false,"defaultSettleDelay":20,"defaultSwapDelay":0,"defaultSwapStyle":"innerHTML","disableSelector":"[hx-disable], [data-hx-disable]","getCacheBusterParam":false,"globalViewTransitions":false,"historyCacheSize":10,"historyEnabled":true,"ignoreTitle":false,"includeIndicatorStyles":true,"indicatorClass":"htmx-indicator","inlineScriptNonce":"nonce","methodsThatUseUrlParams":["get"],"refreshOnHistoryMiss":false,"requestClass":"htmx-request","scrollBehavior":"smooth","scrollIntoViewOnBoost":true,"selfRequestsOnly":false,"settlingClass":"htmx-settling","swappingClass":"htmx-swapping","timeout":1000,"triggerSpecsCache":"cacheObject","useTemplateFragments":false,"withCredentials":false,"wsBinaryType":"blob","wsReconnectDelay":"full-jitter"}'
}
