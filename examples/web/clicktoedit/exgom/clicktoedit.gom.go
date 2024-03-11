package exgom

import (
	"embed"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/swap"

	"github.com/will-wow/typed-htmx-go/examples/web/clicktoedit/form"
	"github.com/will-wow/typed-htmx-go/examples/web/exprint"
	"github.com/will-wow/typed-htmx-go/examples/web/layout/gom/layout"
)

var hx = htmx.NewGomponents()

//go:embed clicktoedit.gom.go
var fs embed.FS
var ex = exprint.New(fs, "//", "")

func Page(form *form.Form) g.Node {
	return layout.Wrapper("Click to edit",
		H1(g.Text("Click To Edit")),
		P(g.Text("The click to edit pattern provides a way to offer inline editing of all or part of a record without a page refresh.")),
		Ul(
			Li(
				g.Text("This pattern starts with a UI that shows the details of a contact. The div has a button that will get the editing UI for the contact from /contact/1/edit"),
			),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("clicktoedit.gom.go", "ViewForm")),
			),
		),
		Ul(
			Li(
				g.Text("This returns a form that can be used to edit the contact"),
			),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("clicktoedit.gom.go", "EditForm")),
			),
		),
		Ul(
			Li(
				g.Text("The form issues a POST back to /edit, following the usual REST-ful pattern."),
			),
			Li(
				g.Text("If there is an error, the form swaps the form with error messages in place of the edit form."),
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
		ViewForm(form),
	)
}

func ViewForm(form *form.Form) g.Node {
	//ex:start:ViewForm
	return Div(
		hx.Target(htmx.TargetThis),
		hx.Swap(swap.OuterHTML),
		Dl(
			Dt(g.Text("First Name")),
			Dd(g.Text(form.FirstName)),
			Dt(g.Text("Last Name")),
			Dd(g.Text(form.LastName)),
			Dt(g.Text("Email")),
			Dd(g.Text(form.Email)),
			Div(Role("group"),
				Button(
					hx.Get("/examples/gomponents/click-to-edit/edit/"),
					g.Text("Click to edit"),
				),
			),
		),
	)
	//ex:end:ViewForm
}

func EditForm(form *form.Form) g.Node {
	//ex:start:EditForm
	return FormEl(
		Method("POST"),
		Action("/examples/gomponents/click-to-edit/edit/"),
		hx.Post("/examples/gomponents/click-to-edit/edit/"),
		hx.Target(htmx.TargetThis),
		hx.Swap(swap.OuterHTML),
		Label(
			g.Text("First Name"),
			Input(
				Type("text"),
				Name("firstName"),
				Value(form.FirstName),
				g.If(
					form.HasError("FirstName"),
					g.Attr("aria-invalid", "true"),
				),
			),
			g.If(
				form.HasError("FirstName"),
				Small(g.Text(form.GetError("FirstName"))),
			),
		),
		Label(
			g.Text("Last Name"),
			Input(
				Type("text"),
				Name("lastName"),
				Value(form.LastName),
				g.If(
					form.HasError("LastName"),
					g.Attr("aria-invalid", "true"),
				),
			),
			g.If(
				form.HasError("LastName"),
				Small(g.Text(form.GetError("LastName"))),
			),
		),
		Label(
			g.Text("Email"),
			Input(
				Type("text"),
				Name("email"),
				Value(form.Email),
				g.If(
					form.HasError("Email"),
					g.Attr("aria-invalid", "true"),
				),
			),
			g.If(
				form.HasError("Email"),
				Small(g.Text(form.GetError("Email"))),
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
				hx.Get("/examples/gomponents/click-to-edit/view/"),
				g.Text("Cancel"),
			),
		),
	)
	//ex:end:EditForm
}
