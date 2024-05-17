package removeme_test

import (
	"fmt"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/removeme"
)

var hx = htmx.NewStringAttrs()

func ExampleRemoveMe() {
	attr := removeme.RemoveMe(hx, time.Second)
	fmt.Println(attr)
	// Output: remove-me='1s'
}
