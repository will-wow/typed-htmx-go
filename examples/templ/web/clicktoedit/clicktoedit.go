package clicktoedit

import (
	"net/http"
	"strings"

	"github.com/will-wow/typed-htmx-go/examples/templ/web/ui"
)

type form struct {
	ui.Form
	FirstName string
	LastName  string
	Email     string
}

func (f *form) validate() (ok bool) {
	if f.FirstName == "" {
		f.SetRequiredError("FirstName")
	}
	if f.LastName == "" {
		f.SetRequiredError("LastName")
	}
	if f.Email == "" {
		f.SetRequiredError("Email")
	} else if !strings.Contains(f.Email, "@") {
		f.SetError("Email", "Invalid email address")
	}

	return !f.HasErrors()
}

func newForm() *form {
	return &form{
		FirstName: "",
		LastName:  "",
		Email:     "",
		Form:      ui.NewForm(),
	}
}

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", demo)
	mux.HandleFunc("GET /view", view)
	mux.HandleFunc("GET /edit", edit)
	mux.HandleFunc("POST /edit", post)

	return mux
}

func demo(w http.ResponseWriter, r *http.Request) {
	form := newForm()

	component := page(form)

	w.WriteHeader(http.StatusOK)
	_ = component.Render(r.Context(), w)
}

func view(w http.ResponseWriter, r *http.Request) {
	form := newForm()

	component := viewForm(form)
	w.WriteHeader(http.StatusOK)
	_ = component.Render(r.Context(), w)
}

func edit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form := newForm()

	component := editForm(form)
	w.WriteHeader(http.StatusOK)
	_ = component.Render(r.Context(), w)
}

func post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form := &form{
		FirstName: r.Form.Get("firstName"),
		LastName:  r.Form.Get("lastName"),
		Email:     r.Form.Get("email"),
		Form:      ui.NewForm(),
	}
	ok := form.validate()

	if !ok {
		component := editForm(form)
		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = component.Render(r.Context(), w)
		return
	}

	component := viewForm(form)
	w.WriteHeader(http.StatusOK)
	_ = component.Render(r.Context(), w)
}
