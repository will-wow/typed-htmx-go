package htmx

import (
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/will-wow/typed-htmx-go/swap"
)

type Builder struct {
	attrs map[string]any
}

// HX starts a new HTMX attributes builder.
func HX() Builder {
	return Builder{
		attrs: map[string]any{},
	}
}

// Build returns the final attribute map, compatible with [templ.Attributes].
func (b Builder) Build() map[string]any {
	return b.attrs
}

// String renders the attributes as HTML attributes.
func (b Builder) String() string {
	attributes := make([]string, len(b.attrs))

	i := 0
	for k, v := range b.attrs {
		attributes[i] = fmt.Sprintf(`%s='%v'`, k, v)
		i++
	}

	// Do a stable sort, which makes testing easier.
	if len(attributes) > 1 {
		sort.StringSlice(attributes).Sort()
		slices.SortStableFunc(attributes, func(a string, b string) int {
			return strings.Compare(a, b)
		})
	}

	return strings.Join(attributes, " ")
}

// Boost  allows you to “boost” normal anchors and form tags to use AJAX instead. This has the [nice fallback] that, if the user does not have javascript enabled, the site will continue to work.
//
// For anchor tags, clicking on the anchor will issue a GET request to the url specified in the href and will push the url so that a history entry is created. The target is the <body> tag, and the innerHTML swap strategy is used by default. All of these can be modified by using the appropriate attributes, except the click trigger.
//
// For forms the request will be converted into a GET or POST, based on the method in the method attribute and will be triggered by a submit. Again, the target will be the body of the page, and the innerHTML swap will be used. The url will not be pushed, however, and no history entry will be created. (You can use the [hx.PushUrl] attribute if you want the url to be pushed.)
//
//	<div { htmx.HX().Boost(true).Build()...} >
//		<a href="/page1">Go To Page 1</a>
//		<a href="/page2">Go To Page 2</a>
//	</div>
//
// These links will issue an ajax GET request to the respective URLs and replace the body’s inner content with it.
//
// Here is an example of a boosted form:
//
//	<form {...htmx.HX().Boost(true).Build()...} action="/example" method="post">
//		<input name="email" type="email" placeholder="Enter email...">
//		<button>Submit</button>
//	</form>
//
// This form will issue an ajax POST to the given URL and replace the body’s inner content with it.
//
// # Notes
//
//   - hx-boost is inherited and can be placed on a parent element
//   - Only links that are to the same domain and that are not local anchors will be boosted
//   - All requests are done via AJAX, so keep that in mind when doing things like redirects
//   - To find out if the request results from a boosted anchor or form, look for HX-Boosted in the request header
//   - Selectively disable boost on child elements with hx-boost="false"
//   - Disable the replacement of elements via boost, and their children, with hx-preserve="true"
//
// HTMX Attribute: [hx-boost]
//
// [hx-boost]: https://htmx.org/attributes/hx-boost/
// [nice fallback]: https://en.wikipedia.org/wiki/Progressive_enhancement
func (b Builder) Boost(boost bool) Builder {
	if boost {
		b.attrs["hx-boost"] = "true"
	} else {
		b.attrs["hx-boost"] = "false"
	}
	return b
}

// Get will cause an element to issue a GET to the specified URL and swap the HTML into the DOM using a swap strategy.
//
//	<div {htmx.HX().Get("/example").Build()...}>Get Some HTML</div>
//
// This example will cause the div to issue a GET to /example and swap the returned HTML into the innerHTML of the div.
//
// Notes:
//   - hx-get is not inherited
//   - By default hx-get does not include any parameters. You can use the hx-params attribute to change this
//   - You can control the target of the swap using the hx-target attribute
//   - You can control the swap strategy by using the hx-swap attribute
//   - You can control what event triggers the request with the hx-trigger attribute
//   - You can control the data submitted with the request in various ways, documented here: [Parameters]
//
// HTMX Attribute: [hx-get]
//
// [hx-get]: https://htmx.org/attributes/hx-get/
// [Parameters]: https://htmx.org/docs/#parameters
func (b Builder) Get(url string) Builder {
	b.attrs["hx-get"] = url
	return b
}

