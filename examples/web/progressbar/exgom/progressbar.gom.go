package exgom

import (
	"embed"
	"fmt"
	"strconv"
	"time"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/will-wow/typed-htmx-go/htmx/swap"
	"github.com/will-wow/typed-htmx-go/htmx/trigger"

	"github.com/will-wow/typed-htmx-go/examples/web/exprint"
	"github.com/will-wow/typed-htmx-go/examples/web/layout/gom/layout"
	"github.com/will-wow/typed-htmx-go/examples/web/progressbar/shared"
	"github.com/will-wow/typed-htmx-go/examples/web/static"

	"github.com/will-wow/typed-htmx-go/htmx"
)

var hx = htmx.NewGomponents()

//go:embed progressbar.gom.go
var fs embed.FS
var ex = exprint.New(fs, "//", "")

func Page() g.Node {
	return layout.Wrapper(
		"Progress Bar",
		Class("progress-bar-demo"),
		H1(g.Text("Progress Bar")),
		P(
			g.Text("This example shows how to implement a smoothly scrolling progress bar."),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("progressbar.gom.go", "demo")),
			),
		),
		P(g.Text("This div is then replaced with a new div containing status and a progress bar that reloads itself every 600ms:")),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("progressbar.gom.go", "running")),
			),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("progressbar.gom.go", "progress")),
			),
		),
		P(
			g.Text("This progress bar is updated every 600 milliseconds, with the"),
			Code(g.Text("width")),
			g.Text("style attribute and "),
			Code(g.Text("aria-valuenow")),
			g.Text("attribute set to current progress value. Because there is an id on the progress bar div, htmx will smoothly transition between requests by settling the style attribute into its new value. This, when coupled with CSS transitions, makes the visual transition continuous rather than jumpy."),
		),
		P(
			g.Text("Finally, when the process is complete, a server returns a"),
			Code(g.Text("HX-Trigger: done")),
			g.Text("header, which triggers an update of the UI to “Complete” state with a restart button added to the UI (we are using the"),
			Code(g.Text("class-tools")),
			g.Text("extension in this example to add fade-in effect on the button):"),
		),
		Pre(
			Code(
				Class("language-go"),
				g.Text(ex.PrintOrErr("progressbar.gom.go", "done")),
			),
		),
		P(
			g.Text("This example uses styling cribbed from the bootstrap progress bar:"),
		),
		Pre(
			Code(
				Class("language-css"),
				g.Text(static.ExCSS.PrintOrErr("main.css", "progress-bar-style")),
			),
		),
		H2(g.Text("Demo")),
		demo(),
	)
}

func demo() g.Node {
	//ex:start:demo
	return Div(
		hx.Target(htmx.TargetThis),
		hx.Swap(swap.OuterHTML),
		H3(g.Text("Start Progress")),
		Button(
			Class("btn"),
			hx.Post("/examples/templ/progress-bar/job/"),
			g.Text("Start Job"),
		),
	)
	//ex:end:demo
}

func JobRunning(jobID int64, progress int) g.Node {
	//ex:start:running
	return Div(
		hx.Trigger(shared.TriggerDone),
		hx.Get("/examples/templ/progress-bar/job/%d/", jobID),
		hx.Swap(swap.OuterHTML),
		hx.Target(htmx.TargetThis),
		H3(Role("status"), ID("pblabel"), TabIndex("-1"), AutoFocus(),
			g.Text(fmt.Sprintf("Job %d Running", jobID)),
		),
		ProgressFetcher(jobID, progress),
	)
	//ex:end:running
}

func Job(jobID int64, progress int) g.Node {
	//ex:start:done
	return Div(
		hx.Target(htmx.TargetThis),
		hx.Swap(swap.OuterHTML),
		H3(Role("status"), ID("pblabel"), TabIndex("-1"), AutoFocus(),
			g.Text(fmt.Sprintf("Job %d Complete", jobID)),
		),
		ProgressBar(progress),
		Button(
			ID("restart-btn"),
			Class("btn"),
			hx.Post("/examples/templ/progress-bar/job/"),
			g.Attr("classes", "add show:600ms"),
			g.Text("Restart Job"),
		),
	)
	//ex:end:done
}

//ex:start:progress
func ProgressFetcher(jobID int64, progress int) g.Node {
	return Div(
		hx.Get("/examples/templ/progress-bar/job/%d/progress/", jobID),
		hx.TriggerExtended(trigger.Every(time.Millisecond*600)),
		hx.Target(htmx.TargetThis),
		hx.Swap(swap.InnerHTML),

		ProgressBar(progress),
	)
}

func ProgressBar(progress int) g.Node {
	return Div(
		Class("progress"),
		Role("progressbar"),
		Aria("valuemin", "0"),
		Aria("valuemax", "100"),
		Aria("valuenow", strconv.Itoa(progress)),
		Aria("labelledby", "pblabel"),
		Div(
			ID("pb"),
			Class("progress-bar"),
			// { progressWidth(progress)... }
		),
	)
}

//ex:end:progress
