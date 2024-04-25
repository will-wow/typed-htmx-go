package htmx_test

import (
	"fmt"

	"github.com/will-wow/typed-htmx-go/htmx"
)

func ExampleNewStringAttrs() {
	hx := htmx.NewStringAttrs()

	fmt.Println(hx.Target(htmx.TargetNext))
	fmt.Println(hx.Preserve())

	// Output:
	// hx-target='next'
	// hx-preserve
}