// Post will cause an element to issue a POST to the specified URL and swap the HTML into the DOM using a swap strategy.
//
//	<button {htmx.HX().Post("/accounts/enable").Target("body").Build()...}>
//	  Enable Your Account
//	</button>
//
// This example will cause the button to issue a POST to /account/enable and swap the returned HTML into the innerHTML of the body.
//
// Notes
//   - hx-post is not inherited
//   - You can control the target of the swap using the hx-target attribute
//   - You can control the swap strategy by using the hx-swap attribute
//   - You can control what event triggers the request with the hx-trigger attribute
//   - You can control the data submitted with the request in various ways, documented here: [Parameters]
//
// HTMX Attribute: [hx-post]
//
// [hx-post]: https://htmx.org/attributes/hx-post/
// [Parameters]: https://htmx.org/docs/#parameters
func (b Builder) Post(url string) Builder {
	b.attrs["hx-post"] = url
	return b
}

// On allows you to embed scripts inline to respond to events directly on an element; similar to the onevent properties found in HTML, such as onClick.
//
// The hx-on* attributes improve upon onevent by enabling the handling of any arbitrary JavaScript event, for enhanced Locality of Behaviour (LoB) even when dealing with non-standard DOM events. For example, these attributes allow you to handle htmx events.
//
// HX().On() attaches to standard DOM events. For htmx custom events, use [Builder.OnHTMX].
//
// If you wish to handle multiple different events, you can simply add multiple attributes to an element.
//
// # Symbols
//
// Like onevent, two symbols are made available to event handler scripts:
//
//   - this - The element on which the hx-on attribute is defined
//   - event - The event that triggered the handler
//
// # Notes
//
// hx-on is not inherited, however due to event bubbling, hx-on attributes on parent elements will typically be triggered by events on child elements.
//
// HTMX Attribute: [hx-on]
//
// [hx-on]: https://htmx.org/attributes/hx-on/
func (b Builder) On(event string, action string) Builder {
	b.attrs[fmt.Sprintf("hx-on:%s", event)] = action
	return b
}

// OnHTMX allows you to embed scripts inline to respond to HTMX events directly on an element; similar to the onevent properties found in HTML, such as onClick.
//
// The hx-on* attributes improve upon onevent by enabling the handling of any arbitrary JavaScript event, for enhanced Locality of Behaviour (LoB) even when dealing with non-standard DOM events. For example, these attributes allow you to handle htmx events.
//
// All htmx and other custom events can be captured, too! To respond to standard DOM events, use [Builder.On] instead.
//
// One gotcha to note is that DOM attributes do not preserve case. This means, unfortunately, an attribute like hx-on:htmx:beforeRequest will not work, because the DOM lowercases the attribute names. Fortunately, htmx supports both camel case event names and also kebab-case event names, so you can use .OnHTMX("before-request") instead.
//
//	<button {htmx.HX().OnHTMX("before-request", "alert('making a request!')").Get("/info").Build()...}} >
//	Get Info!
//	</button>
//
// If you wish to handle multiple different events, you can simply add multiple attributes to an element.
//
// # Symbols
//
// Like onevent, two symbols are made available to event handler scripts:
//
//   - this - The element on which the hx-on attribute is defined
//   - event - The event that triggered the handler
//
// # Notes
//
// hx-on is not inherited, however due to event bubbling, hx-on attributes on parent elements will typically be triggered by events on child elements.
//
// HTMX Attribute: [hx-on]
//
// [hx-on]: https://htmx.org/attributes/hx-on/
func (b Builder) OnHTMX(event string, action string) Builder {
	b.attrs[fmt.Sprintf("hx-on::%s", event)] = action
	return b
}

