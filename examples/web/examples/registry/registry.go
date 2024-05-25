package registry

import (
	"net/http"

	"github.com/will-wow/typed-htmx-go/examples/web/activesearch"
	"github.com/will-wow/typed-htmx-go/examples/web/bulkupdate"
	"github.com/will-wow/typed-htmx-go/examples/web/classtools_ex"
	"github.com/will-wow/typed-htmx-go/examples/web/clicktoedit"
	"github.com/will-wow/typed-htmx-go/examples/web/keyboard"
	"github.com/will-wow/typed-htmx-go/examples/web/progressbar"
	"github.com/will-wow/typed-htmx-go/examples/web/sse_ex"
)

type Example struct {
	Title   string
	Desc    string
	Slug    string
	Handler func(bool) http.Handler
}

var Examples = []Example{
	{
		Title:   "Click to Edit",
		Desc:    "Demonstrates inline editing of a data object",
		Slug:    "click-to-edit",
		Handler: clicktoedit.NewHandler,
	},
	{
		Title:   "Bulk Update",
		Desc:    "Demonstrates bulk updating of multiple rows of data",
		Slug:    "bulk-update",
		Handler: bulkupdate.NewHandler,
	},
	{
		Title:   "Active Search",
		Desc:    "Demonstrates the active search box pattern",
		Slug:    "active-search",
		Handler: activesearch.NewHandler,
	},
	{
		Title:   "Progress Bar",
		Desc:    "Demonstrates a job-runner like progress bar",
		Slug:    "progress-bar",
		Handler: progressbar.NewHandler,
	},
	{
		Title:   "Class Tools",
		Desc:    "Demo of class-tools options",
		Slug:    "class-tools",
		Handler: classtools_ex.NewHandler,
	},
	{
		Title:   "Server-Sent Events",
		Desc:    "Streaming responses with the SSE HTMX Extension",
		Slug:    "sse",
		Handler: sse_ex.NewHandler,
	},
	{
		Title:   "Keyboard Shortcuts",
		Desc:    "Demonstrates how to create keyboard shortcuts for htmx enabled elements",
		Slug:    "keyboard",
		Handler: keyboard.NewHandler,
	},
}
