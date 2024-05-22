package exgom

import (
	"embed"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/sse"
	"github.com/will-wow/typed-htmx-go/htmx/swap"

	"github.com/will-wow/typed-htmx-go/examples/web/exprint"
	"github.com/will-wow/typed-htmx-go/examples/web/layout/gom/layout"
	"github.com/will-wow/typed-htmx-go/examples/web/sse_ex/shared"
)

var hx = htmx.NewGomponents()

//go:embed sse.gom.go
var fs embed.FS
var ex = exprint.New(fs, "//", "")

func Page() g.Node {
	return layout.Wrapper(
		"Server-Sent Events",
		H1(g.Text("Server-Sent Events")),
		Class("sse"),
		P(
			g.Text("A demo countdown using Server-Sent Events, with htmx and typed-htmx-go."),
		),
		P(
			g.Text("When you click the button below, that fetches a new element that uses the sse extension to start stream a countdown."),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("sse.gom.go", "trigger")),
			),
		),
		P(
			g.Text("The new elements uses the sse.Connect attribute to connect to a server-side event streaming countdown."),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("sse.gom.go", "countdown")),
			),
		),
		P(
			g.Text("Each second, the server sends a countdown message that updates the innerHTML of the div."),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("sse.gom.go", "message")),
			),
		),
		P(
			g.Text("After the countdown is complete, the server will send a ResetEvent, that closes removes the sse connection by replacing the sse.Connect element with the initial button using"),
			Code(g.Text("<code>hx.Swap(swap.OuterHTML)</code>")),
			g.Text("."),
		),
		H2(g.Text("Demo")),
		Trigger(),
	)
}

func Trigger() g.Node {
	//ex:start:trigger
	return Button(
		hx.Get("/examples/gomponents/sse/countdown/"),
		hx.Target(htmx.TargetThis),
		hx.Swap(swap.OuterHTML),
		g.Text("Start Countdown"),
	)
	//ex:end:trigger
}

func Countdown() g.Node {
	//ex:start:countdown
	return Div(
		hx.Ext(sse.Extension),
		sse.Connect(hx, "/examples/gomponents/sse/countdown/feed/"),
		sse.Swap(hx, shared.ResetEvent),
		hx.Swap(swap.OuterHTML),
		Div(
			sse.Swap(hx, shared.CountdownEvent),
			hx.Swap(swap.InnerHTML),
		),
	)
	//ex:end:countdown
}

//ex:start:message
func Message(msg string) g.Node {
	return P(g.Text(msg))
}

func Blastoff() g.Node {
	return P(g.Text("Blastoff!"))
}

//ex:end:message
