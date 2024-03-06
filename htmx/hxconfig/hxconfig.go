package hxconfig

import (
	"time"

	"github.com/will-wow/typed-htmx-go/htmx/swap"
)

type Builder struct {
	config map[string]any
}

func New() *Builder {
	return &Builder{
		config: map[string]any{},
	}
}

// HistoryEnabled defaults to true, really only useful for testing
func (b *Builder) HistoryEnabled(value bool) *Builder {
	b.config["historyEnabled"] = &value
	return b
}

// defaults to 10
func (b *Builder) HistoryCacheSize(value int) *Builder {
	b.config["historyCacheSize"] = &value
	return b
}

// defaults to false, if set to true htmx will issue a full page refresh on history misses rather than use an AJAX request
func (b *Builder) RefreshOnHistoryMiss(value bool) *Builder {
	b.config["refreshOnHistoryMiss"] = &value
	return b
}

// defaults to innerHTML
func (b *Builder) DefaultSwapStyle(value swap.Strategy) *Builder {
	b.config["defaultSwapStyle"] = &value
	return b
}

// defaults to 0
func (b *Builder) DefaultSwapDelay(value time.Duration) *Builder {
	// convert to milliseconds
	ms := int(value / time.Millisecond)
	b.config["defaultSwapDelay"] = &ms
	return b
}

// defaults to 20
func (b *Builder) DefaultSettleDelay(value time.Duration) *Builder {
	// convert to milliseconds
	ms := int(value / time.Millisecond)
	b.config["defaultSettleDelay"] = &ms
	return b
}

// defaults to true (determines if the indicator styles are loaded) )
func (b *Builder) IncludeIndicatorStyles(value bool) *Builder {
	b.config["includeIndicatorStyles"] = &value
	return b
}

// defaults to htmx-indicator
func (b *Builder) IndicatorClass(value string) *Builder {
	b.config["indicatorClass"] = &value
	return b
}

// defaults to htmx-request
func (b *Builder) RequestClass(value string) *Builder {
	b.config["requestClass"] = &value
	return b
}

// defaults to htmx-added
func (b *Builder) AddedClass(value string) *Builder {
	b.config["addedClass"] = &value
	return b
}

// defaults to htmx-settling
func (b *Builder) SettlingClass(value string) *Builder {
	b.config["settlingClass"] = &value
	return b
}

// defaults to htmx-swapping
func (b *Builder) SwappingClass(value string) *Builder {
	b.config["swappingClass"] = &value
	return b
}

// defaults to true, can be used to disable htmx’s use of eval for certain features (e.g. trigger filters)
func (b *Builder) AllowEval(value bool) *Builder {
	b.config["allowEval"] = &value
	return b
}

// defaults to true, determines if htmx will process script tags found in new content
func (b *Builder) AllowScriptTags(value bool) *Builder {
	b.config["allowScriptTags"] = &value
	return b
}

// defaults to ”, meaning that no nonce will be added to inline scripts
func (b *Builder) InlineScriptNonce(value string) *Builder {
	b.config["inlineScriptNonce"] = &value
	return b
}

// defaults to ["class", "style", "width", "height"], the attributes to settle during the settling phase
func (b *Builder) AttributesToSettle(value []string) *Builder {
	b.config["attributesToSettle"] = &value
	return b
}

// defaults to false, HTML template tags for parsing content from the server (not IE11 compatible!)
func (b *Builder) UseTemplateFragments(value bool) *Builder {
	b.config["useTemplateFragments"] = &value
	return b
}

// defaults to ["get"], htmx will format requests with these methods by encoding their parameters in the URL, not the request body
func (b *Builder) WSReconnectDelay(value string) *Builder {
	b.config["wsReconnectDelay"] = &value
	return b
}

// defaults to blob, the the type of binary data being received over the WebSocket connection
func (b *Builder) WSBinaryType(value string) *Builder {
	b.config["wsBinaryType"] = &value
	return b
}

