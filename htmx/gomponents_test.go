package htmx_test

import (
	"os"
	"testing"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/swap"
)

var gomHx = htmx.NewGomponents()

func BenchmarkBoost(b *testing.B) {
	attrs := make([]htmx.GomponentsAttrs, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		attrs[i] = gomHx.Boost(true)
	}
}

func ExampleNewGomponents() {
	hx := htmx.NewGomponents()

	component := FormEl(
		hx.Boost(true),
		hx.Post("/submit"),
		hx.Swap(swap.OuterHTML),
		ID("form"),
		Input(
			Name("firstName"),
		),
		Button(
			Type("submit"),
			g.Text("Submit"),
		),
	)

	_ = component.Render(os.Stdout)
	// Output: <form hx-boost="true" hx-post="/submit" hx-swap="outerHTML" id="form"><input name="firstName"><button type="submit">Submit</button></form>
}