// PushURL allows you to push a URL into the browser location history. This creates a new history entry, allowing navigation with the browser’s back and forward buttons. htmx snapshots the current DOM and saves it into its history cache, and restores from this cache on navigation.
//
// The possible values of this attribute are:
//   - true, which pushes the fetched URL into history.
//   - false, which disables pushing the fetched URL if it would otherwise be pushed due to inheritance or hx-boost.
//
// To push a specific URL into history, use [attributes.PushURLPath].
//
// # Example
//
//	<div {htmx.HX().Get("/account").PushURL(true).Build()...}>
//		Go to My Account
//	</div>
//
// # Notes
//
//   - hx-push-url is inherited and can be placed on a parent element
//   - The HX-Push-Url response header has similar behavior and can override this attribute.
//   - The hx-history-elt attribute allows changing which element is saved in the history cache.
//
// HTMX Attribute: [hx-push-url]
//
// [hx-push-url]: https://htmx.org/attributes/hx-push-url/
func (b Builder) PushURL(on bool) Builder {
	b.attrs["hx-push-url"] = boolToString(on)
	return b
}

// PushURLPath allows you to push a URL into the browser location history. This creates a new history entry, allowing navigation with the browser’s back and forward buttons. htmx snapshots the current DOM and saves it into its history cache, and restores from this cache on navigation.
//
// This method takes a URL to be pushed into the location bar. This may be relative or absolute, as per history.pushState().
//
// To simply toggle pushing the URL associated with a link, use [attributes.PushURL].
//
// # Example
//
//	<div {htmx.HX().Get("/account").PushURLPath("/account/home").Build()...}>
//		Go to My Account
//	</div>
//
// # Notes
//
//   - hx-push-url is inherited and can be placed on a parent element
//   - The HX-Push-Url response header has similar behavior and can override this attribute.
//   - The hx-history-elt attribute allows changing which element is saved in the history cache.
//
// HTMX Attribute: [hx-push-url]
//
// [hx-push-url]: https://htmx.org/attributes/hx-push-url/
func (b Builder) PushURLPath(url string) Builder {
	b.attrs["hx-push-url"] = url
	return b
}

// Select allows you to select the content you want swapped from a response. The value of this attribute is a CSS query selector of the element or elements to select from the response.
//
// Here is an example that selects a subset of the response content:
//
//	<div>
//		<button { htmx.HX().Get("/info").Select("#info-details").Swap(swap.outerHTML).Build()... } >
//			Get Info!
//		</button>
//	</div>
//
// So this button will issue a GET to /info and then select the element with the id info-detail, which will replace the entire button in the DOM.
//
// # Notes
//
// hx-select is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-select]
//
// [hx-select]: https://htmx.org/attributes/hx-select/
func (b Builder) Select(selector string) Builder {
	b.attrs["hx-select"] = selector
	return b
}

func (b Builder) SelectOOB(selector string) Builder {
	b.attrs["hx-select-oob"] = selector
	return b
}

func (b Builder) Swap(style swap.Style) Builder {
	b.attrs["hx-swap"] = string(style)
	return b
}

func (b Builder) SwapExtended(swap *swap.Builder) Builder {
	b.attrs["hx-swap"] = swap.String()
	return b
}

func (b Builder) SwapOOB(selector string) Builder {
	b.attrs["hx-swap-oob"] = selector
	return b
}

func (b Builder) Target(selector string) Builder {
	b.attrs["hx-target"] = selector
	return b
}

type TargetElementType string

const (
	TargetElementThis     TargetElementType = "this"
	TargetElementNext     TargetElementType = "next"
	TargetElementPrevious TargetElementType = "previous"
)

func (b Builder) TargetElement(target TargetElementType) Builder {
	b.attrs["hx-target"] = string(target)
	return b
}

type TargetSelectorType string

const (
	TargetSelectorClosest  TargetSelectorType = "closest"
	TargetSelectorFind     TargetSelectorType = "find"
	TargetSelectorNext     TargetSelectorType = "next"
	TargetSelectorPrevious TargetSelectorType = "previous"
)

func (b Builder) TargetSelector(targetType TargetSelectorType, selector string) Builder {
	b.attrs["hx-target"] = fmt.Sprintf("%s %s", targetType, selector)
	return b
}

