package sse_ex

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/will-wow/typed-htmx-go/examples/web/sse_ex/exgom"
	"github.com/will-wow/typed-htmx-go/examples/web/sse_ex/extempl"
	"github.com/will-wow/typed-htmx-go/examples/web/sse_ex/shared"
)

type example struct {
	gom bool
}

func NewHandler(gom bool) http.Handler {
	mux := http.NewServeMux()

	ex := example{gom: gom}

	mux.HandleFunc("GET /{$}", ex.demo)
	mux.HandleFunc("GET /countdown/{$}", ex.countdown)
	mux.HandleFunc("GET /countdown/feed/{$}", ex.feed)

	return mux
}

func (ex *example) demo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if ex.gom {
		_ = exgom.Page().Render(w)
	} else {
		_ = extempl.Page().Render(ctx, w)
	}
}

func (ex *example) countdown(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if ex.gom {
		_ = exgom.Countdown().Render(w)
	} else {
		_ = extempl.Countdown().Render(ctx, w)
	}
}

func (ex *example) feed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	flusher, ok := w.(http.Flusher)
	if !ok {
		slog.Error("flush not supported")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("flush not supported"))
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for i := range 5 {
		countMessage := strconv.Itoa(5 - i)
		ex.sendEvent(w, shared.CountdownEvent, func() {
			if ex.gom {
				_ = exgom.Message(countMessage).Render(w)
			} else {
				_ = extempl.Message(countMessage).Render(ctx, w)
			}
		})
		flusher.Flush()
		time.Sleep(time.Second)
	}

	ex.sendEvent(w, shared.CountdownEvent, func() {
		if ex.gom {
			_ = exgom.Message("Blastoff!").Render(w)
		} else {
			_ = extempl.Message("Blastoff!").Render(ctx, w)
		}
	})
	flusher.Flush()
	time.Sleep(2 * time.Second)

	ex.sendEvent(w, shared.ResetEvent, func() {
		if ex.gom {
			_ = exgom.Trigger().Render(w)
		} else {
			_ = extempl.Trigger().Render(ctx, w)
		}
	})
}

func (ex *example) sendEvent(w http.ResponseWriter, event string, render func()) {
	_, _ = fmt.Fprintf(w, "event: %s\n", event)
	_, _ = fmt.Fprint(w, "data: ")
	render()
	_, _ = fmt.Fprint(w, "\n\n")
}
