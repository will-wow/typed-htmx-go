// package sse connects to an EventSource directly from HTML. It manages the connections to your web server, listens for server events, and then swaps their contents into your htmx webpage in real-time.
//
// [EventSource]: https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
package sse

import (
	"fmt"

	"github.com/will-wow/typed-htmx-go/htmx"
)

// Extension connects to an EventSource directly from HTML. It manages the connections to your web server, listens for server events, and then swaps their contents into your htmx webpage in real-time.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@1.9.12/dist/ext/sse.js"></script>
//
// Extension: [server-sent-events]
//
// [server-sent-events]: https://htmx.org/extensions/server-sent-events/
//
// [EventSource]: https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
const Extension htmx.Extension = "sse"

// Message is the the default name of an empty SSE event.
const Message = "message"

// Connect
func Connect[T any](hx htmx.HX[T], url string) T {
	return hx.Attr("sse-connect", url)
}

// Swap
func Swap[T any](hx htmx.HX[T], messageName string) T {
	return hx.Attr("sse-swap", messageName)
}

// Trigger allows SSE messages to trigger HTTP callbacks using the [htmx.HX.Trigger()] attribute.
func Trigger[T any](hx htmx.HX[T], event string) T {
	prefixedEvent := fmt.Sprintf("sse:%s", event)
	return hx.Attr(htmx.Trigger, prefixedEvent)
}
