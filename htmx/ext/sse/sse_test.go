package sse_test

import (
	"fmt"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/sse"
)

var hx = htmx.NewStringAttrs()

func ExampleExtension() {
	attr := hx.Ext(sse.Extension)
	fmt.Println(attr)
	// Output: hx-ext='sse'
}

func ExampleMessage() {
	attr := sse.Swap(hx, sse.Message)
	fmt.Println(attr)
	// Output: sse-swap='message'
}

func ExampleConnect() {
	attr := sse.Connect(hx, "/chatroom")
	fmt.Println(attr)
	// Output: sse-connect='/chatroom'
}

func ExampleSwap() {
	attr := sse.Swap(hx, "event")
	fmt.Println(attr)
	// Output: sse-swap='event'
}

func ExampleTrigger() {
	attr := sse.Trigger(hx, "event")
	fmt.Println(attr)
	// Output: hx-trigger='sse:event'
}
