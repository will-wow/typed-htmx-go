package bulkupdate

import (
	"fmt"
	"net/http"
)

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", demo)
	mux.HandleFunc("POST /", post)

	return mux
}

type userModel struct {
	name   string
	email  string
	active bool
}

func demo(w http.ResponseWriter, r *http.Request) {
	users := defaultUsers()

	component := page(users)
	w.WriteHeader(http.StatusOK)
	_ = component.Render(r.Context(), w)
}

func post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users := defaultUsers()

	var additions int
	var removals int

	for _, user := range users {
		if r.Form.Has(user.email) && !user.active {
			additions++
		} else if !r.Form.Has(user.email) && user.active {
			removals++
		}
	}

	toast := fmt.Sprintf("Activated %d and deactivated %d users", additions, removals)

	component := updateToast(toast)
	w.WriteHeader(http.StatusOK)
	_ = component.Render(r.Context(), w)
}

func defaultUsers() []userModel {
	return []userModel{
		{name: "Joe Smith", email: "joe@smith.org", active: true},
		{name: "Angie MacDowell", email: "angie@macdowell.org", active: true},
		{name: "Fuqua Tarkenton", email: "fuqua@tarkenton.org", active: true},
		{name: "Kim Yee", email: "kim@yee.org", active: false},
	}
}
