package classtools_ex

import (
	"net/http"

	"github.com/will-wow/typed-htmx-go/examples/web/classtools_ex/exgom"
	"github.com/will-wow/typed-htmx-go/examples/web/classtools_ex/extempl"
)

type example struct {
	gom bool
}

func NewHandler(gom bool) http.Handler {
	mux := http.NewServeMux()

	ex := example{gom: gom}

	mux.HandleFunc("GET /{$}", ex.demo)
	mux.HandleFunc("GET /foo/{$}", ex.demo)

	return mux
}

func (ex *example) demo(w http.ResponseWriter, r *http.Request) {
	if ex.gom {
		_ = exgom.Page().Render(w)
	} else {
		_ = extempl.Page().Render(r.Context(), w)
	}
}
