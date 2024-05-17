package classtools_test

import (
	"fmt"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/classtools"
)

var hx = htmx.NewStringAttrs()

func ExampleClasses() {
	attr := classtools.Classes(hx,
		// Add foo after 500ms
		classtools.Add("foo", 500*time.Millisecond),
		// Remove bar immediately after
		classtools.Remove("bar", 0),
		// Then, start toggling baz every second
		classtools.Toggle("baz", time.Second),
	)
	fmt.Println(attr)
	// Output: classes='add foo:500ms, remove bar:0s, toggle baz:1s'
}

func ExampleClassesParallel() {
	attr := classtools.ClassesParallel(hx, []classtools.Run{
		{
			// Add foo after 500ms
			classtools.Add("foo", 500*time.Millisecond),
			// Remove bar immediately after
			classtools.Remove("bar", 0),
		},
		{
			// Also, toggle baz every second
			classtools.Toggle("baz", time.Second),
		},
	})
	fmt.Println(attr)
	// Output: classes='add foo:500ms, remove bar:0s & toggle baz:1s'
}
