package clicktoedit

import (
	"cmp"

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

func (e gomExample) page(form *form) g.Node {
	return e.layout.Base("Click to edit",
		H1(g.Text("Click To Edit")),
		P(g.Text("The click to edit pattern provides a way to offer inline editing of all or part of a record without a page refresh.")),
		Ul(
			Li(
				g.Text("This pattern starts with a UI that shows the details of a contact. The div has a button that will get the editing UI for the contact from /contact/1/edit"),
			),
		),
		P(
			A(
				Href("https://github.com/will-wow/typed-htmx-go/tree/main/examples/templ/web/clicktoedit"),
				Target("_blank"),
				Rel("noopener"),
				g.Text("Source"),
			),
		),
		H2(g.Text("Demo")),
		e.viewForm(form),
	)
}

func (e gomExample) viewForm(form *form) g.Node {
	return Div(
		ghx.Target(htmx.TargetThis),
		ghx.Swap(swap.OuterHTML),
		Dl(
			Dt(g.Text("First Name")),
			Dd(g.Text(cmp.Or(form.FirstName, "None"))),
			Dt(g.Text("Last Name")),
			Dd(g.Text(cmp.Or(form.LastName, "None"))),
			Dt(g.Text("Email")),
			Dd(g.Text(cmp.Or(form.Email, "None"))),
			Div(Role("group"),
				Button(
					ghx.Get("/examples/gomponents/click-to-edit/edit"),
					g.Text("Click to edit"),
				),
			),
		),
	)
}

func (e gomExample) editForm(form *form) g.Node {
	return FormEl(
		Method("POST"),
		Action("/examples/gomponents/click-to-edit/edit"),
		ghx.Post("/examples/gomponents/click-to-edit/edit"),
		ghx.Target(htmx.TargetThis),
		ghx.Swap(swap.OuterHTML),
		Label(
			g.Text("First Name"),
			Input(
				Type("text"),
				Name("firstName"),
				Value(form.FirstName),
				g.If(
					form.HasError("firstName"),
					g.Attr("aria-invalid", "true"),
				),
				g.If(
					form.HasError("firstName"),
					Small(g.Text(form.GetError("firstName"))),
				),
			),
		),
		Label(
			g.Text("Last Name"),
			Input(
				Type("text"),
				Name("lastName"),
				Value(form.LastName),
				g.If(
					form.HasError("lastName"),
					g.Attr("aria-invalid", "true"),
				),
				g.If(
					form.HasError("lastName"),
					Small(g.Text(form.GetError("lastName"))),
				),
			),
		),
		Label(
			g.Text("Email"),
			Input(
				Type("text"),
				Name("email"),
				Value(form.Email),
				g.If(
					form.HasError("email"),
					g.Attr("aria-invalid", "true"),
				),
				g.If(
					form.HasError("email"),
					Small(g.Text(form.GetError("email"))),
				),
			),
		),
		Div(
			Role("group"),
			Button(
				Type("submit"),
				g.Text("Submit"),
			),
			A(
				Href("/examples/gomponents/click-to-edit/"),
				Role("button"),
				ghx.Get("/examples/gomponents/click-to-edit/view/"),
				g.Text("Cancel"),
			),
		),
	)
}
