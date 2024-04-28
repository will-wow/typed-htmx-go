package exgom

import (
	"embed"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/htmx/swap"

	"github.com/will-wow/typed-htmx-go/examples/web/exprint"
	"github.com/will-wow/typed-htmx-go/examples/web/layout/gom/layout"

	"github.com/will-wow/typed-htmx-go/htmx"
)

var hx = htmx.NewGomponents()

//go:embed progressbar.gom.go
var fs embed.FS
var ex = exprint.New(fs, "//", "")

func Page() g.Node {
	return layout.Wrapper(
		"Progress Bar",
		Class("progress-bar"),
		H1(g.Text("Progress Bar")),
		P(
			g.Text("This example shows how to implement a smoothly scrolling progress bar."),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("progressbar.gom.go", "demo")),
			),
		),
		H2(g.Text("Demo")),
		demo(),
	)
}

func demo() g.Node {
	//ex:start:demo
	return Div(
		hx.Target(htmx.TargetThis),
		hx.Swap(swap.OuterHTML),
		H3(g.Text("Start Progress")),
		Button(
			Class("btn"),
			hx.Post("/examples/templ/progress-bar/job/"),
			g.Text("Start Job"),
		),
	)
	//ex:end:demo
}
