package debug_test

import (
	"fmt"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/debug"
)

var hx = htmx.NewStringAttrs()

func ExampleExtension() {
	attr := hx.Ext(debug.Extension)
	fmt.Println(attr)
	// Output: hx-ext='debug'
}
