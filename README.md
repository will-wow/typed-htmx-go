# typed-htmx-go/hx

Well-documented Go functions for building [HTMX](https://htmx.org) attributes.

[![Go Reference](https://pkg.go.dev/badge/github.com/will-wow/typed-htmx-go.svg)](https://pkg.go.dev/github.com/will-wow/typed-htmx-go/hx)
![Code Coverage](./assets/badge.svg)

[htmx](https://htmx.org) is a powerful and simple tool for building dynamic, server-rendered web applications. It also pairs particularly well with [Templ](https://templ.guide), a JSX-like template builder for Go.

However, when using it I have to have the [docs](https://htmx.org/reference) open, to look up the specifics of each modifier. I wanted the simplicity of HTMX, the editor support of Go, and beautiful integration with Templ, without sacrificing performance. I built typed-htmx-go.

`hx.New()` provides a builder struct that wraps all documented [HTMX attributes](https://htmx.org/reference/) with Go functions, and `.Build()` returns a map that conforms to [templ.Attributes](https://templ.guide/syntax-and-usage/attributes). This allows the result to be spread into a Templ element or be passed to a Templ component. However this library has no actual dependency of Templ, and can be used by anything that can render a `map[string]any` to HTML attributes. You can also use `.String()` to get a formatted string of HTML attributes to directly render in a template.

Each function and option includes a Godoc comment copied from the extensive HTMX docs, so you can access that documentation right from the comfort of your editor.

## Goals

The project has some specific goals that drive the API.

### Complete HTMX attribute support

Every documented HTMX attribute and modifier should have a corresponding Go function. If it's missing something please submit an issue or a PR! And in the meantime, you can always drop back to a raw HTML attribute.

### No stringly-typed options

Many HTMX attributes (like `hx-swap` and `hx-trigger`) support a complex syntax of methods, modifiers, and selectors in the attribute string (like `hx-trigger='click[isActive] consume from:(#parent > #child) queue:first target:#element'`).

That's necessary for a tool that embeds in standard HTML attributes, but it requires a lot of studying the docs to get exactly right.

`hx` strives to provide function signatures and typed options that ensure you're passing the right options to the right modifiers.

Sometimes that means that `hx` will provide multiple functions for a single attribute. For instance, `hx` provides three methods for `hx-target`, stop you from doing `hx-target='this #element'` (which is invalid), and instead guide you towards valid options like:

- `.Target("#element")` => `hx-target="#element'` 
- `.TargetNonStandard(hx.TargetThis)` => `hx-target='this'`
- `.TargetRelative(hx.TargetSelectorNext, "#element")` => `hx-target='next #element'`

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
<form
	method="GET"
	action="/page"
	{ hx.New().
	Get("/page").
	Target("body").
	ReplaceURL(true).
	Swap(
		swap.New().
		ScrollElement("#search-results", swap.Top).
		Swap(time.Second),
	).
	Build()... }
>
```

### Templ compatibility

While this library isn't tied to Templ directly, it should always return attribute maps that work as [Templ attributes](https://templ.guide/syntax-and-usage/attributes) for spreading, and generally work nicely within Templ components.

However, it should also be possible to directly print attributes for use in a standard Go [html/template](https://pkg.go.dev/html/template#HTMLAttrhttps://pkg.go.dev/html/template) (with [HTMLAttr](https://pkg.go.dev/html/template#HTMLAttrhttps://pkg.go.dev/html/template#HTMLAttr)).

TODO: Figure out a safer method of including in an html/template.

### Fully tested

Every attribute function should have a test to make sure it's printing valid HTMX. These are also a good opportunity to try out the API and make sure it's ergonomic in practice.

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
		action="/page"
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
