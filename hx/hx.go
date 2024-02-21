// package htmx provides well-documented Go functions for building HTMX attributes.
package hx

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/will-wow/typed-htmx-go/hx/swap"
)

// An HX constructs HTMX attributes.
type HX struct {
	attrs map[string]any
}

// New starts a new HTMX attributes builder.
func New() HX {
	return HX{
		attrs: map[string]any{},
	}
}

// Build returns the final attribute map, compatible with [templ.Attributes].
func (hx HX) Build() map[string]any {
	return hx.attrs
}

// String renders the attributes as HTML attributes.
func (hx HX) String() string {
	attributes := make([]string, len(hx.attrs))

	i := 0
	for k, v := range hx.attrs {
		attributes[i] = fmt.Sprintf(`%s='%v'`, k, v)
		i++
	}

	// Sort, which makes testing easier.
	if len(attributes) > 1 {
		slices.Sort(attributes)
	}

	return strings.Join(attributes, " ")
}

// Boost allows you to “boost” normal anchors and form tags to use AJAX instead. This has the [nice fallback] that, if the user does not have javascript enabled, the site will continue to work.
//
// For anchor tags, clicking on the anchor will issue a GET request to the url specified in the href and will push the url so that a history entry is created. The target is the <body> tag, and the innerHTML swap strategy is used by default. All of these can be modified by using the appropriate attributes, except the click trigger.
//
// For forms the request will be converted into a GET or POST, based on the method in the method attribute and will be triggered by a submit. Again, the target will be the body of the page, and the innerHTML swap will be used. The url will not be pushed, however, and no history entry will be created. (You can use the [hx.PushUrl] attribute if you want the url to be pushed.)
//
//	<div { hx.New().Boost(true).Build()...} >
//		<a href="/page1">Go To Page 1</a>
//		<a href="/page2">Go To Page 2</a>
//	</div>
//
// These links will issue an ajax GET request to the respective URLs and replace the body’s inner content with it.
//
// Here is an example of a boosted form:
//
//	<form {...hx.New().Boost(true).Build()...} action="/example" method="post">
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
func (hx HX) Boost(boost bool) HX {
	if boost {
		hx.attrs["hx-boost"] = "true"
	} else {
		hx.attrs["hx-boost"] = "false"
	}
	return hx
}

// Get will cause an element to issue a GET to the specified URL and swap the HTML into the DOM using a swap strategy.
//
//	<div {hx.New().Get("/example").Build()...}>Get Some HTML</div>
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
func (hx HX) Get(url string) HX {
	hx.attrs["hx-get"] = url
	return hx
}

// Post will cause an element to issue a POST to the specified URL and swap the HTML into the DOM using a swap strategy.
//
//	<button {hx.New().Post("/accounts/enable").Target("body").Build()...}>
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
func (hx HX) Post(url string) HX {
	hx.attrs["hx-post"] = url
	return hx
}

// On allows you to embed scripts inline to respond to events directly on an element; similar to the onevent properties found in HTML, such as onClick.
//
// The hx-on* attributes improve upon onevent by enabling the handling of any arbitrary JavaScript event, for enhanced Locality of Behaviour (LoB) even when dealing with non-standard DOM events. For example, these attributes allow you to handle htmx events.
//
// HX().On() attaches to standard DOM events. For htmx custom events, use [HX.OnHTMX].
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
func (hx HX) On(event string, action string) HX {
	hx.attrs[fmt.Sprintf("hx-on:%s", event)] = action
	return hx
}

