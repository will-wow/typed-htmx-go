// package loadingstates allows you to easily manage loading states while a request is in flight, including disabling elements, and adding and removing CSS classes.
package loadingstates

import (
	"strconv"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
)

// Extension allows you to easily manage loading states while a request is in flight, including disabling elements, and adding and removing CSS classes.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@1.9.12/dist/ext/loading-states.js"></script>
//
// # Usage
// Add the hx-ext="loading-states" attribute to the body tag or to any parent element containing your htmx attributes.
// Add the following class to your stylesheet to make sure elements are hidden by default:
//
//	[data-loading] {
//	  display: none;
//	}
//
// Extension: [loading-states]
//
// [loading-states]: https://htmx.org/extensions/loading-states/
const Extension htmx.Extension = "loading-states"

// DataLoading shows the element with the default style of inline-block.
//
//	<div data-loading>loading</div>
func DataLoading[T any](hx htmx.HX[T]) T {
	return hx.Attr("data-loading", true)
}

// DataLoadingStyle shows the element with the specified style when a request is in flight.
//
//	<div data-loading="block">loading</div>
//	<div data-loading="flex">loading</div>
func DataLoadingStyle[T any](hx htmx.HX[T], style string) T {
	return hx.Attr("data-loading", style)
}

// DataLoadingClass adds, then removes, CSS classes to the element:
func DataLoadingClass[T any](hx htmx.HX[T], class string) T {
	return hx.Attr("data-loading-class", class)
}

// Removes, then adds back, CSS classes from the element.
func DataLoadingClassRemove[T any](hx htmx.HX[T], class string) T {
	return hx.Attr("data-loading-class-remove", class)
}

// DataLoadingDisable disables an element for the duration of the request.
func DataLoadingDisable[T any](hx htmx.HX[T]) T {
	return hx.Attr("data-loading-disable", true)
}

// DataLoadingAriaBusy adds aria-busy="true" attribute to the element for the duration of the request.
func DataLoadingAriaBusy[T any](hx htmx.HX[T]) T {
	return hx.Attr("data-loading-aria-busy", true)
}

// Some actions may update quickly and showing a loading state in these cases may be more of a distraction. This attribute ensures that the loading state changes are applied only after 200ms if the request is not finished.
//
// You can place the data-loading-delay attribute directly on the element you want to disable, or in any parent element.
func DataLoadingDelay[T any](hx htmx.HX[T]) T {
	return hx.Attr("data-loading-delay", true)
}

// Some actions may update quickly and showing a loading state in these cases may be more of a distraction. This attribute ensures that the loading state changes are applied only after the specified delay if the request is not finished.
//
// You can place the data-loading-delay attribute directly on the element you want to disable, or in any parent element.
func DataLoadingDelayBy[T any](hx htmx.HX[T], delay time.Duration) T {
	return hx.Attr("data-loading-delay", strconv.FormatInt(delay.Milliseconds(), 10))
}

// Allows setting a different target to apply the loading states. The attribute value can be any valid CSS selector. The example below disables the submit button and shows the loading state when the form is submitted.
//
//	<form hx-post="/save"
//		data-loading-target="#loading"
//	  data-loading-class-remove="hidden">
//
//	  <button type="submit" data-loading-disable>Submit</button>
//
//	</form>
//
//	<div id="loading" class="hidden">Loading ...</div>
func DataLoadingTarget[T any](hx htmx.HX[T], targetSelector string) T {
	return hx.Attr("data-loading-target", targetSelector)
}

// DataLoadingPath allows filtering the processing of loading states only for specific requests based on the request path.
//
// You can place the [DataLoadingPath] attribute directly on the loading state element, or in any parent element.
func DataLoadingPath[T any](hx htmx.HX[T], path string) T {
	return hx.Attr("data-loading-path", path)
}

// This attribute is optional and it allows defining a scope for the loading states so only elements within that scope are processed.
func DataLoadingStates[T any](hx htmx.HX[T]) T {
	return hx.Attr("data-loading-states", true)
}
