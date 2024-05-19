// package removeme allows you to remove an element after a specified interval.
package removeme

import (
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
)

// Extension allows you to remove an element after a specified interval.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@1.9.12/dist/ext/remove-me.js"></script>
//
// Extension: [remove-me]
//
// [remove-me]: https://htmx.org/extensions/remove-me/
const Extension htmx.Extension = "remove-me"

// RemoveMe removes the element after the specified interval.
//
// Extension: [remove-me]
//
// [remove-me]: https://htmx.org/extensions/remove-me/
func RemoveMe[T any](hx htmx.HX[T], after time.Duration) T {
	return hx.Attr("remove-me", after.String())
}
