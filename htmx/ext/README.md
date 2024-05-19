# HTMX Extensions

htmx includes a set of extensions out of the box that address common developer needs. These extensions are tested against htmx in each distribution.

While you can always use any extension by adding standard HTML attributes, `typed-htmx-go` has typed support for some extensions.

These extensions each have their own package, and expose function that take a configured `hx` as a first parameter, and return a full attribute.

`hx` also includes an `hx.Ext()` method to register an extension on an element (ie: `{ hx.Ext(classtools.Extension)... }` instead of `hx-ext="class-tools"`).
