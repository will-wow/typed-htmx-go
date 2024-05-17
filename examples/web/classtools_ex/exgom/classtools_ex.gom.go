package exgom

import (
	"embed"
	"time"

	. "github.com/maragudk/gomponents/html"

	g "github.com/maragudk/gomponents"
	"github.com/will-wow/typed-htmx-go/htmx/ext/classtools"

	"github.com/will-wow/typed-htmx-go/examples/web/exprint"
	"github.com/will-wow/typed-htmx-go/examples/web/layout/gom/layout"

	"github.com/will-wow/typed-htmx-go/htmx"
)

var hx = htmx.NewGomponents()

//go:embed classtools_ex.gom.go
var fs embed.FS
var ex = exprint.New(fs, "//", "")

func Page() g.Node {
	return layout.Wrapper(
		"Class Tools",
		Class("class-tools-ex"),
		H1(g.Text("Class Tools")),
		P(
			g.Text("Demonstrates different uses of class-tools"),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("classtools_ex.gom.go", "demo")),
			),
		),
		H2(g.Text("Demo")),
		demo(),
	)
}

func demo() g.Node {
	//ex:start:demo
	return Div(
		hx.Ext(classtools.Extension),
		P(g.Text("Add then remove bold after 1 second, then toggle color every second"),
			classtools.Classes(hx,
				classtools.Add("bold", time.Second),
				classtools.Remove("bold", time.Second),
				classtools.Toggle("color", time.Second),
			),
		),
		P(g.Text("Add then remove bold after 1 second, while toggling color every second"),
			classtools.ClassesParallel(hx, []classtools.Run{
				{
					classtools.Add("bold", time.Second),
					classtools.Remove("bold", time.Second),
				},
				{
					classtools.Toggle("color", time.Second),
				},
			}),
		),
		P(g.Text("Add with no delay"),
			classtools.Classes(hx, classtools.Add("color", 0)),
		),
		P(g.Text("Toggle with 0 delay"),
			classtools.Classes(hx, classtools.Toggle("color", 0)),
		),
	)
	//ex:end:demo
}
