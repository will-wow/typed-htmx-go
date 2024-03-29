package htmx_test

import (
	"fmt"
	"testing"

	"github.com/will-wow/typed-htmx-go/htmx"
)

var templHx = htmx.NewTempl()

func BenchmarkTempl(b *testing.B) {
	attrs := make([]map[string]any, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		attrs[i] = templHx.Boost(true)
	}
}

func ExampleTemplAttrs() {
	attrs := htmx.TemplAttrs(
		templHx.Post("/example"),
		templHx.ReplaceURL(true),
	)

	fmt.Println(attrs)
	// Output: map[hx-post:/example hx-replace-url:true]
}
