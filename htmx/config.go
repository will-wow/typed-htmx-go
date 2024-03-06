package htmx

import (
	"encoding/json"

	"github.com/will-wow/typed-htmx-go/htmx/hxconfig"
)

// Config sets the htmx configuration options on a meta element.
//
//	<meta
//		name="htmx-config"
//		hx.Config(
//			hxconfig.New().DefaultSwapStyle(swap.OuterHTML),
//		)
//	/>
//
// [HTMX Docs]
//
// [HTMX Docs]: https://htmx.org/reference/#config
func (hx *HX[T]) Config(config *hxconfig.Builder) T {
	c := config.Build()
	bytes, err := json.Marshal(c)
	if err != nil {
		var empty T
		return empty
	}
	return hx.attr("content", string(bytes))
}
