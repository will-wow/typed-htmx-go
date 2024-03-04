package exgom

import (
	"time"

	"github.com/lithammer/dedent"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/swap"

	"github.com/will-wow/typed-htmx-go/examples/web/bulkupdate/form"

	"github.com/will-wow/typed-htmx-go/examples/web/layout/gom/layout"
)

var hx = htmx.NewGomponents()

func Page(users []form.UserModel) g.Node {
	return layout.Wrapper(
		"Bulk Update",
		Class("bulk-update"),
		H1(g.Text("Bulk Update")),
		P(
			g.Text("This demo shows how to implement a common pattern where rows are selected and then bulk updated. This is accomplished by putting a form around a table, with checkboxes in the table, and then including the checked values in the form submission"),
			Code(g.Text("(POST request)")),
			g.Text(":"),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(dedent.Dedent(`
					func table(users []form.UserModel) g.Node {
						return FormEl(
							ID("checked-contacts"),
							hx.Post("/examples/gomponents/bulk-update/"),
							hx.SwapExtended(
								swap.New().Strategy(swap.OuterHTML).Settle(3*time.Second),
							),
							hx.Target(
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
										g.Map(users, func(u form.UserModel) g.Node {
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
							UpdateToast(""),
						)
					}
				`)),
			),
		),
		P(
			g.Text("The server will bulk-update the statuses based on the values of the checkboxes. We respond with a small toast message about the update to inform the user, and use ARIA to politely announce the update for accessibility."),
		),
		Pre(
			Code(
				Class("language-css"),
				g.Text(dedent.Dedent(`
					#toast.htmx-settling {
						opacity: 100;
					}

					#toast {
						background: #E1F0DA;
						opacity: 0;
						transition: opacity 3s ease-out;
					}
				`)),
			),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(dedent.Dedent(`
					func UpdateToast(toast string) g.Node {
						return Span(
							ID("toast"),
							g.If(toast != "",
								g.Attr("aria-live", "polite"),
							),
							g.Text(toast),
						)
					}
				`)),
			),
		),
		P(
			g.Text("The cool thing is that, because HTML form inputs already manage their own state, we don’t need to re-render any part of the users table. The active users are already checked and the inactive ones unchecked!"),
		),
		P(
			g.Text("You can see a working example of this code below."),
		),
		P(
			A(
				Href("https://github.com/will-wow/typed-htmx-go/tree/main/examples/templ/web/bulkupdate"),
				Target("_blank"),
				Rel("noopener"),
				g.Text("Source"),
			),
		),
		H2(g.Text("Demo")),
		table(users),
	)
}

func table(users []form.UserModel) g.Node {
	return FormEl(
		ID("checked-contacts"),
		hx.Post("/examples/gomponents/bulk-update/"),
		hx.SwapExtended(
			swap.New().Strategy(swap.OuterHTML).Settle(3*time.Second),
		),
		hx.Target(
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
					g.Map(users, func(u form.UserModel) g.Node {
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
		UpdateToast(""),
	)
}

func UpdateToast(toast string) g.Node {
	return Span(
		ID("toast"),
		g.If(toast != "",
			g.Attr("aria-live", "polite"),
		),
		g.Text(toast),
	)
}
