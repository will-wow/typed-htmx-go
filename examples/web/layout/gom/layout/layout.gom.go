package layout

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/htmx"
)

var ghx = htmx.NewGomponents()

func Wrapper(title string, children ...g.Node) g.Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				Meta(Charset("utf-8")),

				g.If(title != "",
					TitleEl(g.Textf("%s | HX | Examples", title)),
				),
				g.If(title == "",
					TitleEl(g.Text("HX | Examples")),
				),

				Meta(Name("htmx-config"), Content(`{"includeIndicatorStyles":false}`)),
				Meta(Name("color-scheme"), Content("light")),
				Meta(Name("description"), Content("Examples of typed-htmx-go/hx")),
				Meta(Name("referrer"), Content("origin-when-cross-origin")),
				Meta(Name("creator"), Content("Will Ockelmann-Wagner")),
				Script(Src("https://unpkg.com/htmx.org@1.9.10")),
				Link(Rel("stylesheet"), Href("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.classless.min.css")),
				Link(Rel("stylesheet"), Href("https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.9.0/build/styles/default.min.css")),
				Link(Rel("stylesheet"), Href("/static/main.css")),
			),
			Body(
				ghx.Boost(true),
				Main(
					nav(),
					g.Group(children),
				),
				Script(Src("https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.9.0/build/highlight.min.js")),
				Script(Src("https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.9.0/build/languages/go.min.js")),
				Script(Src("/static/main.js")),
			),
		),
	)
}

func nav() g.Node {
	return Nav(
		Ul(
			Li(
				A(
					Href("/examples/gomponents/"),
					Strong(
						g.Text("T"),
						U(g.Text("HX")),
						g.Text("GO"),
					),
				),
			),
		),
		Ul(
			Li(
				A(
					Href("https://pkg.go.dev/github.com/will-wow/typed-htmx-go/hx"),
					Target("_blank"),
					Rel("noopener"),
					g.Text("Docs"),
				),
			),
			Li(
				A(
					Href("https://htmx.org"),
					Target("_blank"),
					Rel("noopener"),
					g.Text("HTMX"),
				),
			),
			Li(
				A(
					Href("/"),
					g.Text("Templ"),
				),
			),
			Li(
				A(
					Href("/examples/gomponents/"),
					g.Text("Gomponents"),
				),
			),
			Li(
				A(
					Href("https://github.com/will-wow/typed-htmx-go"),
					Target("_blank"),
					Rel("noopener"),
					g.Text("GitHub"),
				),
			),
		),
	)
}
