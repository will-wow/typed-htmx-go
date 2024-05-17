package static

import (
	"embed"

	"github.com/will-wow/typed-htmx-go/examples/web/exprint"
)

//go:embed main.css
var fs embed.FS

var ExCSS = exprint.New(fs, "/*", "*/")
