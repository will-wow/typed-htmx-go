package keyboard

import (
	"net/http"

	"github.com/will-wow/typed-htmx-go/examples/web/keyboard/extempl"
)

type example struct {
	gom bool
}

func NewHandler(gom bool) http.Handler {
	mux := http.NewServeMux()

	ex := example{gom: gom}

	mux.HandleFunc("GET /{$}", ex.demo)
	mux.HandleFunc("POST /doit/{$}", ex.doIt)

	return mux
}

func (ex *example) demo(w http.ResponseWriter, r *http.Request) {
	component := extempl.Page()
	_ = component.Render(r.Context(), w)
}

func (ex *example) doIt(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Dit it!"))
}
