package trigger

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx/internal/mod"
)

// Modifier is an enum of the possible hx-trigger modifiers.
type Modifier string

const (
	Once      Modifier = "once"      // the event will only trigger once (e.g. the first click)
	Changed   Modifier = "changed"   // the event will only change if the value of the element has changed. Please pay attention change is the name of the event and changed is the name of the modifier.
	Delay     Modifier = "delay"     // a delay will occur before an event triggers a request. If the event is seen again it will reset the delay.
	Throttle  Modifier = "throttle"  // a throttle will occur after an event triggers a request. If the event is seen again before the delay completes, it is ignored, the element will trigger at the end of the delay.
	From      Modifier = "from"      // allows the event that triggers a request to come from another element in the document (e.g. listening to a key event on the body, to support hot keys)
	Target    Modifier = "target"    // allows you to filter via a CSS selector on the target of the event. This can be useful when you want to listen for triggers from elements that might not be in the DOM at the point of initialization, by, for example, listening on the body, but with a target filter for a child element
	Consume   Modifier = "consume"   // if this option is included the event will not trigger any other htmx requests on parents (or on elements listening on parents)
	Queue     Modifier = "queue"     // determines how events are queued if an event occurs while a request for another event is in flight. Options are:
	Root      Modifier = "root"      // a CSS selector of the root element for intersection. Only used by the intersect event.
	Threshold Modifier = "threshold" // a floating point number between 0.0 and 1.0, indicating what amount of intersection to fire the event on. Only used by the intersect event.
)

// An Event is a builder to create a new user event hx-trigger.
type Event struct {
	event     string
	filter    string
	modifiers map[Modifier]string
}

// NewEvent starts a builder chain for creating a new hx-trigger for user events.
func NewEvent(event string) *Event {
	return &Event{
		event:     event,
		filter:    "",
		modifiers: map[Modifier]string{},
	}
}

// trigger is a no-op method to satisfy the Trigger interface.
func (e *Event) trigger() {}

// String returns the final hx-swap string.
func (e *Event) String() string {
	return mod.Join(e.coreEvent(), e.modifiers)
}

// OverrideEvent overrides the initial event name. This is useful if you are forking a default trigger setup.
func (e *Event) OverrideEvent(event string) *Event {
	e.event = event
	return e
}

// Filter specifies a boolean javascript expression as an event filter.
//
// If this expression evaluates to true the event will be triggered, otherwise it will be ignored.
//
//	<div hx-get="/clicked" hx-trigger="click[ctrlKey]">Control Click Me</div>
//
// Conditions can also refer to global functions or state
//
//	<div hx-get="/clicked" hx-trigger="click[checkGlobalState()]">Control Click Me</div>
//
// And can also be combined using the standard javascript syntax
//
//	<div hx-get="/clicked" hx-trigger="click[ctrlKey&&shiftKey]">Control-Shift Click Me</div>
//
// Note that all symbols used in the expression will be resolved first against the triggering event, and then next against the global namespace, so myEvent[foo] will first look for a property named foo on the event, then look for a global symbol with the name foo
func (e *Event) Filter(filter string) *Event {
	e.filter = filter
	return e
}

// Once makes the event will only trigger once (e.g. the first click)
func (e *Event) Once() *Event {
	e.modifiers[Once] = ""
	return e
}

// Changed makes the event only if the value of the element has changed. Please pay attention change is the name of the event and changed is the name of the modifier.
func (e *Event) Changed() *Event {
	e.modifiers[Changed] = ""
	return e
}

// Delay will cause a delay before an event triggers a request. If the event is seen again it will reset the delay.
func (e *Event) Delay(timing time.Duration) *Event {
	e.modifiers[Delay] = timing.String()
	return e
}

// Throttle will cause a throttle to occur after an event triggers a request. If the event is seen again before the delay completes, it is ignored, the element will trigger at the end of the delay.
func (e *Event) Throttle(timing time.Duration) *Event {
	e.modifiers[Throttle] = timing.String()
	return e
}

// A SelectorModifier is a relative modifier to a CSS selector. This is used for "extended selectors".
// Some attributes only support a subset of these, but any Relative function that takes this type supports the full set..
type SelectorModifier string

const (
	Closest  SelectorModifier = "closest"  // find the closest ancestor element or itself, that matches the given CSS selector
	Find     SelectorModifier = "find"     // find the first child descendant element that matches the given CSS selector
	Next     SelectorModifier = "next"     // scan the DOM forward for the first element that matches the given CSS selector. (e.g. next .error will target the closest following sibling element with error class)
	Previous SelectorModifier = "previous" // scan the DOM backwards fo
)

