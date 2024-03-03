# typed-htmx-go/hx

Well-documented Go functions for building [HTMX](https://htmx.org) attributes.

[![Go Reference](https://pkg.go.dev/badge/github.com/will-wow/typed-htmx-go.svg)](https://pkg.go.dev/github.com/will-wow/typed-htmx-go/hx)
![Code Coverage](./assets/badge.svg)

[htmx](https://htmx.org) is a powerful and simple tool for building dynamic, server-rendered web applications. It also pairs particularly well with [Templ](https://templ.guide) (a JSX-like template language for Go) and [Gomponents](https://www.gomponents.com/) (a Go-native view library).

However, when using it I have to have the [docs](https://htmx.org/reference) open, to look up the specifics of each modifier. I wanted the simplicity of HTMX, the editor support of Go, and beautiful integration with Templ, without sacrificing performance. I built typed-htmx-go.

`hx.NewTempl()` provides an `hx` struct that exposes all documented [HTMX attributes](https://htmx.org/reference/) as Go functions, and returns either [templ.Attributes](https://templ.guide/syntax-and-usage/attributes) for to be spread into a Templ element, or an attribute `g.Node` for Gomponents. You can also support other templating libraries by simply passing a new attribute constructor to `HX{}`.

Each function and option includes a Godoc comment copied from the extensive HTMX docs, so you can access that documentation right from the comfort of your editor.

## Examples

Usage examples are in [examples](./examples) (hosted at [typed-htmx-go.vercel.app](https://typed-htmx-go.vercel.app/))

## HTMX Version

`typed-hx-go` strives to keep up with HTMX releases. It currently supports HTMX `v1.9.10`.

## Goals

The project has some specific goals that drive the API.

### Complete HTMX attribute support

Every documented HTMX attribute and modifier should have a corresponding Go function. If it's missing something please submit an issue or a PR! And in the meantime, you can always drop back to a raw HTML attribute.

### No stringly-typed options

Many HTMX attributes (like `hx-swap` and `hx-trigger`) support a complex syntax of methods, modifiers, and selectors in the attribute string (like `hx-trigger='click[isActive] consume from:(#parent > #child) queue:first target:#element'`).

That's necessary for a tool that embeds in standard HTML attributes, but it requires a lot of studying the docs to get exactly right.

`hx` strives to provide function signatures and typed options that ensure you're passing the right options to the right modifiers.

Sometimes that means that `hx` will provide multiple functions for a single attribute. For instance, `hx` provides three methods for `hx-target`, stop you from doing `hx-target='this #element'` (which is invalid), and instead guide you towards valid options like:

- `hx.Target("#element")` => `hx-target="#element'`
- `hx.TargetNonStandard(hx.TargetThis)` => `hx-target='this'`
- `hx.TargetRelative(hx.TargetSelectorNext, "#element")` => `hx-target='next #element'`

As a corollary to this goal, it should also be difficult to create an invalid attribute. So if modifier must be accompanied by a selector (like the `next` in `hx-target`), then it must be exposed through a two-argument function.

### Full documentation in-editor

The [HTMX References](https://htmx.org/reference/) are through and readable (otherwise this project wouldn't have been possible!) However, having those docs at your fingertips as you write, instead of in a separate tab, is even better.

`hx` strives to have a Go-adjusted copy of every line of documentation from the HTMX References, including examples, included in the godocs of functions and options.

Note: This work is on going. If you see something missing, please submit a PR!

### Transferable HTMX skills

As much as possible, it should be the case that if you know HTMX, you can use `hx`, and using `hx` should prepare you to use raw HTMX. That means that attributes functions should match their HTMX counterparts, names should match terms in the docs, and arguments should occur in the order they are printed in the HTML.

This also means that written `hx` attributes should look like HTMX attributes. So if in HTMX you would write:

```html
<form
  method="GET"
  action="/page"
  hx-get="/page"
  hx-target="body"
  hx-replace-url="true"
  hx-swap="scroll:#search-results:top swap:1s"
></form>
```

The `hx` equivalent should take the same names and values in the same order:

```go
// templ
<form
	method="GET"
	action="/page"
	{ hx.Get("/page")... }
	{ hx.Target("body")... }
	{ hx.ReplaceURL(true)... }
	{ hx.Swap(
		swap.New().
		ScrollElement("#search-results", swap.Top).
		Swap(time.Second),
	)... }
>
```

```go
// gomponents
Form(
	Method("GET"),
	Action("/page"),
	hx.Get("/page"),
	hx.Target("body"),
	hx.ReplaceURL(true),
	hx.Swap(
		swap.New().
		ScrollElement("#search-results", swap.Top).
		Swap(time.Second),
	),
)
```

### Fully tested

Every attribute function should have a test to make sure it's printing valid HTMX. These are also a good opportunity to try out the API and make sure it's ergonomic in practice.

## Install

```bash
go get github.com/will-wow/typed-htmx-go
```

## Usage

```go
import (
	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/swap"
)

var hx = htmx.NewTempl()

templ SearchInput(search string) {
	<form
		method="GET"
		action="/page"
		class="relative mb-2"
		{ hx.Get(currentPage)...}
		{ hx.Target("body")...}
		{ hx.ReplaceURL(true)...}
		{ hx.Swap(swap.New().ScrollElement("#search-results", swap.Top))...}
	>
		@ui.Input(
			htmx.TemplAttrs(
				hx.Get(currentPage),
				hx.Trigger("input changed delay:500ms"),
				templ.Attributes{
					"id":          "search",
					"placeholder": "Search (/)",
					"type":        "search",
					"name":        "search",
					"value":       search,
				},
			),
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

### Install tools

```bash
task tools
```

### Check that everything is ready to commit, and update the coverage badge

```bash
task ready
```

### Publish

```bash
VERSION="0.0.0" task publish
```
