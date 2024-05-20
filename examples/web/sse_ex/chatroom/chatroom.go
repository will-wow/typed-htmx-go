package chatroom

import (
	"context"
	"fmt"
	"slices"
)

const EndEvent = "end"
const ChatEvent = "chat"

type Chat struct {
	Message string
}

type Chatroom struct {
	source         chan Chat
	listeners      []chan Chat
	addListener    chan chan Chat
	removeListener chan (<-chan Chat)
}

func New(ctx context.Context) *Chatroom {
	c := &Chatroom{
		source:         make(chan Chat),
		listeners:      []chan Chat{},
		addListener:    make(chan chan Chat),
		removeListener: make(chan (<-chan Chat)),
	}

	go c.run(ctx)

	return c
}

func (c *Chatroom) Join() <-chan Chat {
	newListener := make(chan Chat)
	c.addListener <- newListener
	return newListener
}

func (c *Chatroom) Leave(channel <-chan Chat) {
	fmt.Println("removing listener")
	c.removeListener <- channel
}

func (c *Chatroom) Send(message string) {
	c.source <- Chat{Message: message}
}

func (c *Chatroom) run(ctx context.Context) {
	defer func() {
		for _, listener := range c.listeners {
			if listener != nil {
				close(listener)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("closing chatroom")
			return
		case newListener := <-c.addListener:
			c.listeners = append(c.listeners, newListener)
		case toRemove := <-c.removeListener:
			c.listeners = slices.DeleteFunc(c.listeners, func(l chan Chat) bool {
				if l == toRemove {
					fmt.Println("removed listener")
					return true
				}
				return false
			})
		case chat, ok := <-c.source:
			if !ok {
				return
			}
			for _, listener := range c.listeners {
				listener <- chat
			}
		}
	}
}
