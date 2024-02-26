package web

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/will-wow/typed-htmx-go/examples/templ/web/bulkupdate"
	"github.com/will-wow/typed-htmx-go/examples/templ/web/clicktoedit"
	"github.com/will-wow/typed-htmx-go/examples/templ/web/examples"
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

func (h *Handler) routes() http.Handler {
	mux := http.NewServeMux()

	// Catch-all
	mux.HandleFunc("/", notFound)

	// Set up a in-memory file server for the embedded static files.
	fileServer := http.FileServerFS(staticFiles)
	mux.Handle("GET /static/", fileServer)

	mux.HandleFunc("/{$}", examples.Handler)
	delegateExample(mux, "/examples/click-to-edit", clicktoedit.Handler())
	delegateExample(mux, "/examples/bulk-update", bulkupdate.Handler())

	return h.recoverPanic(h.logRequest(mux))
}

func delegateExample(mux *http.ServeMux, path string, handler http.Handler) {
	mux.Handle(path+"/", http.StripPrefix(path, handler))
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

func notFound(w http.ResponseWriter, r *http.Request) {
	component := notFoundPage()
	w.WriteHeader(http.StatusNotFound)
	_ = component.Render(r.Context(), w)
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
