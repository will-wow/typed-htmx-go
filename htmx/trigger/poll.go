package trigger

import (
	"fmt"
	"time"
)

type Poll struct {
	timing time.Duration
	filter string
}

// Every creates a new polling trigger.
func Every(timing time.Duration) *Poll {
	return &Poll{
		timing: timing,
		filter: "",
	}
}

// trigger is a no-op method to satisfy the Trigger interface.
func (p *Poll) trigger() {}

// String returns the final hx-trigger string.
func (p *Poll) String() string {
	if p.filter != "" {
		return fmt.Sprintf("every %s [%s]", p.timing.String(), p.filter)
	}

	return fmt.Sprintf("every %s", p.timing.String())
}

// Filter adds a filter to the polling trigger, so that when the timer goes off, the trigger will only occur if the expression evaluates to true.
func (p *Poll) Filter(filter string) *Poll {
	p.filter = filter
	return p
}
