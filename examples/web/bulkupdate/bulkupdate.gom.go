package bulkupdate

import (
	"time"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/swap"

	"github.com/will-wow/typed-htmx-go/examples/web/layout"
)

var ghx = htmx.NewGomponents()

type gomExample struct {
	layout layout.Gomponents
}

func newGomExample() gomExample {
	return gomExample{
		layout: layout.Gomponents{},
	}
}

func (e gomExample) page(users []userModel) g.Node {
	return e.layout.Base(
		"Bulk Update",
		Class("bulk-update"),
		H1(g.Text("Bulk Update")),
		P(
			g.Text("This demo shows how to implement a common pattern where rows are selected and then bulk updated. This is accomplished by putting a form around a table, with checkboxes in the table, and then including the checked values in the form submission"),
			Code(g.Text("(POST request)")),
			g.Text(":"),
		),
		e.table(users),
	)
}

func (e gomExample) table(users []userModel) g.Node {
	return FormEl(
		ID("checked-contacts"),
		ghx.Post("/examples/gomponents/bulk-update/"),
		ghx.SwapExtended(
			swap.New().Strategy(swap.OuterHTML).Settle(3*time.Second),
		),
		ghx.Target(
			"#toast",
		),

		H3(g.Text("Select Rows And Activate Or Deactivate Below")),
		Table(
			THead(
				Tr(
					Td(g.Text("Name")),
					Td(g.Text("Email")),
					Td(g.Text("Activate")),
				),
			),
			TBody(
				ID("tbody"),
				g.Group(
					g.Map(users, func(u userModel) g.Node {
						return Tr(
							Td(g.Text(u.Name)),
							Td(g.Text(u.Email)),
							Td(
								Input(
									Type("checkbox"),
									Name(u.Email),
									g.If(u.Active,
										Checked(),
									),
								),
							),
						)

					}),
				),
			),
		),
		Input(
			Type("submit"),
			Value("Submit"),
		),
		e.updateToast(""),
	)
}

func (e gomExample) updateToast(toast string) g.Node {
	return Span(
		ID("toast"),
		g.If(toast != "",
			g.Attr("aria-live", "polite"),
		),
		g.Text(toast),
	)
}
