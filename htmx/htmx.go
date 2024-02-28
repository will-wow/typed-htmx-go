// package base provides Go functions for building HTMX attributes.
// These should usually be used through a proxy like templ/hx or gomponents/hx.
//
// See [typed-htmx-go.vercel.app] for example usage.
//
// [typed-htmx-go.vercel.app]: https://typed-htmx-go.vercel.app/
package htmx

import (
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx/swap"
	"github.com/will-wow/typed-htmx-go/htmx/trigger"
)

type NewAttr[T any] func(key Attribute, value any) T

// An HX constructs HTMX attributes.
type HX[T any] struct {
	attr NewAttr[T]
}

func NewHX[T any](attr NewAttr[T]) HX[T] {
	return HX[T]{
		attr: attr,
	}
}

// A StandardCSSSelector is any valid CSS selector, like #element or `.class > button`.
type StandardCSSSelector string

// Boost allows you to “boost” normal anchors and form tags to use AJAX instead. This has the [nice fallback] that, if the user does not have javascript enabled, the site will continue to work.
//
// For anchor tags, clicking on the anchor will issue a GET request to the url specified in the href and will push the url so that a history entry is created. The target is the <body> tag, and the innerHTML swap strategy is used by default. All of these can be modified by using the appropriate attributes, except the click trigger.
//
// For forms the request will be converted into a GET or POST, based on the method in the method attribute and will be triggered by a submit. Again, the target will be the body of the page, and the innerHTML swap will be used. The url will not be pushed, however, and no history entry will be created. (You can use the [HX.PushUrl] attribute if you want the url to be pushed.)
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
func (hx *HX[T]) Boost(boost bool) T {
	if boost {
		return hx.attr("hx-boost", "true")
	}
	return hx.attr("hx-boost", "false")
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
func (hx *HX[T]) Get(url string) T {
	return hx.attr(Get, url)
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
func (hx *HX[T]) Post(url string) T {
	return hx.attr(Post, url)
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
func (hx *HX[T]) On(event string, action string) T {
	return hx.attr(Attribute(fmt.Sprintf("hx-on:%s", event)), action)
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
func (hx *HX[T]) OnHTMX(event string, action string) T {
	return hx.attr(Attribute(fmt.Sprintf("hx-on::%s", event)), action)
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
func (hx *HX[T]) PushURL(on bool) T {
	return hx.attr(PushURL, boolToString(on))
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
func (hx *HX[T]) PushURLPath(url string) T {
	return hx.attr(PushURL, url)
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
func (hx *HX[T]) Select(selector StandardCSSSelector) T {
	return hx.attr(Select, string(selector))
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
func (hx *HX[T]) SelectOOB(selectors ...StandardCSSSelector) T {
	return hx.attr(SelectOOB, joinStringLikes(selectors, ","))
}

type SelectOOBStrategy struct {
	Selector StandardCSSSelector
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
func (hx *HX[T]) SelectOOBWithStrategy(selectors ...SelectOOBStrategy) T {
	values := make([]string, len(selectors))
	for i, s := range selectors {
		if s.Strategy == "" {
			values[i] = string(s.Selector)
		} else {
			values[i] = fmt.Sprintf("%s:%s", s.Selector, s.Strategy)
		}
	}

	return hx.attr(SelectOOB, strings.Join(values, ","))
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
func (hx *HX[T]) Swap(strategy swap.Strategy) T {
	return hx.attr(Swap, string(strategy))
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
func (hx *HX[T]) SwapExtended(swap *swap.Builder) T {
	return hx.attr(Swap, swap.String())
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
func (hx *HX[T]) SwapOOB() T {
	return hx.attr(SwapOOB, "true")
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
func (hx *HX[T]) SwapOOBWithStrategy(strategy swap.Strategy) T {
	return hx.attr(SwapOOB, string(strategy))
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
func (hx *HX[T]) SwapOOBSelector(strategy swap.Strategy, extendedSelector string) T {
	return hx.attr(SwapOOB, fmt.Sprintf("%s:%s", strategy, extendedSelector))
}

type TargetSelector string

const (
	TargetThis     TargetSelector = "this"     // indicates that the element that the hx-target attribute is on is the target.
	TargetNext     TargetSelector = "next"     // resolves to element.nextElementSibling
	TargetPrevious TargetSelector = "previous" // resolves to element.previousElementSibling
)

var TargetRelative = makeRelativeSelector[SelectorModifier, TargetSelector]()

// Target allows you to target a different element for swapping than the one issuing the AJAX request.
//
// You can pass an extended selector to this method, using [RelativeSelector].
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
func (hx *HX[T]) Target(extendedSelector TargetSelector) T {
	return hx.attr(Target, string(extendedSelector))
}

// Trigger allows you to specify what event triggers an AJAX request.
//
// For usage with modifiers and polling, see [HX.TriggerExtended()].
//
// HTMX Attribute: [hx-trigger]
//
// [hx-trigger]: https://htmx.org/attributes/hx-trigger/
func (hx *HX[T]) Trigger(event string) T {
	return hx.attr(Trigger, event)
}

// TriggerExtended allows you to specify what triggers an AJAX request, with modifiers for changing the behavior of the trigger.
// A trigger value can be one of the following:
//
//   - An event name (e.g. “click” or “my-custom-event”) followed by an event filter and a set of event modifiers
//   - A polling definition of the form every <timing declaration>
//   - A comma-separated list of such events
//
// See [trigger.Event] and [trigger.Poll] for more information on options.
//
// HTMX Attribute: [hx-trigger]
//
// [hx-trigger]: https://htmx.org/attributes/hx-trigger/
func (hx *HX[T]) TriggerExtended(triggers ...trigger.Trigger) T {
	values := make([]string, len(triggers))
	for i, t := range triggers {
		values[i] = t.String()
	}

	return hx.attr(Trigger, strings.Join(values, ", "))
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
func (hx *HX[T]) Vals(vals any) T {
	json, err := json.Marshal(vals)
	if err != nil {
		// Silently ignore the value if there is an error, because there's not a good way to report an error when constructing templ attributes.
		var empty T
		return empty
	}
	return hx.attr(Vals, string(json))
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
func (hx *HX[T]) ValsJS(vals map[string]string) T {
	return hx.attr(Vals, mapToJS(vals))
}

// Additional Attributes

// Confirm allows you to confirm an action before issuing a request. This can be useful in cases where the action is destructive and you want to ensure that the user really wants to do it.
//
// Here is an example:
//
//	<button {hx.New().Delete("/account").Confirm("Are you sure you wish to delete you account?").Build()...}>
//	  Delete My Account
//	</button>
//
// # Event details
//
// The event triggered by hx-confirm contains additional properties in its detail:
//
//   - triggeringEvent: the event that triggered the original request
//   - issueRequest(skipConfirmation=false): a callback which can be used to confirm the AJAX request
//   - question: the value of the hx-confirm attribute on the HTML element
//
// # Notes
//
//   - hx-confirm is inherited and can be placed on a parent element
//   - hx-confirm uses the browser’s window.confirm by default. You can customize this behavior as shown in this example.
//   - a boolean skipConfirmation can be passed to the issueRequest callback; if true (defaults to false), the window.confirm will not be called and the AJAX request is issued directly
//
// HTMX Attribute: [hx-confirm]
//
// [hx-confirm]: https://htmx.org/attributes/hx-confirm/
func (hx *HX[T]) Confirm(msg string) T {
	return hx.attr(Confirm, msg)
}

// Delete will cause an element to issue a DELETE to the specified URL and swap the HTML into the DOM using a swap strategy:
//
//	<button {hx.New().Delete("/account").Target("body").Build()...} >
//		Delete Your Account
//	</button>
//
// This example will cause the button to issue a DELETE to /account and swap the returned HTML into the innerHTML of the body.
//
// # Notes
//
//   - hx-delete is not inherited
//   - You can control the target of the swap using the [HX.Target] attribute
//   - You can control the swap strategy by using the [HX.Swap] attribute
//   - You can control what event triggers the request with the [HX.Trigger] attribute
//   - You can control the data submitted with the request in various ways, documented here: [Parameters]
//   - To remove the element following a successful DELETE, return a 200 status code with an empty body; if the server responds with a 204, no swap takes place, documented here: [Requests & Responses]
//
// HTMX Attribute: [hx-delete]
//
// [hx-delete]: https://htmx.org/attributes/hx-delete
// [Parameters]: https://htmx.org/docs/#parameters
// [Requests & Responses]: https://htmx.org/docs/#requests
func (hx *HX[T]) Delete(url string) T {
	return hx.attr(Delete, url)
}

// Disable will disable htmx processing for a given element and all its children. This can be useful as a backup for HTML escaping, when you include user generated content in your site, and you want to prevent malicious scripting attacks.
//
// The value of the tag is ignored, and it cannot be reversed by any content beneath it.
//
// HTMX Attribute: [hx-disable]
//
// [hx-disable]: https://htmx.org/attributes/hx-disable
func (hx *HX[T]) Disable() T {
	return hx.attr(Disable, true)
}

type DisabledEltModifier string

const DisabledEltClosest DisabledEltModifier = "closest"

type DisabledEltSelector string

const DisabledEltThis DisabledEltSelector = "this"

var DisabledEltRelative = makeRelativeSelector[DisabledEltModifier, DisabledEltSelector]()

// DisabledElt allows you to specify elements that will have the disabled attribute added to them for the duration of the request.
//
// The value of this attribute is a CSS query selector of the element or elements to apply the class to, or the keyword closest, followed by a CSS selector, which will find the closest ancestor element or itself, that matches the given CSS selector (e.g. closest tr), or the keyword this
//
// Here is an example with a button that will disable itself during a request:
//
//	<button { hx.New().Post("/example").DisabledElt(hx.This).Build()...} >
//		Post It!
//	</button>
//
// When a request is in flight, this will cause the button to be marked with the disabled attribute, which will prevent further clicks from occurring.
//
// HTMX Attribute: [hx-disabled-elt]
//
// [hx-disabled-elt]: https://htmx.org/attributes/hx-disabled-elt
func (hx *HX[T]) DisabledElt(selector DisabledEltSelector) T {
	return hx.attr(DisabledElt, string(selector))
}

// Disinherit allows you to disable automatic attribute inheritance for one or multiple specified attributes.
//
// The default behavior for htmx is to “inherit” many attributes automatically: that is, an attribute such as hx-target may be placed on a parent element, and all child elements will inherit that target.
//
// An example scenario is to allow you to place an hx-boost on the body element of a page, but overriding that behavior in a specific part of the page to allow for more specific behaviors.
//
//	<div {hx.New().Boost(true).Select("#content").Target("#content").Disinherit(hx.Target).Build()...} >
//		<!-- hx-select is automatically set to parent value; hx-target is not inherited -->
//	  <button {hx.New().Get("/test").Build()...}></button>
//	</div>
//
//	<div {hx.New().Select("#content").Build()...} >
//		<div {hx.New().Boost(true).Target("#content").Disinherit(hx.Select).Build()...}>
//	  	<!-- hx-target is automatically inherited from parent value -->
//	    <!-- hx-select is not inherited, because the direct parent does
//	    disables inheritance, despite not specifying hx-select itself -->
//	    <button {hx.New().Get("/test").Build()...}></button>
//	  </div>
//	</div>
//
// Notes
//
//   - Read more about [Attribute Inheritance]
//
// HTMX Attribute: [hx-disinherit]
//
// [hx-disinherit]: https://htmx.org/attributes/hx-disinherit/
// [Attribute Inheritance]: https://htmx.org/docs/#inheritance
func (hx *HX[T]) Disinherit(attr ...Attribute) T {
	// Convert to strings for joining.
	attrStrings := make([]string, len(attr))
	for i, a := range attr {
		attrStrings[i] = string(a)
	}

	return hx.attr(Disinherit, strings.Join(attrStrings, " "))
}

// DisinheritAll allows you to disable automatic attribute inheritance for all attributes.
//
// The default behavior for htmx is to “inherit” many attributes automatically: that is, an attribute such as hx-target may be placed on a parent element, and all child elements will inherit that target.
//
// An example scenario is to allow you to place an hx-boost on the body element of a page, but overriding that behavior in a specific part of the page to allow for more specific behaviors.
//
//	<div { hx.New().Boost(true).Select("#content").Target("#content").DisinheritAll().Build()...} hx-boost="true" >
//		<a href="/page1">Go To Page 1</a> <!-- boosted with the attribute settings above -->
//	  <a href="/page2" {hx.New().Unset(hx.Boost).Build()...} >Go To Page 1</a> <!-- not boosted -->
//	  <button {hx.New().Get("/test").TargetNonStandard(hx.TargetThis).Build()... }></button> <!-- hx-select is not inherited -->
//	</div>
//
// Notes
//
//   - Read more about [Attribute Inheritance]
//
// HTMX Attribute: [hx-disinherit]
//
// [hx-disinherit]: https://htmx.org/attributes/hx-disinherit/
// [Attribute Inheritance]: https://htmx.org/docs/#inheritance
func (hx *HX[T]) DisinheritAll() T {
	return hx.attr(Disinherit, "*")
}

// An EncodingContentType is a valid encoding override for an [HX.Encoding()].
type EncodingContentType string

// support file uploads in an ajax request
const EncodingMultipart EncodingContentType = "multipart/form-data"

// Encoding allows you to switch the request encoding from the usual application/x-www-form-urlencoded encoding to multipart/form-data, usually to support file uploads in an ajax request.
//
// The value of this attribute should be "multipart/form-data".
//
// The hx-encoding tag may be placed on parent elements.
//
// # Notes
//
//   - hx-encoding is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-encoding]
//
// [hx-encoding]: https://htmx.org/attributes/hx-encoding
func (hx *HX[T]) Encoding(encoding EncodingContentType) T {
	return hx.attr(Encoding, string(encoding))
}

// Ext enables an htmx [extension] for an element and all its children.
//
// The value can be one or more extension names to apply.
//
// The hx-ext tag may be placed on parent elements if you want a plugin to apply to an entire swath of the DOM, and on the body tag for it to apply to all htmx requests.
//
// # Notes
//
//   - hx-ext is both inherited and merged with parent elements, so you can specify extensions on any element in the DOM hierarchy and it will apply to all child elements.
//
// HTMX Attribute: [hx-ext]
//
// [hx-ext]: https://htmx.org/attributes/hx-ext
// [extension]: https://htmx.org/extensions
func (hx *HX[T]) Ext(ext ...string) T {
	return hx.attr(Ext, strings.Join(ext, ","))
}

// ExtIgnore ignores an [extension] that is defined by a parent node.
//
//	<div {hx.New().Ext("example").Build()...}>
//	  "Example" extension is used in this part of the tree...
//	  <div {hx.New().ExtIgnore("example").Build()...}>
//	    ... but it will not be used in this part.
//	  </div>
//	</div>
//
// HTMX Attribute: [hx-ext]
//
// [hx-ext]: https://htmx.org/attributes/hx-ext
// [extension]: https://htmx.org/extensions
func (hx *HX[T]) ExtIgnore(ext string) T {
	return hx.attr(Ext, fmt.Sprintf("ignore:%s", ext))
}

// Headers allows you to add to the headers that will be submitted with an AJAX request.
//
// The value of this attribute is a list of name-expression values in JSON (JavaScript Object Notation) format.
//
// For values computed at runtime, see [HX.HeadersJS()].
//
// # Notes
//
//   - hx-headers is inherited and can be placed on a parent element.
//   - A child declaration of a header overrides a parent declaration.
//
// HTMX Attribute: [hx-headers]
//
// [hx-headers]: https://htmx.org/attributes/hx-headers
func (hx *HX[T]) Headers(headers any) T {
	json, err := json.Marshal(headers)
	if err != nil {
		// Silently ignore the value if there is an error, because there's not a good way to report an error when constructing templ attributes.
		var empty T
		return empty
	}
	return hx.attr(Headers, string(json))
}

// HeadersJS allows you to add to the headers that will be submitted with an AJAX request, with values evaluated as JavaScript expressions at runtime.
//
// # Security Considerations
//
// Be aware that you are introducing security considerations, especially when dealing with user input such as query strings or user-generated content, which could introduce a Cross-Site Scripting (XSS) vulnerability.
//
// For values static JSON, see [HX.Headers()].
//
// # Notes
//
//   - hx-headers is inherited and can be placed on a parent element.
//   - A child declaration of a header overrides a parent declaration.
//
// HTMX Attribute: [hx-headers]
//
// [hx-headers]: https://htmx.org/attributes/hx-headers
func (hx *HX[T]) HeadersJS(headers map[string]string) T {
	return hx.attr(Headers, mapToJS(headers))
}

// History when set to false on any element in the current document, or any html fragment loaded into the current document by htmx, will prevent sensitive data being saved to the localStorage cache when htmx takes a snapshot of the page state.
//
// History navigation will work as expected, but on restoration the URL will be requested from the server instead of the history cache.
//
// Here is an example:
//
//	<html>
//	<body>
//	<div {hx.New().History(false).Build()...}>
//	 ...
//	</div>
//	</body>
//	</html>
//
// # Notes
//   - hx-history="false" can be present anywhere in the document to embargo the current page state from the history cache (i.e. even outside the element specified for the history snapshot hx-history-elt).
//
// HTMX Attribute: [hx-history]
//
// [hx-history]: https://htmx.org/attributes/hx-history/
func (hx *HX[T]) History(on bool) T {
	return hx.attr(History, boolToString(on))
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
func (hx *HX[T]) HistoryElt() T {
	return hx.attr(HistoryElt, true)
}

type IncludeSelector string

const IncludeThis IncludeSelector = "this"

var IncludeRelative = makeRelativeSelector[SelectorModifier, IncludeSelector]()

// Include allows you to include additional element values in an AJAX request.
//
// HTMX Attribute: [hx-include]
//
// [hx-include]: https://htmx.org/attributes/hx-include/
func (hx *HX[T]) Include(selector IncludeSelector) T {
	return hx.attr(Include, string(selector))
}

type IndicatorModifier string

const IndicatorClosest IndicatorModifier = "closest"

type IndicatorSelector string

var IndicatorRelative = makeRelativeSelector[IndicatorModifier, IndicatorSelector]()

// The hx-indicator attribute allows you to specify the element that will have the htmx-request class added to it for the duration of the request. This can be used to show spinners or progress indicators while the request is in flight.
//
// The value of this attribute is a CSS query selector of the element or elements to apply the class to, or the keyword `closest` followed by a CSS selector, which will find the closest ancestor element or itself, that matches the given CSS selector (e.g. closest tr);
//
// HTMX Attribute: [hx-indicator]
//
// [hx-indicator]: https://htmx.org/attributes/hx-indicator/
func (hx *HX[T]) Indicator(extendedSelector IndicatorSelector) T {
	return hx.attr(Indicator, string(extendedSelector))
}

// ParamsAll allows you to include all parameters with an AJAX request (default).
//
//	<div {hx.New().Get("/example").ParamsAll().Build()...}>Get Some HTML, Including Params</div>
//
// This div will include all the parameters that a POST would, but they will be URL encoded and included in the URL, as per usual with a GET.
//
// # Notes
//
//   - hx-params is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-params]
//
// [hx-params]: https://htmx.org/attributes/hx-params/
func (hx *HX[T]) ParamsAll() T {
	return hx.attr(Params, "*")
}

// ParamsNone allows you to include no parameters with an AJAX request.
//
// # Notes
//
//   - hx-params is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-params]
//
// [hx-params]: https://htmx.org/attributes/hx-params/
func (hx *HX[T]) ParamsNone() T {
	return hx.attr(Params, "none")
}

// Params allows you to filter the parameters that will be submitted with an AJAX request.
//
// # Notes
//
//   - hx-params is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-params]
//
// [hx-params]: https://htmx.org/attributes/hx-params/
func (hx *HX[T]) Params(paramNames ...string) T {
	return hx.attr(Params, strings.Join(paramNames, ","))
}

// ParamsNot allows you to include all params except the comma separated list of parameter
// when submitting an AJAX request.
//
// # Notes
//
//   - hx-params is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-params]
//
// [hx-params]: https://htmx.org/attributes/hx-params/
func (hx *HX[T]) ParamsNot(paramNames ...string) T {
	return hx.attr(Params, fmt.Sprintf("not %s", strings.Join(paramNames, ",")))
}

func (hx *HX[T]) Patch(url string) T {
	return hx.attr(Patch, url)
}

// Preserve allows you to keep an element unchanged during HTML replacement. Elements with hx-preserve set are preserved by id when htmx updates any ancestor element. You must set an unchanging id on elements for hx-preserve to work. The response requires an element with the same id, but its type and other attributes are ignored.
//
// Note that some elements cannot unfortunately be preserved properly, such as <input type="text"> (focus and caret position are lost), iframes or certain types of videos. To tackle some of these cases we recommend the [morphdom extension], which does a more elaborate DOM reconciliation.
//
// # Notes
//
// hx-preserve is not inherited
//
// HTMX Attribute: [hx-preserve]
//
// [hx-preserve]: https://htmx.org/attributes/hx-preserve/
// [morphdom extension]: https://htmx.org/extensions/morphdom
func (hx *HX[T]) Preserve() T {
	return hx.attr(Preserve, true)
}

func (hx *HX[T]) Prompt(msg string) T {
	return hx.attr(Prompt, msg)
}

func (hx *HX[T]) Put(url string) T {
	return hx.attr(Put, url)
}

// ReplaceURL allows you to replace the current url of the browser location history.
//
// The possible values of this attribute are:
//
//   - true, which replaces the fetched URL in the browser navigation bar.
//   - false, which disables replacing the fetched URL if it would otherwise be replaced due to inheritance.
//
// Here is an example:
//
//	<div {hx.New().Get("/account").ReplaceURL(true).Build()...} >
//		Go to My Account
//	</div>
//
// This will cause htmx to snapshot the current DOM to localStorage and replace the URL `/account’ in the browser location bar.
//
// # Notes
//   - hx-replace-url is inherited and can be placed on a parent element
//   - The HX-Replace-Url response header has similar behavior and can override this attribute.
//   - The hx-history-elt attribute allows changing which element is saved in the history cache.
//   - The hx-push-url attribute is a similar and more commonly used attribute, which creates a new history entry rather than replacing the current one.
//
// HTMX Attribute: [hx-replace]
//
// [hx-replace]: https://htmx.org/attributes/hx-replace/
func (hx *HX[T]) ReplaceURL(on bool) T {
	return hx.attr(ReplaceURL, boolToString(on))
}

// ReplaceURLWith allows you to replace the current url of the browser location history with
// a URL to be replaced into the location bar. This may be relative or absolute, as per [history.replaceState()].
//
//	<div {hx.New().Get("/account").ReplaceURLWith("/account/home").Build()...} >
//		Go to My Account
//	</div>
//
// This will replace the URL `/account/home’ in the browser location bar.
//
// # Notes
//   - hx-replace-url is inherited and can be placed on a parent element
//   - The HX-Replace-Url response header has similar behavior and can override this attribute.
//   - The hx-history-elt attribute allows changing which element is saved in the history cache.
//   - The hx-push-url attribute is a similar and more commonly used attribute, which creates a new history entry rather than replacing the current one.
//
// HTMX Attribute: [hx-replace]
//
// [hx-replace]: https://htmx.org/attributes/hx-replace/
// [history.replaceState()]: https://developer.mozilla.org/en-US/docs/Web/API/History/replaceState
func (hx *HX[T]) ReplaceURLWith(url string) T {
	return hx.attr(ReplaceURL, url)
}

// RequestConfig describes static hx-request attributes
// See https://htmx.org/attributes/hx-request/
type RequestConfig struct {
	Timeout     time.Duration // the timeout for the request
	Credentials bool          // if the request will send credentials
	NoHeaders   bool          // strips all headers from the request
}

func (r RequestConfig) String() string {
	opts := []string{}

	if r.Timeout > 0 {
		opts = append(opts, fmt.Sprintf(`"timeout":%d`, r.Timeout.Milliseconds()))
	}
	if r.Credentials {
		opts = append(opts, `"credentials": true`)
	}
	if r.NoHeaders {
		opts = append(opts, `"noHeaders": true`)
	}

	return strings.Join(opts, ",")
}

// Request allows you to configure various aspects of the request.
// These attributes are set using a JSON-like syntax.
//
// HTMX Attribute: [hx-request]
//
// [hx-request]: https://htmx.org/attributes/hx-request/
func (hx *HX[T]) Request(request RequestConfig) T {
	return hx.attr(Request, request.String())
}

// RequestConfigJS describes runtime hx-request attributes
// See https://htmx.org/attributes/hx-request/
type RequestConfigJS struct {
	Timeout     string // the timeout for the request in milliseconds
	Credentials string // if the request will send credentials
	NoHeaders   string // strips all headers from the request
}

func (r RequestConfigJS) String() string {
	opts := []string{}

	if r.Timeout != "" {
		opts = append(opts, fmt.Sprintf(`"timeout":%s`, r.Timeout))
	}
	if r.Credentials != "" {
		opts = append(opts, fmt.Sprintf(`"credentials": %s`, r.Credentials))
	}
	if r.NoHeaders != "" {
		opts = append(opts, fmt.Sprintf(`"noHeaders": %s`, r.NoHeaders))
	}

	value := strings.Join(opts, ",")
	return fmt.Sprintf("js: %s", value)
}

// RequestJS allows you to configure various aspects of the request, with each value being a valid JavaScript expression.
// To pass a literal string, use wrap it in quotes like "'string'".
//
// HTMX Attribute: [hx-request]
//
// [hx-request]: https://htmx.org/attributes/hx-request/
func (hx *HX[T]) RequestJS(request RequestConfigJS) T {
	return hx.attr(Request, request.String())
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

type SyncSelector string

const SyncThis SyncSelector = "this"

var SyncRelative = makeRelativeSelector[SelectorModifier, SyncSelector]()

// SyncStrategy allows you to synchronize AJAX requests between multiple elements.
//
// The hx-sync attribute consists of a CSS selector to indicate the element to synchronize on. By default, this will use the [SyncDrop] strategy.
//
// You can pass [hx.SyncThis] as a selector to synchronize requests from the current element.
//
// # Notes
//   - hx-sync is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-sync]
//
// [hx-sync]: https://htmx.org/attributes/hx-sync/
func (hx *HX[T]) Sync(extendedSelector SyncSelector) T {
	return hx.attr(Sync, string(extendedSelector))
}

// SyncStrategy allows you to synchronize AJAX requests between multiple elements.
//
// The hx-sync attribute consists of a CSS selector to indicate the element to synchronize on, followed optionally by a colon and then by an optional syncing strategy.
//
// You can pass "hx.This" as a selector to synchronize requests from the current element.
//
// # Notes
//   - hx-sync is inherited and can be placed on a parent element
//
// HTMX Attribute: [hx-sync]
//
// [hx-sync]: https://htmx.org/attributes/hx-sync/
func (hx *HX[T]) SyncStrategy(extendedSelector SyncSelector, strategy SyncStrategy) T {
	return hx.attr(Sync, fmt.Sprintf("%s:%s", extendedSelector, strategy))
}

// Validate will cause an element to validate itself by way of the HTML5 Validation API before it submits a request.
//
// Only <form> elements validate data by default, but other elements do not. Adding hx-validate="true" to <input>, <textarea> or <select> enables validation before sending requests.
//
// # Notes
//   - hx-validate is not inherited
//
// HTMX Attribute: [hx-validate]
//
// [hx-validate]: https://htmx.org/attributes/hx-validate/
func (hx *HX[T]) Validate(validate bool) T {
	return hx.attr(Validate, boolToString(validate))
}

// Non-standard attributes

// Unset sets the value of the selected attributes as "unset"  to clear a property that would normally be inherited (e.g. hx-confirm).
func (hx *HX[T]) Unset(attr Attribute) T {
	return hx.attr(attr, "unset")
}

// // set sets a valid attribute to a value.
// func (hx *HX[T]) set(key Attribute, value any) T {
// 	return T{string(key): value}
// }

// // set sets a non-standard attribute to a value.
// func (hx *HX[T]) setOther(key string, value any) T {
// 	return T{key: value}
// }

// An Attribute is a valid HTMX attribute name. Used for general type changes like `unset` and `disinherit`.
type Attribute string

const (
	Get         Attribute = "hx-get"
	Post        Attribute = "hx-post"
	PushURL     Attribute = "hx-push-url"
	Select      Attribute = "hx-select"
	SelectOOB   Attribute = "hx-select-oob"
	Swap        Attribute = "hx-swap"
	SwapOOB     Attribute = "hx-swap-oob"
	Target      Attribute = "hx-target"
	Trigger     Attribute = "hx-trigger"
	Vals        Attribute = "hx-vals"
	Boost       Attribute = "hx-boost"
	Confirm     Attribute = "hx-confirm"
	Delete      Attribute = "hx-delete"
	Disable     Attribute = "hx-disable"
	DisabledElt Attribute = "hx-disabled-elt"
	Disinherit  Attribute = "hx-disinherit"
	Encoding    Attribute = "hx-encoding"
	Ext         Attribute = "hx-ext"
	Headers     Attribute = "hx-headers"
	History     Attribute = "hx-history"
	HistoryElt  Attribute = "hx-history-elt"
	Include     Attribute = "hx-include"
	Indicator   Attribute = "hx-indicator"
	Params      Attribute = "hx-params"
	Patch       Attribute = "hx-patch"
	Preserve    Attribute = "hx-preserve"
	Prompt      Attribute = "hx-prompt"
	Put         Attribute = "hx-put"
	ReplaceURL  Attribute = "hx-replace-url"
	Request     Attribute = "hx-request"
	Sync        Attribute = "hx-sync"
	Validate    Attribute = "hx-validate"
)

// A SelectorModifier is a relative modifier to a CSS selector. This is used for "extended selectors".
// Some attributes only support a subset of these, but any Relative function that takes this type supports the full set..
type SelectorModifier string

const (
	Closest  SelectorModifier = "closest"  // find the closest ancestor element or itself, that matches the given CSS selector
	Find     SelectorModifier = "find"     // find the first child descendant element that matches the given CSS selector
	Next     SelectorModifier = "next"     // scan the DOM forward for the first element that matches the given CSS selector. (e.g. next .error will target the closest following sibling element with error class)
	Previous SelectorModifier = "previous" // scan the DOM backwards fo
)

func boolToString(hx bool) string {
	if hx {
		return "true"
	}
	return "false"
}

func mapToJS(vals map[string]string) string {
	values := make([]string, len(vals))

	i := 0
	for k, v := range vals {
		// Note that values are not wrapped in "", because they are javascript expressions.
		values[i] = fmt.Sprintf(`%s:%s`, quoteJSIdentifier(k), v)
		i++
	}
	// Sort by keys.
	slices.Sort(values)

	return fmt.Sprintf("js:{%s}", strings.Join(values, ","))
}

// reJSIdentifier is a regular expression to match valid JavaScript identifiers.
var reJSIdentifier = regexp.MustCompile(`^[a-zA-Z_$][a-zA-Z0-9_$]*$`)

// quoteJSIdentifier quotes a string if it is not a valid JavaScript identifier, for use as a key.
func quoteJSIdentifier(identifier string) string {
	if reJSIdentifier.MatchString(identifier) {
		return identifier
	}
	return fmt.Sprintf(`"%s"`, identifier)
}

// joinStringLikes joins a slice of string-like values into a single string.
func joinStringLikes[T ~string](elems []T, sep string) string {
	var stringElems = make([]string, len(elems))
	for i, x := range elems {
		stringElems[i] = string(x)
	}
	return strings.Join(stringElems, sep)
}

func makeRelativeSelector[Modifier ~string, Selector ~string]() func(Modifier, string) Selector {
	return func(modifier Modifier, selector string) Selector {
		return Selector(fmt.Sprintf("%s %s", modifier, selector))
	}
}
