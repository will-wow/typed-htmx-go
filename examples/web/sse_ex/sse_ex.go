package sse_ex

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/will-wow/typed-htmx-go/examples/web/sse_ex/chatroom"
	"github.com/will-wow/typed-htmx-go/examples/web/sse_ex/extempl"
)

type example struct {
	gom  bool
	room *chatroom.Chatroom
}

func NewHandler(gom bool) http.Handler {
	mux := http.NewServeMux()

	ctx := context.TODO()

	c := chatroom.New(ctx)

	ex := example{gom: gom, room: c}

	mux.HandleFunc("GET /{$}", ex.demo)
	mux.HandleFunc("GET /chatroom/{$}", ex.chatroom)
	mux.HandleFunc("GET /chatroom/feed/{$}", ex.feed)
	mux.HandleFunc("POST /chatroom/post/{$}", ex.newPost)

	return mux
}

func (ex *example) demo(w http.ResponseWriter, r *http.Request) {
	component := extempl.Page()
	_ = component.Render(r.Context(), w)
}

func (ex *example) chatroom(w http.ResponseWriter, r *http.Request) {
	component := extempl.Chatroom()
	_ = component.Render(r.Context(), w)
}

func (ex *example) feed(w http.ResponseWriter, r *http.Request) {
	rc := http.NewResponseController(w)

	err := rc.SetWriteDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		fmt.Println("failed to set write deadline", err)
		return
	}

	feedTimeout := time.NewTimer(25 * time.Second)

	ctx := r.Context()
	room := ex.room.Join()
	defer func() {
		// Stop and drain the timer
		feedTimeout.Stop()

		// Leave the room
		ex.room.Leave(room)
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		select {
		case <-feedTimeout.C:
			fmt.Println("end feed")
			component := extempl.Entry()

			_, _ = fmt.Fprintf(w, "event: %s\n", chatroom.EndEvent)
			_, _ = fmt.Fprint(w, "data: ")
			_ = component.Render(ctx, w)
			_, _ = fmt.Fprint(w, "\n\n")
			_ = rc.Flush()
			return
		case <-ctx.Done():
			fmt.Println("ctx done")
			// TODO: header
			return
		case msg := <-room:
			fmt.Println("msg", msg.Message)
			component := extempl.ChatMessage(msg.Message)

			_, _ = fmt.Fprintf(w, "event: %s\n", chatroom.ChatEvent)
			_, _ = fmt.Fprint(w, "data: ")
			_ = component.Render(ctx, w)
			_, _ = fmt.Fprint(w, "\n\n")
			_ = rc.Flush()
		}
	}
}

func (ex *example) newPost(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	if message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len([]rune(message)) > 140 {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}

	ex.room.Send(message)
	w.WriteHeader(http.StatusNoContent)
}
