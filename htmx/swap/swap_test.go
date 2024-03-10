package swap_test

import (
	"fmt"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx/swap"
)

func ExampleNew_default() {
	builder := swap.New()
	fmt.Println(builder.String())
	// Output: innerHTML
}

func ExampleBuilder_Strategy() {
	builder := swap.New().Strategy(swap.OuterHTML)
	fmt.Println(builder.String())
	// Output: outerHTML
}

func ExampleBuilder_Transition() {
	builder := swap.New().Transition()
	fmt.Println(builder.String())
	// Output: innerHTML transition:true
}

func ExampleBuilder_Swap_timing() {
	builder := swap.New().Swap(500 * time.Millisecond)
	fmt.Println(builder.String())
	// Output: innerHTML swap:500ms
}

func ExampleBuilder_Settle_timing() {
	builder := swap.New().Settle(500 * time.Millisecond)
	fmt.Println(builder.String())
	// Output: innerHTML settle:500ms
}

func ExampleBuilder_IgnoreTitle() {
	builder := swap.New().IgnoreTitle()
	fmt.Println(builder.String())
	// Output: innerHTML ignoreTitle:true
}

func ExampleBuilder_Scroll() {
	builder := swap.New().Scroll(swap.Top)
	fmt.Println(builder.String())
	// Output: innerHTML scroll:top
}

func ExampleBuilder_ScrollElement() {
	builder := swap.New().ScrollElement("#example", swap.Top)
	fmt.Println(builder.String())
	// Output: innerHTML scroll:#example:top
}

func ExampleBuilder_Show() {
	builder := swap.New().Show(swap.Bottom)
	fmt.Println(builder.String())
	// Output: innerHTML show:bottom
}

func ExampleBuilder_ShowElement() {
	builder := swap.New().ShowElement("#example", swap.Bottom)
	fmt.Println(builder.String())
	// Output: innerHTML show:#example:bottom
}

func ExampleBuilder_Show_window() {
	builder := swap.New().ShowElement(swap.ShowWindow, swap.Top)
	fmt.Println(builder.String())
	// Output: innerHTML show:window:top
}

func ExampleBuilder_ShowNone() {
	builder := swap.New().ShowNone()
	fmt.Println(builder.String())
	// Output: innerHTML show:none
}

func ExampleBuilder_FocusScroll() {
	builder := swap.New().FocusScroll(true)
	fmt.Println(builder.String())
	// Output: innerHTML focus-scroll:true
}

func ExampleBuilder_Clear() {
	builder := swap.New().Transition().Clear(swap.Transition)
	fmt.Println(builder.String())
	// Output: innerHTML
}
