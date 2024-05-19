// package ajaxheader adds the X-Requested-With header to requests with the value “XMLHttpRequest”.
package ajaxheader

import "github.com/will-wow/typed-htmx-go/htmx"

// Extension adds the X-Requested-With header to requests with the value “XMLHttpRequest”.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@1.9.12/dist/ext/ajax-header.js"></script>
//
// # Usage
//
//	<body { hx.Ext(ajaxheader.Extension)... } >
//
// Extension: [ajax-header]
//
// [ajax-header]: https://htmx.org/extensions/ajax-header/
const Extension htmx.Extension = "ajax-header"
