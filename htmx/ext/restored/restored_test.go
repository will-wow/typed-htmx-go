package restored_test

import (
	"fmt"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/restored"
)

var hx = htmx.NewStringAttrs()

func ExampleExtension() {
	attr := hx.Ext(restored.Extension)
	fmt.Println(attr)
	// Output: hx-ext='restored'
}

func ExampleEvent() {
	attr := hx.Trigger(restored.Event)
	fmt.Println(attr)
	// Output: hx-trigger='restored'
}
