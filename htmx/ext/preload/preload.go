package preload

import "github.com/will-wow/typed-htmx-go/htmx"

// Extension allows you to load HTML fragments into your browser’s cache before they are requested by the user, so that additional pages appear to users to load nearly instantaneously. As a developer, you can customize its behavior to fit your applications needs and use cases.
const Extension htmx.Extension = "preload"

func Preload[T any](hx htmx.HX[T]) T {
	return hx.Attr("preload", true)
}

type PreloadEvent string

const (
	MouseDown PreloadEvent = "mousedown"    // The default behavior for this extension is to begin loading a resource when the user presses the mouse down.
	MouseOver PreloadEvent = "mouseover"    // To preload links more aggressively, you can trigger the preload to happen when the user’s mouse hovers over the link instead.
	Init      PreloadEvent = "preload:init" // The extension itself generates an event called preload:init that can be used to trigger preloads as soon as an object has been processed by htmx.
)

func PreloadOn[T any](hx htmx.HX[T], event PreloadEvent) T {
	return hx.Attr("preload", string(event))
}

// PreloadImages preloads linked image resources after an HTML page (or page fragment) is preloaded.
func PreloadImages[T any](hx htmx.HX[T], preloadImages bool) T {
	return hx.Attr("preload-images", boolToString(preloadImages))
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