// A FromSelector is a non-standard selector for the From modifier.
type FromSelector string

const (
	FromDocument FromSelector = "document" // listen for events on the document
	FromWindow   FromSelector = "window"   // listen for events on the window
	FromNext     FromSelector = "next"     // resolves to element.nextElementSibling
	FromPrevious FromSelector = "previous" // resolves to element.previousElementSibling
)

// FromRelative creates a relative selector for an Event.From modifier.
// It always wraps the selector in (), in case it contains a space.
func FromRelative(modifier SelectorModifier, selector string) FromSelector {
	return FromSelector(fmt.Sprintf("%s (%s)", modifier, selector))
}

var disambiguatedRe = regexp.MustCompile(`\(`)

// From allows the event that triggers a request to come from another element in the document (e.g. listening to a key event on the body, to support hot keys)
// A standard CSS selector resolves to all elements matching that selector. Thus, from:input would listen on every input on the page.
// If the selector contains whitespace, it will be wrapped in () to disambiguate it from other modifiers.
func (e *Event) From(extendedSelector FromSelector) *Event {
	// Wrap the selector in () to disambiguate it, if not done already by [FromRelative].
	var selector string
	if disambiguatedRe.MatchString(string(extendedSelector)) {
		selector = string(extendedSelector)
	} else {
		selector = fmt.Sprintf("(%s)", extendedSelector)
	}
	e.modifiers[From] = selector
	return e
}

// Target allows you to filter via a CSS selector on the target of the event. This can be useful when you want to listen for triggers from elements that might not be in the DOM at the point of initialization, by, for example, listening on the body, but with a target filter for a child element.
// If the selector contains whitespace, it will be wrapped in () to disambiguate it from other modifiers.
func (e *Event) Target(selector string) *Event {
	e.modifiers[Target] = disambiguateSelector(selector)
	return e
}

// Consume causes the event not to trigger any other htmx requests on parents (or on elements listening on parents).
func (e *Event) Consume() *Event {
	e.modifiers[Consume] = ""
	return e
}

// A QueueOption determines how events are queued if an event occurs while a request for another event is in flight.
type QueueOption string

const (
	First QueueOption = "first" // queue the first event
	Last  QueueOption = "last"  // queue the last event (default)
	All   QueueOption = "all"   // queue all events (issue a request for each event)
	None  QueueOption = "none"  // do not queue new events
)

// Queue determines how events are queued if an event occurs while a request for another event is in flight.
func (e *Event) Queue(option QueueOption) *Event {
	e.modifiers[Queue] = string(option)
	return e
}

// Clear removes a modifier entirely from the builder.
// Used to undo an previously set modifier.
func (s *Event) Clear(modifier Modifier) *Event {
	delete(s.modifiers, modifier)
	return s
}

// coreEvent returns the event name with the filter appended, if present.
func (e *Event) coreEvent() string {
	if e.filter == "" {
		return e.event
	}
	return fmt.Sprintf("%s[%s]", e.event, e.filter)
}

// An IntersectEvent fires once when an element first intersects the viewport. This supports additional options to a normal trigger, [IntersectEvent.Root] and [IntersectEvent.Threshold].
type IntersectEvent struct {
	Event
}

// NewIntersectEvent configures a trigger that fires once when an element first intersects the viewport. This supports additional options to a normal trigger, [IntersectEvent.Root] and [IntersectEvent.Threshold].
func NewIntersectEvent() *IntersectEvent {
	return &IntersectEvent{
		Event{
			event:     "intersect",
			filter:    "",
			modifiers: map[Modifier]string{},
		},
	}
}

// Root configures a CSS selector of the root element for intersection.
func (e *IntersectEvent) Root(selector string) *IntersectEvent {
	e.modifiers[Root] = disambiguateSelector(selector)
	return e
}

// Threshold takes a floating point number between 0.0 and 1.0, indicating what amount of intersection to fire the event on
func (e *IntersectEvent) Threshold(threshold float64) *IntersectEvent {
	e.modifiers[Threshold] = strconv.FormatFloat(threshold, 'f', -1, 64)
	return e
}

// disambiguateSelector surrounds a selector with parentheses if it contains a space, for the from and target modifiers.
func disambiguateSelector(selector string) string {
	if strings.Contains(selector, " ") {
		return fmt.Sprintf("(%s)", selector)
	}
	return selector
}
