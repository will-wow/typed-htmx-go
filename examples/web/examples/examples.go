package examples

import (
	"net/http"

	"github.com/will-wow/typed-htmx-go/examples/web/examples/exgom"
	"github.com/will-wow/typed-htmx-go/examples/web/examples/extempl"
)

type Routes struct {
	gom bool
}

func NewRoutes(gom bool) Routes {
	return Routes{gom: gom}
}

func (e Routes) NewIndexHandler(w http.ResponseWriter, r *http.Request) {
	if e.gom {
		exgom.Page().Render(w)
	} else {
		component := extempl.Page()
		_ = component.Render(r.Context(), w)
	}
}
