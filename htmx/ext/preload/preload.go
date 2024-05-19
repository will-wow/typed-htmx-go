// package preload allows you to load HTML fragments into your browser’s cache before they are requested by the user, so that additional pages appear to users to load nearly instantaneously. As a developer, you can customize its behavior to fit your applications needs and use cases.
package preload

import (
	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/internal/util"
)

// Extension allows you to load HTML fragments into your browser’s cache before they are requested by the user, so that additional pages appear to users to load nearly instantaneously. As a developer, you can customize its behavior to fit your applications needs and use cases.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@1.9.12/dist/ext/preload.js"></script>
//
// Extension: [preload]
//
// [preload]: https://htmx.org/extensions/preload/
const Extension htmx.Extension = "preload"

// Preload adds a preload attribute to any hyperlinks and hx-get elements you want to preload. By default, resources will be loaded as soon as the mousedown event begins, giving your application a roughly 100-200ms head start on serving responses.
//
// Extension: [preload]
//
// [preload]: https://htmx.org/extensions/preload/
func Preload[T any](hx htmx.HX[T]) T {
	return hx.Attr("preload", true)
}

// A PreloadEvent represents the event that triggers the preload.
type PreloadEvent string

const (
	MouseDown PreloadEvent = "mousedown"    // The default behavior for this extension is to begin loading a resource when the user presses the mouse down.
	MouseOver PreloadEvent = "mouseover"    // To preload links more aggressively, you can trigger the preload to happen when the user’s mouse hovers over the link instead.
	Init      PreloadEvent = "preload:init" // The extension itself generates an event called preload:init that can be used to trigger preloads as soon as an object has been processed by htmx.
)

// PreloadOn adds a preload attribute to any hyperlinks and hx-get elements you want to preload, specifying the event that triggers the preload.
//
// Extension: [preload]
//
// [preload]: https://htmx.org/extensions/preload/#preload-mouseover
func PreloadOn[T any](hx htmx.HX[T], event PreloadEvent) T {
	return hx.Attr("preload", string(event))
}

// PreloadImages preloads linked image resources after an HTML page (or page fragment) is preloaded.
//
// Extension: [preload]
//
// [preload]: https://htmx.org/extensions/preload/#preloading-of-linked-images
func PreloadImages[T any](hx htmx.HX[T], preloadImages bool) T {
	return hx.Attr("preload-images", util.BoolToString(preloadImages))
}
