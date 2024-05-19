package alpinemorph_test

import (
	"fmt"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/alpinemorph"
)

var hx = htmx.NewStringAttrs()

func ExampleExtension() {
	attr := hx.Ext(alpinemorph.Extension)
	fmt.Println(attr)
	// Output: hx-ext='alpine-morph'
}
