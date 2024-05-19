// package restored triggers an event restored whenever a back button even is detected while using [htmx.HX.boost].
package restored

import "github.com/will-wow/typed-htmx-go/htmx"

// Extension triggers an event restored whenever a back button even is detected while using hx-boost.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@1.9.12/dist/ext/restored.js"></script>
//
// # Usage
//
// A page utilizing hx-boost that will reload the h1 each time the back button is pressed:
//
//	<body { hx.Boost(true)... }>
//		<h1 { hx.Ext(restored.Extension)... } { hx.Trigger(restored.Event)... } { hx.Get("/header")... }>Come back!</h1>
//		<a href="/other_page">I'll be back</a>
//	</body>
//
// Extension: [restored]
//
// [restored]: https://htmx.org/extensions/restored/
const Extension htmx.Extension = "restored"

// Event is the event name that is triggered when the back button is pressed and [Extension] is enabled.
const Event = "restored"
