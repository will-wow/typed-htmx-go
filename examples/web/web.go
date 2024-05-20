package web

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/will-wow/typed-htmx-go/examples/web/activesearch"
	"github.com/will-wow/typed-htmx-go/examples/web/bulkupdate"
	"github.com/will-wow/typed-htmx-go/examples/web/classtools_ex"
	"github.com/will-wow/typed-htmx-go/examples/web/clicktoedit"
	"github.com/will-wow/typed-htmx-go/examples/web/examples"
	"github.com/will-wow/typed-htmx-go/examples/web/progressbar"
	"github.com/will-wow/typed-htmx-go/examples/web/sse_ex"
)

//go:embed "static"
var staticFiles embed.FS

type Handler struct {
	logger *slog.Logger
}

func NewHttpHandler() http.Handler {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	handler := &Handler{
		logger: logger,
	}

	return handler.routes()
}

var templIndexRoutes = examples.NewRoutes(false)
var gomIndexRoutes = examples.NewRoutes(true)

func (h *Handler) routes() http.Handler {
	mux := http.NewServeMux()

	// Catch-all
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)

		component := notFoundPage()
		_ = component.Render(r.Context(), w)
	})

	// Set up a in-memory file server for the embedded static files.
	fileServer := http.FileServerFS(staticFiles)
	mux.Handle("GET /static/", fileServer)

	mux.HandleFunc("/{$}", templIndexRoutes.NewIndexHandler)
	mux.HandleFunc("/examples/gomponents/{$}", gomIndexRoutes.NewIndexHandler)
	delegateExample(mux, "click-to-edit", clicktoedit.NewHandler)
	delegateExample(mux, "bulk-update", bulkupdate.NewHandler)
	delegateExample(mux, "active-search", activesearch.NewHandler)
	delegateExample(mux, "progress-bar", progressbar.NewHandler)
	delegateExample(mux, "class-tools", classtools_ex.NewHandler)
	delegateExample(mux, "sse", sse_ex.NewHandler)

	return h.recoverPanic(h.logRequest(mux))
}

func delegateExample(mux *http.ServeMux, path string, handler func(bool) http.Handler) {
	prefix := fmt.Sprintf("/examples/templ/%s", path)
	mux.Handle(prefix+"/", http.StripPrefix(prefix, handler(false)))

	prefix = fmt.Sprintf("/examples/gomponents/%s", path)
	mux.Handle(prefix+"/", http.StripPrefix(prefix, handler(true)))
}

func (h *Handler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			proto  = r.Proto
			method = r.Method
			uri    = r.URL.RequestURI()
		)

		h.logger.Info("received request", "ip", ip, "proto", proto, "method", method, "uri", uri)

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always be run in the event
		// of a panic as Go unwinds the stack).
		defer func() {
			// Use the builtin recover function to check if there has been a
			// panic or not. If there has...
			if err := recover(); err != nil {
				// Set a "Connection: close" header on the response.
				w.Header().Set("Connection", "close")
				component := serverErrorPage(fmt.Sprintf("%s", err))
				w.WriteHeader(http.StatusInternalServerError)
				_ = component.Render(r.Context(), w)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
