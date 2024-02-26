package examples

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	component := page()
	w.WriteHeader(http.StatusOK)

	_ = component.Render(r.Context(), w)
}
