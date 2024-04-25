package swap_test

import (
	"fmt"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/swap"
)

var hx = htmx.NewStringAttrs()

func ExampleNew_default() {
	fmt.Println(
		hx.SwapExtended(swap.New()),
	)
	// output: hx-swap='innerHTML'
}

func ExampleBuilder_Strategy() {
	fmt.Println(
		hx.SwapExtended(swap.New().Strategy(swap.OuterHTML)),
	)
	// output: hx-swap='outerHTML'
}

func ExampleBuilder_Transition() {
	fmt.Println(
		hx.SwapExtended(swap.New().Transition()),
	)
	// output: hx-swap='innerHTML transition:true'
}

func ExampleBuilder_Swap_timing() {
	fmt.Println(
		hx.SwapExtended(swap.New().Swap(500 * time.Millisecond)),
	)
	// output: hx-swap='innerHTML swap:500ms'
}

func ExampleBuilder_Settle_timing() {
	fmt.Println(
		hx.SwapExtended(swap.New().Settle(500 * time.Millisecond)),
	)
	// output: hx-swap='innerHTML settle:500ms'
}

func ExampleBuilder_IgnoreTitle() {
	fmt.Println(
		hx.SwapExtended(swap.New().IgnoreTitle()),
	)
	// output: hx-swap='innerHTML ignoreTitle:true'
}

func ExampleBuilder_Scroll() {
	fmt.Println(
		hx.SwapExtended(swap.New().Scroll(swap.Top)),
	)
	// output: hx-swap='innerHTML scroll:top'
}

func ExampleBuilder_ScrollElement() {
	fmt.Println(
		hx.SwapExtended(swap.New().ScrollElement("#example", swap.Top)),
	)
	// output: hx-swap='innerHTML scroll:#example:top'
}

func ExampleBuilder_Show() {
	fmt.Println(
		hx.SwapExtended(swap.New().Show(swap.Bottom)),
	)
	// output: hx-swap='innerHTML show:bottom'
}

func ExampleBuilder_ShowElement() {
	fmt.Println(
		hx.SwapExtended(swap.New().ShowElement("#example", swap.Bottom)),
	)
	// output: hx-swap='innerHTML show:#example:bottom'
}

func ExampleBuilder_Show_window() {
	fmt.Println(
		hx.SwapExtended(swap.New().ShowElement(swap.ShowWindow, swap.Top)),
	)
	// output: hx-swap='innerHTML show:window:top'
}

func ExampleBuilder_ShowNone() {
	fmt.Println(
		hx.SwapExtended(swap.New().ShowNone()),
	)
	// output: hx-swap='innerHTML show:none'
}

func ExampleBuilder_FocusScroll() {
	fmt.Println(
		hx.SwapExtended(swap.New().FocusScroll(true)),
	)
	// output: hx-swap='innerHTML focus-scroll:true'
}

func ExampleBuilder_FocusScroll_disable() {
	fmt.Println(
		hx.SwapExtended(swap.New().FocusScroll(false)),
	)
	// output: hx-swap='innerHTML focus-scroll:false'
}

func ExampleBuilder_Clear() {
	fmt.Println(
		hx.SwapExtended(swap.New().Transition().Clear(swap.Transition)),
	)
	// output: hx-swap='innerHTML'
}
