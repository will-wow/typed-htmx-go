package clicktoedit

import (
	"net/http"

	"github.com/will-wow/typed-htmx-go/examples/web/clicktoedit/exgom"
	"github.com/will-wow/typed-htmx-go/examples/web/clicktoedit/extempl"
	"github.com/will-wow/typed-htmx-go/examples/web/clicktoedit/form"
	"github.com/will-wow/typed-htmx-go/examples/web/ui"
)

type example struct {
	gom bool
}

func NewHandler(gom bool) http.Handler {
	mux := http.NewServeMux()

	ex := example{gom: gom}

	mux.HandleFunc("GET /{$}", ex.demo)
	mux.HandleFunc("GET /view/", ex.view)
	mux.HandleFunc("GET /edit/", ex.edit)
	mux.HandleFunc("POST /edit/", ex.post)

	return mux
}

func (e example) demo(w http.ResponseWriter, r *http.Request) {
	form := form.New()

	if e.gom {
		_ = exgom.Page(form).Render(w)
	} else {
		component := extempl.Page(form)
		_ = component.Render(r.Context(), w)
	}
}

func (e example) view(w http.ResponseWriter, r *http.Request) {
	form := form.New()

	if e.gom {
		_ = exgom.ViewForm(form).Render(w)
	} else {
		component := extempl.ViewForm(form)
		_ = component.Render(r.Context(), w)
	}
}

func (e example) edit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form := form.New()

	if e.gom {
		_ = exgom.EditForm(form).Render(w)
	} else {
		component := extempl.EditForm(form)
		_ = component.Render(r.Context(), w)
	}
}

func (e example) post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form := &form.Form{
		FirstName: r.Form.Get("firstName"),
		LastName:  r.Form.Get("lastName"),
		Email:     r.Form.Get("email"),
		Form:      ui.NewForm(),
	}
	ok := form.Validate()

	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)

		if e.gom {
			_ = exgom.EditForm(form).Render(w)
		} else {
			component := extempl.EditForm(form)
			_ = component.Render(r.Context(), w)
		}
		return
	}

	if e.gom {
		_ = exgom.ViewForm(form).Render(w)
	} else {
		component := extempl.ViewForm(form)
		_ = component.Render(r.Context(), w)
	}
}
