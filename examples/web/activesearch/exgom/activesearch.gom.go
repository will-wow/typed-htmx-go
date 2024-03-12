package exgom

import (
	"embed"
	"time"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/htmx/trigger"

	"github.com/will-wow/typed-htmx-go/examples/web/activesearch/shared"
	"github.com/will-wow/typed-htmx-go/examples/web/exprint"
	"github.com/will-wow/typed-htmx-go/examples/web/layout/gom/layout"

	"github.com/will-wow/typed-htmx-go/htmx"
)

var hx = htmx.NewGomponents()

//go:embed activesearch.gom.go
var fs embed.FS
var ex = exprint.New(fs, "//", "")

func Page() g.Node {
	return layout.Wrapper(
		"Active Search",
		Class("active-search"),
		H1(g.Text("Active Search")),
		P(
			g.Text("This example actively searches a contacts database as the user enters text."),
		),
		P(
			g.Text("We start with a search input and an empty table:"),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("activesearch.gom.go", "search")),
			),
		),
		P(
			g.Text("The input issues a "), Code(g.Text("POST")), g.Text(", to "), Code(g.Text("/search")), g.Text(", on the input event and sets the body of the table to be the resulting content. Note that the keyup event could be used as well, but would not fire if the user pasted text with their mouse (or any other non-keyboard method)."),
		),
		P(
			g.Text("We add the "), Code(g.Text("delay:500ms")), g.Text(", modifier to the trigger to delay sending the query until the user stops typing. Additionally, we add the "), Code(g.Text("changed")), g.Text(", modifier to the trigger to ensure we don’t send new queries when the user doesn’t change the value of the input (e.g. they hit an arrow key, or pasted the same value)."),
		),
		P(
			g.Text("Since we use a search type input we will get an x in the input field to clear the input. To make this trigger a new POST we have to specify another trigger. We specify another trigger by using a comma to separate them. The"), Code(g.Text("search")), g.Text("trigger will be run when the field is cleared but it also makes it possible to override the 500 ms input event delay by just pressing enter."),
		),
		P(
			g.Text("Finally, we show an indicator when the search is in flight with the "), Code(g.Text("hx-indicator")), g.Text(", attribute."),
		),
		H2(g.Text("Demo")),
		search(),
	)
}

func search() g.Node {
	return g.Group([]g.Node{
		//ex:start:search
		H3(
			g.Text("Search Contacts"),
			Span(Class("htmx-indicator"),
				Img(Src("/static/img/bars.svg")), g.Text("Searching..."),
			),
		),
		Input(
			Class("form-control"),
			Type("search"),
			Name("search"),
			Placeholder("Begin Typing To Search Users..."),
			hx.Post("/examples/templ/active-search/search/"),
			hx.TriggerExtended(
				trigger.On("input").Changed().Delay(time.Millisecond*500),
				trigger.On("search"),
			),
			hx.Target("#search-results"),
			hx.Indicator(".htmx-indicator"),
		),
		Table(Class("table"),
			THead(
				Tr(
					Th(g.Text("First Name")),
					Th(g.Text("Last Name")),
					Th(g.Text("Email")),
				),
			),
			TBody(ID("search-results")),
		),
		//ex:end:search
	})
}

func SearchResults(users []shared.User) g.Node {
	return g.Group(g.Map(users, func(user shared.User) g.Node {
		return Tr(
			Td(g.Text(user.FirstName)),
			Td(g.Text(user.LastName)),
			Td(g.Text(user.Email)),
		)
	}))
}
