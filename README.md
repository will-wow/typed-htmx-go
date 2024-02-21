# typed-htmx-go

Well-documented Go functions for building [HTMX](https://htmx.org) attributes.

[![Go Reference](https://pkg.go.dev/badge/github.com/will-wow/typed-htmx-go.svg)](https://pkg.go.dev/github.com/will-wow/typed-htmx-go/hx)
![Code Coverage](./assets/badge.svg)

[htmx](https://htmx.org) is a powerful and simple tool for building dynamic, server-rendered web applications. It also pairs particularly well with [Templ](https://templ.guide), a JSX-like template builder for Go.

However, when using it I constantly have to have the [docs](https://htmx.org/reference) open, to look up the specifics of each modifier. I wanted the simplicity of HTMX, the editor support of Go, and beautiful integration with Templ, without sacrificing performance. I built typed-htmx-go.

typed-htmx-go provides a builder struct that wraps all documented [HTMX attributes](https://htmx.org/reference/) with Go functions, and `.Build()` returns a map that conforms to [templ.Attributes](https://templ.guide/syntax-and-usage/attributes). This allows the result to be spread into a Templ element or be passed to a Templ component. However this library has no actual dependency of Templ, and can be used by anything that can render a `map[string]any` to HTML attributes. You can also use `.String()` to get a formatted string of HTML attributes to directly render in a template.

Each function and option includes a Godoc comment copied from the extensive HTMX docs, so you can access that documentation right from the comfort of your editor.

## Install

```bash
go get github.com/will-wow/typed-htmx-go
```

## Usage

```go
import (
	"github.com/will-wow/typed-htmx-go/hx"
	"github.com/will-wow/typed-htmx-go/hx/swap"
)

templ SearchInput(search string) {
	<form
		method="GET"
		action={ templ.URL(currentPage) }
		class="relative mb-2"
		{ hx.New().
		Get(currentPage).
		Target("body").
		ReplaceURL(true).
		Swap(swap.New().ScrollElement("#search-results", swap.Top)).
		Build()... }
	>
		@ui.Input(
			hx.New().
				Get(currentPage).
				Trigger("input changed delay:500ms").
				More(map[string]any{
					"id":          "search",
					"placeholder": "Search (/)",
					"type":        "search",
					"name":        "search",
					"value":       search,
				}).
				Build(),
			"peer")
		<div class="absolute bottom-0 right-2 top-0 flex items-center leading-none peer-focus:hidden">
			@icon.Search()
		</div>
	</form>
}
```

## Contributing

### Install Tasklist

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```
