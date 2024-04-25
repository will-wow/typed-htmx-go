package htmx

import "fmt"

// NewStringAttrs returns a HX instance that returns stringified attributes for direct use in HTML.
func NewStringAttrs() HX[string] {
	return NewHX(func(k Attribute, v any) string {
		switch v := v.(type) {
		// For strings, print the key='value' pair.
		case string:
			return fmt.Sprintf(`%s='%v'`, k, v)
		// For booleans, print just the key if true.
		case bool:
			if v {
				return string(k)
			}
		}

		return ""
	})
}
