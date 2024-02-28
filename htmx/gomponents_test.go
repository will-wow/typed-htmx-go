package htmx_test

import (
	"testing"

	"github.com/will-wow/typed-htmx-go/htmx"
)

var gomHx = htmx.NewGomponents()

func BenchmarkBoost(b *testing.B) {
	attrs := make([]htmx.GomponentsAttrs, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		attrs[i] = gomHx.Boost(true)
	}
}
