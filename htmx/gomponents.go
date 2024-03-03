package htmx

import (
	"io"
	"strings"
	"text/template"

	g "github.com/maragudk/gomponents"
)

func NewGomponents() HX[GomponentsAttrs] {
	return NewHX(
		func(key Attribute, value any) GomponentsAttrs {
			return GomponentsAttrs{key: key, value: value}
		},
	)
}

type GomponentsAttrs struct {
	key   Attribute
	value any
}

var _ g.Node = GomponentsAttrs{key: "", value: false}

func (a GomponentsAttrs) Render(w io.Writer) error {
	switch v := a.value.(type) {
	// For strings, print the key='value' pair.
	case string:
		_, err := w.Write([]byte(" " + string(a.key) + `="` + template.HTMLEscapeString(v) + `"`))
		return err
	// For booleans, print just the key if true.
	case bool:
		if v {
			_, err := w.Write([]byte(" " + string(a.key)))
			return err
		}
	}
	return nil
}

// Type satisfies nodeTypeDescriber.
func (n GomponentsAttrs) Type() g.NodeType {
	return g.AttributeType
}

// String satisfies fmt.Stringer.
func (n GomponentsAttrs) String() string {
	var b strings.Builder
	_ = n.Render(&b)
	return b.String()
}
