package bulkupdate

import (
	"fmt"
	"net/http"
)

func NewHandler(gom bool) http.Handler {
	mux := http.NewServeMux()

	ex := example{gom: gom}

	mux.HandleFunc("GET /", ex.demo)
	mux.HandleFunc("POST /", ex.post)

	return mux
}

type userModel struct {
	Name   string
	Email  string
	Active bool
}

var tEx = newTemplExample()
var gEx = newGomExample()

type example struct {
	gom bool
}

func (e example) demo(w http.ResponseWriter, r *http.Request) {
	users := defaultUsers()

	if e.gom {
		_ = gEx.page(users).Render(w)
	} else {
		component := tEx.page(users)
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
		_ = gEx.updateToast(toast).Render(w)
	} else {
		component := tEx.updateToast(toast)
		_ = component.Render(r.Context(), w)
	}
}

func defaultUsers() []userModel {
	return []userModel{
		{Name: "Joe Smith", Email: "joe@smith.org", Active: true},
		{Name: "Angie MacDowell", Email: "angie@macdowell.org", Active: true},
		{Name: "Fuqua Tarkenton", Email: "fuqua@tarkenton.org", Active: true},
		{Name: "Kim Yee", Email: "kim@yee.org", Active: false},
	}
}
