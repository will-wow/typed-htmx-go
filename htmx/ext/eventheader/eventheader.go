// package eventheader the Triggering-Event header to requests. The value of the header is a JSON serialized version of the event that triggered the request.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@1.9.12/dist/ext/event-header.js"></script>
//
// # Usage
//
//	<button { hx.Ext(eventheader.Extension)... } >
//		Click Me!
//	</button>
//
// Sends something like this:
//
//	Triggering-Event: '{ "isTrusted": false, "htmx-internal-data": { "handled": true }, "screenX": 0, "screenY": 0, "clientX": 0, "clientY": 0, "ctrlKey": false, "shiftKey": false, "altKey": false, "metaKey": false, "button": 0, "buttons": 0, "relatedTarget": null, "pageX": 0, "pageY": 0, "x": 0, "y": 0, "offsetX": 0, "offsetY": 0, "movementX": 0, "movementY": 0, "fromElement": null, "toElement": "button", "layerX": 0, "layerY": 0, "view": "Window", "detail": 0, "sourceCapabilities": null, "which": 1, "NONE": 0, "CAPTURING_PHASE": 1, "AT_TARGET": 2, "BUBBLING_PHASE": 3, "type": "click", "target": "button", "currentTarget": "button", "eventPhase": 2, "bubbles": true, "cancelable": true, "defaultPrevented": true, "composed": true, "timeStamp": 188.86999995447695, "srcElement": "button", "returnValue": false, "cancelBubble": false, "path": [ "button", "div#work-area", "body", "html", "Node", "Window" ] }'
//
// Extension: [event-header]
//
// [event-header]: https://htmx.org/extensions/event-header/
package eventheader

import "github.com/will-wow/typed-htmx-go/htmx"

const Extension htmx.Extension = "event-header"
