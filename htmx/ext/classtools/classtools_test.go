package classtools_test

import (
	"fmt"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/classtools"
)

var hx = htmx.NewStringAttrs()

func ExampleClasses() {
	attr := classtools.Classes(hx, []classtools.Run{
		{
			classtools.Add("foo"),
			classtools.Remove("bar", time.Millisecond*500),
		},
		{
			classtools.Toggle("baz", time.Second),
		},
	})
	fmt.Println(attr)
	// Output: classes='add foo, remove bar:500ms & toggle baz:1s'
}
