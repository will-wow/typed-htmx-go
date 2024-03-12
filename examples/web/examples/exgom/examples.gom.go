package exgom

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/examples/web/layout/gom/layout"
)

func Page() g.Node {
	return layout.Wrapper("",
		H1(g.Text("UI Gomponents Examples")),
		P(g.Text("Below are a set of UX patterns implemented in htmx with minimal HTML and styling.")),
		P(
			g.Text("These are ported from the "),
			A(Href("https://htmx.org/examples/"), Target("_blank"), Rel("noopener"), g.Text("htmx examples")),
			g.Text(" and are intended showcase the use of "),
			Code(g.Text("hx")),
			g.Text(" when building HTMX applications."),
		),
		P(
			g.Text("You can copy and paste them and then adjust them for your needs."),
		),
		Table(
			THead(
				Tr(
					Th(g.Text("Pattern")),
					Th(g.Text("Description")),
				),
			),
			TBody(
				exampleRow(
					"/examples/gomponents/click-to-edit",
					"Click To Edit",
					"Demonstrates inline editing of a data object",
				),
				exampleRow(
					"/examples/gomponents/bulk-update",
					"Bulk Update",
					"Demonstrates bulk updating of multiple rows of data",
				),
				exampleRow(
					"/examples/gomponents/active-search",
					"Active Search",
					"Demonstrates the active search box pattern",
				),
			),
		),
	)
}

func exampleRow(link, name, description string) g.Node {
	return Tr(
		Td(A(Href(link), g.Text(name))),
		Td(g.Text(description)),
	)
}
