// package responsetargets allows you to specify different target elements to be swapped when different HTTP response codes are received.
package responsetargets

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/will-wow/typed-htmx-go/htmx"
)

// Extension allows you to specify different target elements to be swapped when different HTTP response codes are received.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@1.9.12/dist/ext/response-targets.js"></script>
//
// Extension: [response-targets]
//
// [response-targets]: https://htmx.org/extensions/response-targets/
const Extension htmx.Extension = "response-targets"

// A Code is a complete or partial HTTP response code.
type Code interface {
	code() string
}

// A Status is a complete HTTP response code. You can wrap the http.Status* constants with [Status].
type Status int

var _ Code = Status(0)

func (s Status) code() string {
	return strconv.Itoa(int(s))
}

// an errorCode is the string "error", used to cover all 4xx and 5xx HTTP response codes.
type errorCode string

var _ Code = errorCode("")

// Error is a status code that covers all 4xx and 5xx HTTP response codes.
const Error errorCode = "error"

func (e errorCode) code() string {
	return string(e)
}

// A wildcard is a partial HTTP response code with a wildcard component.
type wildcard []int

var _ Code = (wildcard)(nil)

// Wildcard creates a wildcard code with the given digits.
// For example, Wildcard(4, 1) results in hx-target-41*, and matches all 41x HTTP response codes.
func Wildcard(digits ...int) wildcard {
	return digits
}

func (w wildcard) code() string {
	builder := strings.Builder{}
	for _, digit := range w {
		_, _ = builder.WriteString(strconv.Itoa(digit))
	}
	_ = builder.WriteByte('*')
	return builder.String()
}

type wildcardX []int

// WildcardX creates a wildcard code with the given digits, and uses an 'x' instead of a '*' in the generated attribute.
// For example, WildcardX(4, 1) results in hx-target-41x, and matches all 41x HTTP response codes.
func WildcardX(digits ...int) wildcardX {
	return digits
}

func (w wildcardX) code() string {
	builder := strings.Builder{}
	for _, digit := range w {
		_, _ = builder.WriteString(strconv.Itoa(digit))
	}
	_ = builder.WriteByte('x')
	return builder.String()
}

// Target specifies a target element to be swapped when specific HTTP response codes are received.
//
// Extension: [response-targets]
//
// [response-targets]: https://htmx.org/extensions/response-targets/
func Target[T any](hx htmx.HX[T], code Code, extendedSelector htmx.TargetSelector) T {
	attr := fmt.Sprintf("hx-target-%s", code.code())
	return hx.Attr(htmx.Attribute(attr), string(extendedSelector))
}
