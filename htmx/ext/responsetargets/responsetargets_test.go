package responsetargets_test

import (
	"fmt"
	"net/http"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/responsetargets"
)

var hx = htmx.NewStringAttrs()

func ExampleTarget_code() {
	attr := responsetargets.Target(hx, responsetargets.Code(http.StatusNotFound), htmx.TargetRelative(htmx.Next, "div"))
	fmt.Println(attr)
	// Output: hx-target-404='next div'
}

func ExampleTarget_error() {
	attr := responsetargets.Target(hx, responsetargets.Error, htmx.TargetThis)
	fmt.Println(attr)
	// Output: hx-target-error='this'
}

func ExampleTarget_wildcard() {
	attr := responsetargets.Target(hx, responsetargets.Wildcard(4, 0), htmx.TargetRelative(htmx.Next, "div"))
	fmt.Println(attr)
	// Output: hx-target-40*='next div'
}

func ExampleTarget_wildcardX() {
	attr := responsetargets.Target(hx, responsetargets.WildcardX(4, 0), htmx.TargetRelative(htmx.Next, "div"))
	fmt.Println(attr)
	// Output: hx-target-40x='next div'
}
