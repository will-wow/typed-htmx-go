package htmx

import "github.com/a-h/templ"

func NewTempl() HX[templ.Attributes] {
	return NewHX(
		func(key Attribute, value any) templ.Attributes {
			return templ.Attributes{string(key): value}
		},
	)
}
