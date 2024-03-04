// package on holds constants for the camel-cased HTMX event names.
package on

// An Event is a camel-cased HTMX event name, used as an argument to the [htmx.HX.On] attribute.
type Event = string

const (
	Abort                 Event = "htmx:abort"                    // send this event to an element to abort a request
	AfterOnLoad           Event = "htmx:after-on-load"            // triggered after an AJAX request has completed processing a successful response
	AfterProcessNode      Event = "htmx:after-process-node"       // triggered after htmx has initialized a node
	AfterRequest          Event = "htmx:after-request"            // triggered after an AJAX request has completed
	AfterSettle           Event = "htmx:after-settle"             // triggered after the DOM has settled
	AfterSwap             Event = "htmx:after-swap"               // triggered after new content has been swapped in
	BeforeCleanupElement  Event = "htmx:before-cleanup-element"   // triggered before htmx disables an element or removes it from the DOM
	BeforeOnLoad          Event = "htmx:before-on-load"           // triggered before any response processing occurs
	BeforeProcessNode     Event = "htmx:before-process-node"      // triggered before htmx initializes a node
	BeforeRequest         Event = "htmx:before-request"           // triggered before an AJAX request is made
	BeforeSwap            Event = "htmx:before-swap"              // triggered before a swap is done, allows you to configure the swap
	BeforeSend            Event = "htmx:before-send"              // triggered just before an ajax request is sent
	ConfigRequest         Event = "htmx:config-request"           // triggered before the request, allows you to customize parameters, headers
	Confirm               Event = "htmx:confirm"                  // triggered after a trigger occurs on an element, allows you to cancel (or delay) issuing the AJAX request
	HistoryCacheError     Event = "htmx:history-cache-error"      // triggered on an error during cache writing
	HistoryCacheMiss      Event = "htmx:history-cache-miss"       // triggered on a cache miss in the history subsystem
	HistoryCacheMissError Event = "htmx:history-cache-miss-error" // triggered on a unsuccessful remote retrieval
	HistoryCacheMissLoad  Event = "htmx:history-cache-miss-load"  // triggered on a successful remote retrieval
	HistoryRestore        Event = "htmx:history-restore"          // triggered when htmx handles a history restoration action
	BeforeHistorySave     Event = "htmx:before-history-save"      // triggered before content is saved to the history cache
	Load                  Event = "htmx:load"                     // triggered when new content is added to the DOM
	NoSSESourceError      Event = "htmx:no-sse-source-error"      // triggered when an element refers to a SSE event in its trigger, but no parent SSE source has been defined
	OnLoadError           Event = "htmx:on-load-error"            // triggered when an exception occurs during the onLoad handling in htmx
	OOBAfterSwap          Event = "htmx:oob-after-swap"           // triggered after an out of band element as been swapped in
	OOBBeforeSwap         Event = "htmx:oob-before-swap"          // triggered before an out of band element swap is done, allows you to configure the swap
	OOBErrorNoTarget      Event = "htmx:oob-error-no-target"      // triggered when an out of band element does not have a matching ID in the current DOM
	Prompt                Event = "htmx:prompt"                   // triggered after a prompt is shown
	PushedIntoHistory     Event = "htmx:pushed-into-history"      // triggered after an url is pushed into history
	ResponseError         Event = "htmx:response-error"           // triggered when an HTTP response error (non-200 or 300 response code) occurs
	SendError             Event = "htmx:send-error"               // triggered when a network error prevents an HTTP request from happening
	SSEError              Event = "htmx:sse-error"                // triggered when an error occurs with a SSE source
	SSEOpen               Event = "htmx:sse-open"                 // triggered when a SSE source is opened
	SwapError             Event = "htmx:swap-error"               // triggered when an error occurs during the swap phase
	TargetError           Event = "htmx:target-error"             // triggered when an invalid target is specified
	Timeout               Event = "htmx:timeout"                  // triggered when a request timeout occurs
	ValidationValidate    Event = "htmx:validation:validate"      // triggered before an element is validated
	ValidationFailed      Event = "htmx:validation:failed"        // triggered when an element fails validation
	ValidationHalted      Event = "htmx:validation:halter"        // triggered when a request is halted due to validation errors
	XHRAbort              Event = "htmx:xhr:abort"                // triggered when an ajax request aborts
	XHRLoadEnd            Event = "htmx:xhr:loadend"              // triggered when an ajax request ends
	XHRLoadStart          Event = "htmx:xhr:loadstart"            // triggered when an ajax request starts
	XHRProgress           Event = "htmx:xhr:progress"             // triggered periodically during an ajax request that supports progress events
)