func (b Builder) Trigger(event string) Builder {
	b.attrs["hx-trigger"] = event
	return b
}

func (b Builder) Vals(vals any) Builder {
	json, err := json.Marshal(vals)
	if err != nil {
		// Silently ignore the value if there is an error, because there's not a good way to report an error when constructing templ attributes.
		return b
	}
	b.attrs["hx-vals"] = string(json)
	return b
}

func (b Builder) ValsJS(vals any) Builder {
	json, err := json.Marshal(vals)
	if err != nil {
		// Silently ignore the value if there is an error, because there's not a good way to report an error when constructing templ attributes.
		return b
	}
	b.attrs["hx-vals"] = fmt.Sprintf("js:%s", json)
	return b
}

// Additional Attributes

func (b Builder) Confirm(msg string) Builder {
	b.attrs["hx-confirm"] = msg
	return b
}
func (b Builder) Delete(url string) Builder {
	b.attrs["hx-delete"] = url
	return b
}

func (b Builder) Disable() Builder {
	b.attrs["hx-disable"] = true
	return b
}

func (b Builder) DisabledElt(selector string) Builder {
	b.attrs["hx-disabled-elt"] = selector
	return b
}

// TODO: Typed disinherit https://htmx.org/attributes/hx-disinherit/
func (b Builder) Disinherit(attr string) Builder {
	b.attrs["hx-disinherit"] = attr
	return b
}

type Encoding string

const (
	EncodingMultipart Encoding = "multipart/form-data"
)

func (b Builder) Encoding(encoding Encoding) Builder {
	b.attrs["hx-encoding"] = encoding
	return b
}
func (b Builder) Ext(ext string) Builder {
	b.attrs["hx-ext"] = ext
	return b
}

func (b Builder) Headers(headers any) Builder {
	json, err := json.Marshal(headers)
	if err != nil {
		// Silently ignore the value if there is an error, because there's not a good way to report an error when constructing templ attributes.
		return b
	}
	b.attrs["hx-headers"] = fmt.Sprintf("js:%b", json)
	return b
}

func (b Builder) HeadersJS(headers any) Builder {
	json, err := json.Marshal(headers)
	if err != nil {
		// Silently ignore the value if there is an error, because there's not a good way to report an error when constructing templ attributes.
		return b
	}
	b.attrs["hx-headers"] = string(json)
	return b
}

func (b Builder) History(on bool) Builder {
	b.attrs["hx-history"] = boolToString(on)
	return b
}

// HistoryElt allows you to specify the element that will be used to snapshot and restore page state during navigation. By default, the body tag is used. This is typically good enough for most setups, but you may want to narrow it down to a child element. Just make sure that the element is always visible in your application, or htmx will not be able to restore history navigation properly.
//
// Here is an example:
//
//	<html>
//	<body>
//	<div id="content" {htmx.HX().HistoryElt().Build()...}>
//	 ...
//	</div>
//	</body>
//	</html>
//
// # Notes
//
// - hx-history-elt is not inherited
// - In most cases we don’t recommend narrowing the history snapshot
//
// HTMX Attribute: [hx-history-elt]
//
// [hx-history-elt]: https://htmx.org/attributes/hx-history-elt/
func (b Builder) HistoryElt() Builder {
	b.attrs["hx-history-elt"] = true
	return b
}

// include additional data in requests
func (b Builder) Include(selector string) Builder {
	b.attrs["hx-include"] = selector
	return b
}

// include additional data in requests
func (b Builder) IncludeThis() Builder {
	b.attrs["hx-include"] = "this"
	return b
}

type IncludeExtension string

const (
	IncludeClosest  IncludeExtension = "closest"
	IncludeFind     IncludeExtension = "find"
	IncludeNext     IncludeExtension = "next"
	IncludePrevious IncludeExtension = "previous"
)

// include additional data in requests
func (b Builder) IncludeExtended(extension IncludeExtension, selector string) Builder {
	b.attrs["hx-include"] = fmt.Sprintf("%s %s", extension, selector)
	return b
}

