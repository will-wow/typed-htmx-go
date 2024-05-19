package eventheader_test

import (
	"fmt"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/eventheader"
)

var hx = htmx.NewStringAttrs()

func ExampleExtension() {
	attr := hx.Ext(eventheader.Extension)
	fmt.Println(attr)
	// Output: hx-ext='event-header'
}
