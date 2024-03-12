package activesearch

import (
	"net/http"
	"strings"

	"github.com/will-wow/typed-htmx-go/examples/web/activesearch/exgom"
	"github.com/will-wow/typed-htmx-go/examples/web/activesearch/extempl"
	"github.com/will-wow/typed-htmx-go/examples/web/activesearch/shared"
)

type example struct {
	gom bool
}

func NewHandler(gom bool) http.Handler {
	mux := http.NewServeMux()

	ex := example{gom: gom}

	mux.HandleFunc("GET /{$}", ex.demo)
	mux.HandleFunc("POST /search/", ex.search)

	return mux
}

func (ex *example) demo(w http.ResponseWriter, r *http.Request) {
	if ex.gom {
		_ = exgom.Page().Render(w)
	} else {
		component := extempl.Page()
		_ = component.Render(r.Context(), w)
	}
}

func (ex *example) search(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "could not parse form", http.StatusBadRequest)
		return
	}

	// get the search search from the request
	search := r.FormValue("search")

	var filtered []shared.User

	if search == "" {
		filtered = shared.Data
	} else {
		search = strings.ToLower(search)

		filtered = make([]shared.User, 0, len(shared.Data))
		for _, user := range shared.Data {
			if strings.Contains(strings.ToLower(user.FirstName), search) ||
				strings.Contains(strings.ToLower(user.LastName), search) ||
				strings.Contains(strings.ToLower(user.Email), search) ||
				strings.Contains(strings.ToLower(user.City), search) {
				filtered = append(filtered, user)
			}
		}
	}

	if ex.gom {
		_ = exgom.SearchResults(filtered).Render(w)
	} else {
		component := extempl.SearchResults(filtered)
		_ = component.Render(r.Context(), w)
	}
}
