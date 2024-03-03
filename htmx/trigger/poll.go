package trigger

import (
	"fmt"
	"time"
)

type Poll struct {
	string string
}

// trigger is a no-op method to satisfy the Trigger interface.
func (p *Poll) trigger() {}

// String returns the final hx-trigger string.
func (p *Poll) String() string {
	return p.string
}

// NewPoll creates a new polling trigger.
func NewPoll(timing time.Duration) *Poll {
	return &Poll{
		string: fmt.Sprintf("every %s", timing.String()),
	}
}

// NewPoll creates a new polling trigger with a javascript expression as a filter. When the timer goes off, the trigger will only occur if the expression evaluates to true.
func NewFilteredPoll(timing time.Duration, filter string) *Poll {
	return &Poll{
		string: fmt.Sprintf("every %s [%s]", timing.String(), filter),
	}
}
