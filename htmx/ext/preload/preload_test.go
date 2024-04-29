package preload_test

import (
	"fmt"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/preload"
)

var hx = htmx.NewStringAttrs()

func ExamplePreload() {
	attr := preload.Preload(hx)
	fmt.Println(attr)
	// Output: preload
}

func ExamplePreloadOn() {
	attr := preload.PreloadOn(hx, preload.MouseOver)
	fmt.Println(attr)
	// Output: preload='mouseover'
}

func ExamplePreloadOn_init() {
	attr := preload.PreloadOn(hx, preload.Init)
	fmt.Println(attr)
	// Output: preload='preload:init'
}

func ExamplePreloadImages() {
	attr := preload.PreloadImages(hx, true)
	fmt.Println(attr)
	// Output: preload-images='true'
}
