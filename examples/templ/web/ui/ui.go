// package ui holds shared ui helpers
package ui

func HasError(s string) string {
	if s == "" {
		return ""
	}
	return "true"
}

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

type Form struct {
	errors map[string]string
}

func NewForm() Form {
	return Form{
		errors: make(map[string]string),
	}
}

func (f *Form) HasErrors() bool {
	return len(f.errors) > 0
}

func (f *Form) HasError(field string) bool {
	return f.errors[field] != ""
}

func (f *Form) SetError(field string, message string) {
	f.errors[field] = message
}

func (f *Form) GetError(field string) string {
	return f.errors[field]
}

func (f *Form) SetRequiredError(field string) {
	f.errors[field] = "This field is required"
}

func (f *Form) AriaInvalid(field string) string {
	if f.errors[field] != "" {
		return "true"
	}
	return ""
}
