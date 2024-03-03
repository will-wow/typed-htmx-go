package htmx

import "github.com/a-h/templ"

// NewTempl returns a HX instance for use with Templ.
// Each attribute returns a templ.Attributes map, which can be spread into a templ element.
func NewTempl() HX[templ.Attributes] {
	return NewHX(
		func(key Attribute, value any) templ.Attributes {
			return templ.Attributes{string(key): value}
		},
	)
}

// TemplAttrs merges the given attributes into a single map.
// This is helpful for passing many attributes to a templ component.
func TemplAttrs(attrs ...templ.Attributes) templ.Attributes {
	out := templ.Attributes{}
	for _, a := range attrs {
		for k, v := range a {
			out[k] = v
		}
	}
	return out
}
