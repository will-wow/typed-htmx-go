// package swap provides a builder for the hx-swap attribute value.
package swap

import (
	"fmt"

	"github.com/will-wow/typed-htmx-go/hx/internal/mod"
)

// Modifier is an enum of the possible hx-swap modifiers.
type Modifier string

const (
	Transition  Modifier = "transition"
	Swap        Modifier = "swap"
	Settle      Modifier = "settle"
	IgnoreTitle Modifier = "ignoreTitle"
	Scroll      Modifier = "scroll"
	Show        Modifier = "show"
	FocusScroll Modifier = "focus-scroll"
)

// Builder is a builder to create a new hx-swap attribute.
type Builder struct {
	strategy  Strategy
	modifiers map[Modifier]string
}

// New starts a builder chain for creating a new hx-swap attribute.
// It contains methods to add and remove modifiers from the swap.
// Subsequent calls can override previous modifiers of the same type - for instance, .Scroll(Top).Scroll(Bottom) will result in `hx-swap="scroll:bottom"`.
// Call .End() to get the final hx-swap string.
func New() *Builder {
	return &Builder{
		strategy:  InnerHTML,
		modifiers: map[Modifier]string{},
	}
}

// String returns the final hx-swap string.
func (s *Builder) String() string {
	return mod.Join(string(s.strategy), s.modifiers)
}

// Strategy specifies how the response will be swapped in relative to the target of an AJAX request.
type Strategy string

const (
	InnerHTML   Strategy = "innerHTML"   // Replace the inner html of the target element
	OuterHTML   Strategy = "outerHTML"   // Replace the entire target element with the response
	BeforeBegin Strategy = "beforebegin" // Insert the response before the target element
	AfterBegin  Strategy = "afterbegin"  // Insert the response before the first child of the target element
	BeforeEnd   Strategy = "beforeend"   // Insert the response after the last child of the target element
	AfterEnd    Strategy = "afterend"    // Insert the response after the target element
	Delete      Strategy = "delete"      // Deletes the target element regardless of the response
	None        Strategy = "none"        // Does not append content from response (out of band items will still be processed).
)

// Strategy allows you to specify how the response will be swapped in relative to the target of an AJAX request. If you do not specify the option, the default is htmx.config.defaultSwapStyle (innerHTML).
func (s *Builder) Strategy(strategy Strategy) *Builder {
	s.strategy = strategy
	return s
}

// Transition enables the new View Transitions API when a swap occurs.
// You can also enable this feature globally by setting the htmx.config.globalViewTransitions config setting to true.
func (s *Builder) Transition() *Builder {
	s.modifiers[Transition] = "true"
	return s
}

// SwapTiming modifies the amount of time that htmx will wait after receiving a response to swap the content.
// This attribute can be used to synchronize htmx with the timing of CSS transition effects.
func (s *Builder) SwapTiming(wait string) *Builder {
	s.modifiers[Swap] = wait
	return s
}

// SettleTiming modifies the time between the swap and the settle logic.
// This attribute can be used to synchronize htmx with the timing of CSS transition effects.
func (s *Builder) SettleTiming(wait string) *Builder {
	s.modifiers[Settle] = wait
	return s
}

// IgnoreTitle turns off the default title behavior,
// where htmx will update the title of the page if it finds a <title> tag in the response content.
func (s *Builder) IgnoreTitle() *Builder {
	s.modifiers[IgnoreTitle] = "true"
	return s
}

// ScrollDirection specifies the direction to scroll/show an element after a swap.
type ScrollDirection string

const (
	Top    ScrollDirection = "top"
	Bottom ScrollDirection = "bottom"
)

// Scroll will scroll the target element to the top/bottom after the swap.
func (s *Builder) Scroll(scrollDirection ScrollDirection) *Builder {
	s.modifiers[Scroll] = string(scrollDirection)
	return s
}

// ScrollElement will scroll the selected element to the top/bottom after the swap.
// The selector is a CSS selector that identifies the element to scroll.
func (s *Builder) ScrollElement(selector string, scrollDirection ScrollDirection) *Builder {
	s.modifiers[Scroll] = fmt.Sprintf("%s:%s", selector, scrollDirection)
	return s
}

// Show will scroll the viewport to show the target element after the swap.
func (s *Builder) Show(scrollDirection ScrollDirection) *Builder {
	s.modifiers[Show] = string(scrollDirection)
	return s
}

// ShowElement will scroll the viewport to show the selected element after the swap.
// The selector is a CSS selector that identifies the element to show.
func (s *Builder) ShowElement(selector string, scrollDirection ScrollDirection) *Builder {
	s.modifiers[Show] = fmt.Sprintf("%s:%s", selector, scrollDirection)
	return s
}

// ShowWindow will scroll the viewport to the top/bottom after the swap.
func (s *Builder) ShowWindow(scrollDirection ScrollDirection) *Builder {
	s.modifiers[Show] = fmt.Sprintf("window:%s", scrollDirection)
	return s
}

// ShowNone will disable the default show:top behavior for boosted links and forms.
// You can disable it globally with htmx.config.scrollIntoViewOnBoost, or you can use hx-swap="show:none" on an element basis.
func (s *Builder) ShowNone() *Builder {
	s.modifiers[Show] = "none"
	return s
}

// FocusScroll overrides the behavior of scrolling of a focused element after a swap.
//
// htmx preserves focus between requests for inputs that have a defined id attribute. By default htmx prevents auto-scrolling to focused inputs between requests which can be unwanted behavior on longer requests when the user has already scrolled away. To enable focus scroll you can use focus-scroll:true.
//
// Alternatively, if you want the page to automatically scroll to the focused element after each request you can change the htmx global configuration value htmx.config.defaultFocusScroll to true. Then disable it for specific requests using focus-scroll:false.
func (s *Builder) FocusScroll(value bool) *Builder {
	s.modifiers[FocusScroll] = boolToString(value)
	return s
}

// Clear removes a modifier entirely from the builder.
// Used to undo an previously set modifier.
func (s *Builder) Clear(modifier Modifier) *Builder {
	delete(s.modifiers, modifier)
	return s
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
