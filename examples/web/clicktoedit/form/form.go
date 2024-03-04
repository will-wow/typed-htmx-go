package form

import (
	"strings"

	"github.com/will-wow/typed-htmx-go/examples/web/ui"
)

type Form struct {
	ui.Form
	FirstName string
	LastName  string
	Email     string
}

func (f *Form) Validate() (ok bool) {
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

func New() *Form {
	return &Form{
		FirstName: "",
		LastName:  "",
		Email:     "",
		Form:      ui.NewForm(),
	}
}
