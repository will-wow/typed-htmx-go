package htmx_test

import (
	"os"
	"strings"
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

func TestRender(t *testing.T) {
	tests := []struct {
		name string
		node htmx.GomponentsAttrs
		want string
	}{
		{
			name: "string value",
			node: gomHx.Boost(true),
			want: ` hx-boost="true"`,
		},
		{
			name: "bool value",
			node: gomHx.Preserve(),
			want: ` hx-preserve`,
		},
		{
			name: "invalid value",
			node: htmx.GomponentsAttrs{},
			want: ``,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check that render renders the expected value to a writer.
			builder := strings.Builder{}
			err := tt.node.Render(&builder)
			if err != nil {
				t.Errorf("got error %v, want %s", err, tt.want)
			}
			got := builder.String()
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
			// Check that .String() returns the same expected value.
			string := tt.node.String()
			if string != tt.want {
				t.Errorf("for String got %s, want %s", string, tt.want)
			}
		})
	}

}