// defaults to [hx-disable], [data-hx-disable], htmx will not process elements with this attribute on it or a parent
func (b *Builder) DisableSelector(value string) *Builder {
	b.config["disableSelector"] = &value
	return b
}

// defaults to false, allow cross-site Access-Control requests using credentials such as cookies, authorization headers or TLS client certificates
func (b *Builder) WithCredentials(value bool) *Builder {
	b.config["withCredentials"] = &value
	return b
}

// defaults to 0, the number of milliseconds a request can take before automatically being terminated
func (b *Builder) Timeout(value time.Duration) *Builder {
	// convert to milliseconds
	ms := int(value / time.Millisecond)
	b.config["timeout"] = &ms
	return b
}

type ScrollBehavior string

const (
	ScrollBehaviorAuto   ScrollBehavior = "auto"   // auto will behave like a vanilla link.
	ScrollBehaviorSmooth ScrollBehavior = "smooth" // smooth (default) will smoothscroll to the top of the page
)

// defaults to ‘smooth’, the behavior for a boosted link on page transitions. The allowed values are auto and smooth. Smooth will smoothscroll to the top of the page while auto will behave like a vanilla link.
func (b *Builder) ScrollBehavior(value ScrollBehavior) *Builder {
	b.config["scrollBehavior"] = &value
	return b
}

// if the focused element should be scrolled into view, defaults to false and can be overridden using the focus-scroll swap modifier.
func (b *Builder) DefaultFocusScroll(value bool) *Builder {
	b.config["defaultFocusScroll"] = &value
	return b
}

// defaults to false, if set to true htmx will include a cache-busting parameter in GET requests to avoid caching partial responses by the browser
func (b *Builder) GetCacheBusterParam(value bool) *Builder {
	b.config["getCacheBusterParam"] = &value
	return b
}

// if set to true, htmx will use the View Transition API when swapping in new content.
func (b *Builder) GlobalViewTransitions(value bool) *Builder {
	b.config["globalViewTransitions"] = &value
	return b
}

// An HTTPMethod is a named HTTP Method used for [Config.MethodsThatUseUrlParams]
type HTTPMethod string

const (
	MethodGet     HTTPMethod = "get"
	MethodPost    HTTPMethod = "post"
	MethodPut     HTTPMethod = "put"
	MethodDelete  HTTPMethod = "delete"
	MethodPatch   HTTPMethod = "patch"
	MethodHead    HTTPMethod = "head"
	MethodOptions HTTPMethod = "options"
)

// defaults to ["get"], htmx will format requests with these methods by encoding their parameters in the URL, not the request body
func (b *Builder) MethodsThatUseUrlParams(value []HTTPMethod) *Builder {
	b.config["methodsThatUseUrlParams"] = &value
	return b
}

// defaults to false, if set to true will only allow AJAX requests to the same domain as the current document
func (b *Builder) SelfRequestsOnly(value bool) *Builder {
	b.config["selfRequestsOnly"] = &value
	return b
}

// defaults to false, if set to true htmx will not update the title of the document when a title tag is found in new content
func (b *Builder) IgnoreTitle(value bool) *Builder {
	b.config["ignoreTitle"] = &value
	return b
}

// defaults to true, whether or not the target of a boosted element is scrolled into the viewport. If hx-target is omitted on a boosted element, the target defaults to body, causing the page to scroll to the top.
func (b *Builder) ScrollIntoViewOnBoost(value bool) *Builder {
	b.config["scrollIntoViewOnBoost"] = &value
	return b
}

// defaults to null, the cache to store evaluated trigger specifications into, improving parsing performance at the cost of more memory usage. You may define a simple object to use a never-clearing cache, or implement your own system using a proxy object
func (b *Builder) TriggerSpecsCache(value string) *Builder {
	b.config["triggerSpecsCache"] = &value
	return b
}

func (b *Builder) Build() map[string]any {
	return b.config
}