func (b Builder) Indicator(selector string) Builder {
	b.attrs["hx-indicator"] = selector
	return b
}
func (b Builder) IndicatorClosest(selector string) Builder {
	b.attrs["hx-indicator"] = fmt.Sprintf("closest %s", selector)
	return b
}

func (b Builder) ParamsAll() Builder {
	b.attrs["hx-params"] = "*"
	return b
}

func (b Builder) ParamsNone() Builder {
	b.attrs["hx-params"] = "none"
	return b
}

func (b Builder) Params(params []string) Builder {
	b.attrs["hx-params"] = strings.Join(params, ",")
	return b
}

func (b Builder) ParamsNot(params []string) Builder {
	b.attrs["hx-params"] = fmt.Sprintf("not %s", strings.Join(params, ","))
	return b
}

func (b Builder) Patch(url string) Builder {
	b.attrs["hx-patch"] = url
	return b
}

func (b Builder) Preserve() Builder {
	b.attrs["hx-preserve"] = true
	return b
}

func (b Builder) Prompt(msg string) Builder {
	b.attrs["hx-prompt"] = msg
	return b
}

func (b Builder) Put(url string) Builder {
	b.attrs["hx-put"] = url
	return b
}

func (b Builder) ReplaceURL(on bool) Builder {
	b.attrs["hx-replace-url"] = boolToString(on)
	return b
}

func (b Builder) ReplaceURLWith(url string) Builder {
	b.attrs["hx-replace-url"] = url
	return b
}

// Request describes the hx-request attributes
// See https://htmx.org/attributes/hx-request/
type Request struct {
	TimeoutMS   int  // the timeout for the request in milliseconds
	Credentials bool // if the request will send credentials
	NoHeaders   bool // strips all headers from the request
	JS          bool // You may make the values dynamically evaluated by adding this prefix.
}

func (r Request) String() string {
	opts := []string{}

	if r.TimeoutMS > 0 {
		opts = append(opts, fmt.Sprintf(`"timeout":%d`, r.TimeoutMS))
	}
	if r.Credentials {
		opts = append(opts, `"credentials": true`)
	}
	if r.NoHeaders {
		opts = append(opts, `"noHeaders": true`)
	}

	value := strings.Join(opts, ",")

	if r.JS {
		return fmt.Sprintf("js: %s", value)
	} else {
		return value
	}
}

func (b Builder) Request(request Request) Builder {
	b.attrs["hx-request"] = request.String()
	return b
}

type SyncStrategy string

const (
	SyncDefault    SyncStrategy = ""
	SyncDrop       SyncStrategy = "drop"        // drop (ignore) this request if an existing request is in flight (the default)
	SyncAbort      SyncStrategy = "abort"       // drop (ignore) this request if an existing request is in flight, and, if that is not the case, abort this request if another request occurs while it is still in flight
	SyncReplace    SyncStrategy = "replace"     // abort the current request, if any, and replace it with this request
	SyncQueue      SyncStrategy = "queue"       // place this request in the request queue associated with the given element
	SyncQueueFirst SyncStrategy = "queue first" // queue the first request to show up while a request is in flight
	SyncQueueLast  SyncStrategy = "queue last"  // queue the last request to show up while a request is in flight
	SyncQueueAll   SyncStrategy = "queue all"   // queue all requests that show up while a request is in flight
)

func (b Builder) Sync(selector string) Builder {
	b.attrs["hx-sync"] = selector
	return b
}

func (b Builder) SyncStrategy(selector string, strategy SyncStrategy) Builder {
	b.attrs["hx-sync"] = fmt.Sprintf("%s:%s", selector, strategy)
	return b
}

func (b Builder) Validate() Builder {
	b.attrs["hx-validate"] = true
	return b
}

// More allow you to merge arbitrary maps into the final attributes.
// This allows additional attributes to be passed down in a single map.
func (b Builder) More(more map[string]any) Builder {
	for k, v := range more {
		b.attrs[k] = v
	}
	return b
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