// OnHTMX allows you to embed scripts inline to respond to HTMX events directly on an element; similar to the onevent properties found in HTML, such as onClick.
//
// The hx-on* attributes improve upon onevent by enabling the handling of any arbitrary JavaScript event, for enhanced Locality of Behaviour (LoB) even when dealing with non-standard DOM events. For example, these attributes allow you to handle htmx events.
//
// All htmx and other custom events can be captured, too! To respond to standard DOM events, use [HX.On] instead.
//
// One gotcha to note is that DOM attributes do not preserve case. This means, unfortunately, an attribute like hx-on:htmx:beforeRequest will not work, because the DOM lowercases the attribute names. Fortunately, htmx supports both camel case event names and also kebab-case event names, so you can use .OnHTMX("before-request") instead.
//
//	<button {hx.New().OnHTMX("before-request", "alert('making a request!')").Get("/info").Build()...}} >
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
func (hx HX) OnHTMX(event string, action string) HX {
	hx.attrs[fmt.Sprintf("hx-on::%s", event)] = action
	return hx
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
//	<div {hx.New().Get("/account").PushURL(true).Build()...}>
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
func (hx HX) PushURL(on bool) HX {
	hx.attrs["hx-push-url"] = boolToString(on)
	return hx
}

// PushURLPath allows you to push a URL into the browser location history. This creates a new history entry, allowing navigation with the browser’s back and forward buttons. htmx snapshots the current DOM and saves it into its history cache, and restores from this cache on navigation.
//
// This method takes a URL to be pushed into the location bar. This may be relative or absolute, as per history.pushState().
//
// To simply toggle pushing the URL associated with a link, use [attributes.PushURL].
//
// # Example
//
//	<div {hx.New().Get("/account").PushURLPath("/account/home").Build()...}>
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
func (hx HX) PushURLPath(url string) HX {
	hx.attrs["hx-push-url"] = url
	return hx
}

// Select allows you to select the content you want swapped from a response. The value of this attribute is a CSS query selector of the element or elements to select from the response.
//
// Here is an example that selects a subset of the response content:
//
//	<div>
//		<button { hx.New().Get("/info").Select("#info-details").Swap(swap.outerHTML).Build()... } >
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
func (hx HX) Select(selector string) HX {
	hx.attrs["hx-select"] = selector
	return hx
}

// SelectOOB allows you to select content from a response to be swapped in via an out-of-band swap.
// The value of this attribute is comma separated list of elements to be swapped out of band. This attribute is almost always paired with hx-select.
//
// Here is an example that selects a subset of the response content:
//
//	<div>
//	   <div id="alert"></div>
//	    <button
//				{ hx.New().
//					Get("/info").
//					Select("#info-details").
//					Swap(swap.OuterHTML).
//					SelectOOB("#alert").
//					Build()... }
//			>
//	        Get Info!
//	    </button>
//	</div>
//
// This button will issue a GET to /info and then select the element with the id info-details, which will replace the entire button in the DOM, and, in addition, pick out an element with the id alert in the response and swap it in for div in the DOM with the same ID.
//
// Each value in the comma separated list of values can specify any valid hx-swap strategy by separating the selector and the swap strategy with a :.
//
// For example, to prepend the alert content instead of replacing it:
//
//	<div>
//	   <div id="alert"></div>
//	    <button
//				{ hx.New().
//					Get("/info").
//					Select("#info-details").
//					Swap(swap.OuterHTML).
//					SelectOOB("#alert:afterbegin").
//					Build()... }
//			>
//	        Get Info!
//	    </button>
//	</div>
//
// # Notes
//
// hx-select-oob is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-select-oob]
//
// [hx-select-oob]: https://htmx.org/attributes/hx-select-oob/
func (hx HX) SelectOOB(selectors ...string) HX {
	hx.attrs["hx-select-oob"] = strings.Join(selectors, ",")
	return hx
}

type SelectOOBStrategy struct {
	Selector string
	Strategy swap.Strategy
}

// SelectOOBWithStrategy allows you to select content from a response to be swapped in via an out-of-band swap, with an optional strategy for each selector.
//
// The value of this attribute is comma separated list of elements to be swapped out of band. This attribute is almost always paired with hx-select.
//
// Each value in the comma separated list of values can specify any valid hx-swap strategy by separating the selector and the swap strategy with a :.
//
// For example, to prepend the alert content instead of replacing it:
//
//	<div>
//	   <div id="alert"></div>
//	    <button
//				{ hx.New().
//					Get("/info").
//					Select("#info-details").
//					Swap(swap.OuterHTML).
//					SelectOOBWithStrategy(
//						htmx.SelectOOBStrategy{Selector:"#alert", Strategy: swap.AfterBegin},
//					).
//					Build()... }
//			>
//	        Get Info!
//	    </button>
//	</div>
//
// # Notes
//
// hx-select-oob is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-select-oob]
//
// [hx-select-oob]: https://htmx.org/attributes/hx-select-oob
func (hx HX) SelectOOBWithStrategy(selectors ...SelectOOBStrategy) HX {
	values := make([]string, len(selectors))
	for i, s := range selectors {
		if s.Strategy == "" {
			values[i] = s.Selector
		} else {
			values[i] = fmt.Sprintf("%s:%s", s.Selector, s.Strategy)
		}
	}

	hx.attrs["hx-select-oob"] = strings.Join(values, ",")
	return hx
}

// Swap allows you to specify how the response will be swapped in relative to the target of an AJAX request.
//
// So in this code:
//
//	<div {hx.New().Get("/example").Swap(swap.AfterEnd).Build()...} >
//		Get Some HTML & Append It
//	</div>
//
// The div will issue a request to /example and append the returned content after the div.
//
// For advanced usage with modifiers, see [HX.SwapExtended()].
//
// HTMX Attribute: [hx-swap]
//
// [hx-swap]: https://htmx.org/attributes/hx-swap
func (hx HX) Swap(strategy swap.Strategy) HX {
	hx.attrs["hx-swap"] = string(strategy)
	return hx
}

// SwapExtended allows you to specify how the response will be swapped in relative to the target of an AJAX request, with modifiers for changing the behavior of the swap.
//
// So in this code:
//
//	<div {hx.New().Get("/example").SwapExtended(swap.New().Strategy(swap.AfterEnd)).Build()...} >
//		Get Some HTML & Append It
//	</div>
//
// The div will issue a request to /example and append the returned content after the div.
//
// For documentation about modifiers, see [swap.Builder].
//
// HTMX Attribute: [hx-swap]
//
// [hx-swap]: https://htmx.org/attributes/hx-swap
func (hx HX) SwapExtended(swap *swap.Builder) HX {
	hx.attrs["hx-swap"] = swap.String()
	return hx
}

// SwapOOP allows you to specify that some content in a response should be swapped into the DOM somewhere other than the target by ID, that is “Out of Band”. This allows you to piggy back updates to other element updates on a response.
//
// Consider the following response HTML:
//
//	<div>
//	...
//	</div>
//	<div id="alerts" {hx.New().SwapOOB().Build()...}>
//		 Saved!
//	</div>
//
// The first div will be swapped into the target the usual manner. The second div, however, will be swapped in as a replacement for the element with the id alerts, and will not end up in the target.
//
// If the value is true or outerHTML (which are equivalent) the element will be swapped inline.
//
// # Notes
//
// hx-swap-oob is not inherited
//
// HTMX Attribute: [hx-swap-oob]
//
// [hx-swap-oob]: https://htmx.org/attributes/hx-swap-oob
func (hx HX) SwapOOB() HX {
	hx.attrs["hx-swap-oob"] = "true"
	return hx
}

// SwapOOBWithStrategy allows you to specify that some content in a response should be swapped into the DOM somewhere other than the target by ID with a swap strategy, that is “Out of Band”. This allows you to piggy back updates to other element updates on a response.
//
// Consider the following response HTML:
//
//	<div>
//	...
//	</div>
//	<div id="alerts" {hx.New().SwapOOBWithStrategy(swap.AfterBegin).Build()...}>
//		 Saved!
//	</div>
//
// The first div will be swapped into the target the usual manner. The second div, however, will be swapped in after the start of the element with the id #alerts, and will not end up in the target.
//
// If the value is outerHTML (which is equivalent to [HX.SwapOOB]) the element will be swapped inline.
//
// # Notes
//
// hx-swap-oob is not inherited
//
// HTMX Attribute: [hx-swap-oob]
//
// [hx-swap-oob]: https://htmx.org/attributes/hx-swap-oob
func (hx HX) SwapOOBWithStrategy(strategy swap.Strategy) HX {
	hx.attrs["hx-swap-oob"] = string(strategy)
	return hx
}

// SwapOOP allows you to specify that some content in a response should be swapped into the DOM somewhere other than the target by selector, that is “Out of Band”. This allows you to piggy back updates to other element updates on a response.
//
// Consider the following response HTML:
//
//	<div>
//	...
//	</div>
//	<div {hx.New().SwapOOBSelector(swap.OuterHTML, "#alerts").Build()...}>
//		 Saved!
//	</div>
//
// The first div will be swapped into the target the usual manner. The second div, however, will be swapped in as a replacement for the element with the id #alerts, and will not end up in the target.
//
// If the value is outerHTML the element will be swapped inline.
//
// # Notes
//
// hx-swap-oob is not inherited
//
// HTMX Attribute: [hx-swap-oob]
//
// [hx-swap-oob]: https://htmx.org/attributes/hx-swap-oob
func (hx HX) SwapOOBSelector(strategy swap.Strategy, selector string) HX {
	hx.attrs["hx-swap-oob"] = fmt.Sprintf("%s:%s", strategy, selector)
	return hx
}

// Target allows you to target a different element for swapping than the one issuing the AJAX request. The value of this attribute can be:
//
//   - A CSS query selector of the element to target.
//   - this which indicates that the element that the hx-target attribute is on is the target.
//   - closest <CSS selector> which will find the closest ancestor element or itself, that matches the given CSS selector (e.g. closest tr will target the closest table row to the element).
//   - find <CSS selector> which will find the first child descendant element that matches the given CSS selector.
//   - next which resolves to element.nextElementSibling
//   - next <CSS selector> which will scan the DOM forward for the first element that matches the given CSS selector. (e.g. next .error will target the closest following sibling element with error class)
//   - previous which resolves to element.previousElementSibling
//   - previous <CSS selector> which will scan the DOM backwards for the first element that matches the given CSS selector. (e.g previous .error will target the closest previous sibling with error class)
//
// For targeting a special target like `this`, see [HX.TargetSpecial()].
//
// For targeting finding the nearest element, see [HX.TargetRelative()].
//
// Here is an example that targets a div:
//
//	<div>
//		<div id="response-div"></div>
//	 	<button {
//			hx.New().
//			Post("/register").
//			Target("#response-div").
//			Swap(swap.BeforeEnd).
//			Build()...}
//			>
//	 		Register!
//	 	</button>
//	</div>
//
// The response from the /register url will be appended to the div with the id response-div.
//
// # Notes
//
// hx-target is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-target]
//
// [hx-target]: https://htmx.org/attributes/hx-target
func (hx HX) Target(selector string) HX {
	hx.attrs["hx-target"] = selector
	return hx
}

// A TargetSpecialType is a special HTMX target for swapping.
type TargetSpecialType string

const (
	TargetThis     TargetSpecialType = "this"     // indicates that the element that the hx-target attribute is on is the target.
	TargetNext     TargetSpecialType = "next"     // resolves to element.nextElementSibling
	TargetPrevious TargetSpecialType = "previous" // resolves to element.previousElementSibling
)

// TargetSpecial allows you to target a different element for swapping than the one issuing the AJAX request. The value of this attribute can be:
//
//   - this which indicates that the element that the hx-target attribute is on is the target.
//   - next which resolves to element.nextElementSibling
//   - previous which resolves to element.previousElementSibling
//
// For targeting with a general selector target, see [HX.Target()].
//
// For targeting finding the nearest element, see [HX.TargetRelative()].
//
// This example uses hx-target="this" to make a link that updates itself when clicked:
//
// <a hx-post="/new-link" hx-target="this" hx-swap="outerHTML">New link</a>
//
// # Notes
//
// hx-target is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-target]
//
// [hx-target]: https://htmx.org/attributes/hx-target
func (hx HX) TargetSpecial(target TargetSpecialType) HX {
	hx.attrs["hx-target"] = string(target)
	return hx
}

// A TargetSelectorType is a special named HTMX target for swapping.
type TargetSelectorType string

const (
	TargetSelectorClosest  TargetSelectorType = "closest"  // find the closest ancestor element or itself, that matches the given CSS selector
	TargetSelectorFind     TargetSelectorType = "find"     // find the first child descendant element that matches the given CSS selector
	TargetSelectorNext     TargetSelectorType = "next"     // which will scan the DOM forward for the first element that matches the given CSS selector.
	TargetSelectorPrevious TargetSelectorType = "previous" // scan the DOM backwards for the first element that matches the given CSS selector
)

// TargetRelative allows you to target a different element for swapping than the one issuing the AJAX request, and find the target relative to the current element. The value of this attribute can be:
//
//   - closest <CSS selector> which will find the closest ancestor element or itself, that matches the given CSS selector (e.g. closest tr will target the closest table row to the element).
//   - find <CSS selector> which will find the first child descendant element that matches the given CSS selector.
//   - next <CSS selector> which will scan the DOM forward for the first element that matches the given CSS selector. (e.g. next .error will target the closest following sibling element with error class)
//   - previous <CSS selector> which will scan the DOM backwards for the first element that matches the given CSS selector. (e.g previous .error will target the closest previous sibling with error class)
//
// For targeting a special target like `this`, see [HX.TargetElement()].
//
// Here is an example that targets the previous div by ID:
//
//	<div>
//		<div id="response-div">Not me</div>
//		<div id="response-div"></div>
//	 	<button {
//			hx.New().
//			Post("/register").
//			TargetRelative(htmx.TargetSelectorPrevious, "#response-div").
//			Swap(swap.BeforeEnd).
//			Build()...}
//			>
//	 		Register!
//	 	</button>
//	</div>
//
// The response from the /register url will be appended to the first previous div with the id response-div.
//
// # Notes
//
// hx-target is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-target]
//
// [hx-target]: https://htmx.org/attributes/hx-target
func (hx HX) TargetRelative(targetType TargetSelectorType, selector string) HX {
	hx.attrs["hx-target"] = fmt.Sprintf("%s %s", targetType, selector)
	return hx
}

// Trigger allows you to specify what triggers an AJAX request. A trigger value can be one of the following:
//
// TODO: Support trigger similarly to [HX.SwapExtended]
//
// HTMX Attribute: [hx-trigger]
//
// [hx-trigger]: https://htmx.org/attributes/hx-trigger/
func (hx HX) Trigger(event string) HX {
	hx.attrs["hx-trigger"] = event
	return hx
}

// Vals allows you to add to the parameters that will be submitted with an AJAX request.
//
// The value of this attribute is a list of name-expression values in JSON (JavaScript Object Notation) format, marshaled from a struct or map.
//
// By default, the value of hx-vals must be valid JSON. It is not dynamically computed.
//
// # Notes
//
// hx-vals is inherited and can be placed on a parent element.
// A child declaration of a variable overrides a parent declaration.
// Input values with the same name will be overridden by variable declarations.
//
// HTMX Attribute: [hx-vals]
//
// [hx-vals]: https://htmx.org/attributes/hx-vals
func (hx HX) Vals(vals any) HX {
	json, err := json.Marshal(vals)
	if err != nil {
		// Silently ignore the value if there is an error, because there's not a good way to report an error when constructing templ attributes.
		return hx
	}
	hx.attrs["hx-vals"] = string(json)
	return hx
}

// ValsJS allows you to add to the parameters that will be submitted with an AJAX request, using JavaScript to compute the values.
//
// Pass a map[string]string to this method, to generate a Javascript object. The values should be valid JavaScript expressions.
//
// When using evaluated code you can access the event object. This example includes the value of the last typed key within the input.
//
//	<div {hx.New().Get("/example").Trigger("keyup").ValsJS(map[string]string{"lastKey": "event.key"}).Build()...} >
//		<input type="text" />
//	</div>
//
// # Security Considerations
//
// If you use the javascript: prefix, be aware that you are introducing security considerations, especially when dealing with user input such as query strings or user-generated content, which could introduce a Cross-Site Scripting (XSS) vulnerability.
//
// # Notes
//
// hx-vals is inherited and can be placed on a parent element.
// A child declaration of a variable overrides a parent declaration.
// Input values with the same name will be overridden by variable declarations.
//
// HTMX Attribute: [hx-vals]
//
// [hx-vals]: https://htmx.org/attributes/hx-val
func (hx HX) ValsJS(vals map[string]string) HX {
	values := make([]string, len(vals))

	i := 0
	for k, v := range vals {
		// Note that values are not wrapped in "", because they are javascript expressions.
		values[i] = fmt.Sprintf(`%s: %s`, k, v)
		i++
	}
	// Sort by keys.
	slices.Sort(values)

	hx.attrs["hx-vals"] = fmt.Sprintf("js:{%s}", strings.Join(values, ", "))
	return hx
}

// Additional Attributes

func (hx HX) Confirm(msg string) HX {
	hx.attrs["hx-confirm"] = msg
	return hx
}
func (hx HX) Delete(url string) HX {
	hx.attrs["hx-delete"] = url
	return hx
}

func (hx HX) Disable() HX {
	hx.attrs["hx-disable"] = true
	return hx
}

func (hx HX) DisabledElt(selector string) HX {
	hx.attrs["hx-disabled-elt"] = selector
	return hx
}

// TODO: Typed disinherit https://htmx.org/attributes/hx-disinherit/
func (hx HX) Disinherit(attr string) HX {
	hx.attrs["hx-disinherit"] = attr
	return hx
}

type Encoding string

const (
	EncodingMultipart Encoding = "multipart/form-data"
)

func (hx HX) Encoding(encoding Encoding) HX {
	hx.attrs["hx-encoding"] = encoding
	return hx
}
func (hx HX) Ext(ext string) HX {
	hx.attrs["hx-ext"] = ext
	return hx
}

func (hx HX) Headers(headers any) HX {
	json, err := json.Marshal(headers)
	if err != nil {
		// Silently ignore the value if there is an error, because there's not a good way to report an error when constructing templ attributes.
		return hx
	}
	hx.attrs["hx-headers"] = fmt.Sprintf("js:%s", json)
	return hx
}

func (hx HX) HeadersJS(headers any) HX {
	json, err := json.Marshal(headers)
	if err != nil {
		// Silently ignore the value if there is an error, because there's not a good way to report an error when constructing templ attributes.
		return hx
	}
	hx.attrs["hx-headers"] = string(json)
	return hx
}

func (hx HX) History(on bool) HX {
	hx.attrs["hx-history"] = boolToString(on)
	return hx
}

// HistoryElt allows you to specify the element that will be used to snapshot and restore page state during navigation. By default, the body tag is used. This is typically good enough for most setups, but you may want to narrow it down to a child element. Just make sure that the element is always visible in your application, or htmx will not be able to restore history navigation properly.
//
// Here is an example:
//
//	<html>
//	<body>
//	<div id="content" {hx.New().HistoryElt().Build()...}>
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
func (hx HX) HistoryElt() HX {
	hx.attrs["hx-history-elt"] = true
	return hx
}

// include additional data in requests
func (hx HX) Include(selector string) HX {
	hx.attrs["hx-include"] = selector
	return hx
}

// include additional data in requests
func (hx HX) IncludeThis() HX {
	hx.attrs["hx-include"] = "this"
	return hx
}

type IncludeExtension string

const (
	IncludeClosest  IncludeExtension = "closest"
	IncludeFind     IncludeExtension = "find"
	IncludeNext     IncludeExtension = "next"
	IncludePrevious IncludeExtension = "previous"
)

// include additional data in requests
func (hx HX) IncludeExtended(extension IncludeExtension, selector string) HX {
	hx.attrs["hx-include"] = fmt.Sprintf("%s %s", extension, selector)
	return hx
}

func (hx HX) Indicator(selector string) HX {
	hx.attrs["hx-indicator"] = selector
	return hx
}
func (hx HX) IndicatorClosest(selector string) HX {
	hx.attrs["hx-indicator"] = fmt.Sprintf("closest %s", selector)
	return hx
}

func (hx HX) ParamsAll() HX {
	hx.attrs["hx-params"] = "*"
	return hx
}

func (hx HX) ParamsNone() HX {
	hx.attrs["hx-params"] = "none"
	return hx
}

func (hx HX) Params(params []string) HX {
	hx.attrs["hx-params"] = strings.Join(params, ",")
	return hx
}

func (hx HX) ParamsNot(params []string) HX {
	hx.attrs["hx-params"] = fmt.Sprintf("not %s", strings.Join(params, ","))
	return hx
}

func (hx HX) Patch(url string) HX {
	hx.attrs["hx-patch"] = url
	return hx
}

func (hx HX) Preserve() HX {
	hx.attrs["hx-preserve"] = true
	return hx
}

func (hx HX) Prompt(msg string) HX {
	hx.attrs["hx-prompt"] = msg
	return hx
}

func (hx HX) Put(url string) HX {
	hx.attrs["hx-put"] = url
	return hx
}

func (hx HX) ReplaceURL(on bool) HX {
	hx.attrs["hx-replace-url"] = boolToString(on)
	return hx
}

func (hx HX) ReplaceURLWith(url string) HX {
	hx.attrs["hx-replace-url"] = url
	return hx
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

func (hx HX) Request(request Request) HX {
	hx.attrs["hx-request"] = request.String()
	return hx
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

func (hx HX) Sync(selector string) HX {
	hx.attrs["hx-sync"] = selector
	return hx
}

func (hx HX) SyncStrategy(selector string, strategy SyncStrategy) HX {
	hx.attrs["hx-sync"] = fmt.Sprintf("%s:%s", selector, strategy)
	return hx
}

func (hx HX) Validate() HX {
	hx.attrs["hx-validate"] = true
	return hx
}

// More allow you to merge arbitrary maps into the final attributes.
// This allows additional attributes to be passed down in a single map.
func (hx HX) More(more map[string]any) HX {
	for k, v := range more {
		hx.attrs[k] = v
	}
	return hx
}

func boolToString(hx bool) string {
	if hx {
		return "true"
	}
	return "false"
}
