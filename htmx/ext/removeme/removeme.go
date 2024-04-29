package removeme

import (
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
)

// Extension allows you to remove an element after a specified interval.
//
// extension: [remove-me]
//
// [remove-me]: https://htmx.org/extensions/remove-me/
const Extension htmx.Extension = "remove-me"

// RemoveMe removes the element after the specified interval.
func RemoveMe[T any](hx htmx.HX[T], after time.Duration) T {
	return hx.Attr("remove-me", after.String())
}
