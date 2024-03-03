package clicktoedit

import (
	"net/http"
	"strings"

	"github.com/will-wow/typed-htmx-go/examples/web/ui"
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

var tEx = newTemplExample()
var gEx = newGomExample()

type example struct {
	gom bool
}

func NewHandler(gom bool) http.Handler {
	mux := http.NewServeMux()

	ex := example{gom: gom}

	mux.HandleFunc("GET /{$}", ex.demo)
	mux.HandleFunc("GET /view", ex.view)
	mux.HandleFunc("GET /edit", ex.edit)
	mux.HandleFunc("POST /edit", ex.post)

	return mux
}

func (e example) demo(w http.ResponseWriter, r *http.Request) {
	form := newForm()

	if e.gom {
		_ = gEx.page(form).Render(w)
	} else {
		component := tEx.page(form)
		_ = component.Render(r.Context(), w)
	}
}

func (e example) view(w http.ResponseWriter, r *http.Request) {
	form := newForm()

	if e.gom {
		_ = gEx.viewForm(form).Render(w)
	} else {
		component := tEx.viewForm(form)
		_ = component.Render(r.Context(), w)
	}
}

func (e example) edit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form := newForm()

	if e.gom {
		_ = gEx.editForm(form).Render(w)
	} else {
		component := tEx.editForm(form)
		_ = component.Render(r.Context(), w)
	}
}

func (e example) post(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusUnprocessableEntity)

		if e.gom {
			_ = gEx.editForm(form).Render(w)
		} else {
			component := tEx.editForm(form)
			_ = component.Render(r.Context(), w)
		}
		return
	}

	if e.gom {
		_ = gEx.viewForm(form).Render(w)
	} else {
		component := tEx.viewForm(form)
		_ = component.Render(r.Context(), w)
	}
}
