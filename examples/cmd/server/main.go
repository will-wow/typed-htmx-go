package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/will-wow/typed-htmx-go/examples/web"
)

func main() {
	handler := web.NewHttpHandler()

	//nolint:exhaustruct
	server := &http.Server{
		Addr:         "localhost:8080",
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	fmt.Printf("Listening on %v\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		os.Exit(1)
	}
}
