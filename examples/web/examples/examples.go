package examples

import (
	"net/http"
)

var tEx = newTemplExample()
var gEx = newGomExample()

type Routes struct {
	gom bool
}

func NewRoutes(gom bool) Routes {
	return Routes{gom: gom}
}

func (e Routes) NewIndexHandler(w http.ResponseWriter, r *http.Request) {
	if e.gom {
		gEx.page().Render(w)
	} else {
		component := tEx.page()
		_ = component.Render(r.Context(), w)
	}
}

func (e Routes) NewNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	component := tEx.notFoundPage()
	_ = component.Render(r.Context(), w)
}

var ServerErrorPage = tEx.serverErrorPage
