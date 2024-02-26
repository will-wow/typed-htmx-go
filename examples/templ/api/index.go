package api

import (
	"net/http"

	"github.com/will-wow/typed-htmx-go/examples/templ/web"
)

func Index(w http.ResponseWriter, r *http.Request) {
	handler := web.NewHttpHandler()
	handler.ServeHTTP(w, r)
}
