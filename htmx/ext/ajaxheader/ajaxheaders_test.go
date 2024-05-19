package ajaxheader_test

import (
	"fmt"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/ajaxheader"
)

var hx = htmx.NewStringAttrs()

func ExampleExtension() {
	attr := hx.Ext(ajaxheader.Extension)
	fmt.Println(attr)
	// Output: hx-ext='ajax-header'
}
