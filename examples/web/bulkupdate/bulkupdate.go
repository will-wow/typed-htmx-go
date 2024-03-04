package bulkupdate

import (
	"fmt"
	"net/http"

	"github.com/will-wow/typed-htmx-go/examples/web/bulkupdate/exgom"
	"github.com/will-wow/typed-htmx-go/examples/web/bulkupdate/extempl"
	"github.com/will-wow/typed-htmx-go/examples/web/bulkupdate/form"
)

func NewHandler(gom bool) http.Handler {
	mux := http.NewServeMux()

	ex := example{gom: gom}

	mux.HandleFunc("GET /", ex.demo)
	mux.HandleFunc("POST /", ex.post)

	return mux
}

type example struct {
	gom bool
}

func (e example) demo(w http.ResponseWriter, r *http.Request) {
	users := defaultUsers()

	if e.gom {
		_ = exgom.Page(users).Render(w)
	} else {
		component := extempl.Page(users)
		_ = component.Render(r.Context(), w)
	}
}

func (e example) post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users := defaultUsers()

	var additions int
	var removals int

	for _, user := range users {
		if r.Form.Has(user.Email) && !user.Active {
			additions++
		} else if !r.Form.Has(user.Email) && user.Active {
			removals++
		}
	}

	toast := fmt.Sprintf("Activated %d and deactivated %d users", additions, removals)

	if e.gom {
		_ = exgom.UpdateToast(toast).Render(w)
	} else {
		component := extempl.UpdateToast(toast)
		_ = component.Render(r.Context(), w)
	}
}

func defaultUsers() []form.UserModel {
	return []form.UserModel{
		{Name: "Joe Smith", Email: "joe@smith.org", Active: true},
		{Name: "Angie MacDowell", Email: "angie@macdowell.org", Active: true},
		{Name: "Fuqua Tarkenton", Email: "fuqua@tarkenton.org", Active: true},
		{Name: "Kim Yee", Email: "kim@yee.org", Active: false},
	}
}
