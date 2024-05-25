package exgom

import (
	"embed"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/htmx/trigger"

	"github.com/will-wow/typed-htmx-go/examples/web/exprint"
	"github.com/will-wow/typed-htmx-go/examples/web/layout/gom/layout"

	"github.com/will-wow/typed-htmx-go/htmx"
)

var hx = htmx.NewGomponents()

//go:embed keyboard.gom.go
var fs embed.FS
var ex = exprint.New(fs, "//", "")

func Page() g.Node {
	return layout.Wrapper(
		"Keyboard Shortcuts",
		H1(g.Text("Keyboard Shortcuts")),
		Class("keyboard-shortcuts"),
		P(
			g.Text("In this example we show how to create a keyboard shortcut for an action."),
		),
		P(
			g.Text("We start with a simple button that loads some content from the server:"),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("keyboard.gom", "demo")),
			),
		),
		P(
			g.Text("Note that the button responds to both the click event (as usual) and also the keyup event when alt-shift-D is pressed. The from: modifier is used to listen for the keyup event on the body element, thus making it a “global” keyboard shortcut."),
		),
		P(
			g.Text("You can trigger the demo below by either clicking on the button, or by hitting alt-shift-D."),
		),
		P(
			g.Text("You can find out the conditions needed for a given keyboard shortcut here:"),
		),
		P(
			A(
				Href("https://javascript.info/keyboard-events"),
				Target("_blank"),
				g.Text("https://javascript.info/keyboard-events"),
			),
		),
		H2(g.Text("Demo")),
		demo(),
	)
}

func demo() g.Node {
	//ex:start:demo
	return Button(
		hx.TriggerExtended(
			trigger.On("click"),
			trigger.
				On("keyup").
				From("body").
				When("altKey&&shiftKey&&key=='D'"),
		),
	)
	//ex:end:demo
}
