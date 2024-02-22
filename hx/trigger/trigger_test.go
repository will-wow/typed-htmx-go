package trigger_test

import (
	"testing"
	"time"

	"github.com/will-wow/typed-htmx-go/hx/trigger"
)

func TestNewEvent(t *testing.T) {
	tests := []struct {
		name    string
		trigger trigger.Trigger
		want    string
	}{
		{
			name:    "simple Event",
			trigger: trigger.NewEvent("click"),
			want:    "click",
		},
		{
			name:    "Filter",
			trigger: trigger.NewEvent("click").Filter("checkGlobalState()"),
			want:    "click[checkGlobalState()]",
		},
		{
			name:    "Once",
			trigger: trigger.NewEvent("click").Once(),
			want:    "click once",
		},
		{
			name:    "Changed",
			trigger: trigger.NewEvent("click").Changed(),
			want:    "click changed",
		},
		{
			name:    "Delay",
			trigger: trigger.NewEvent("click").Delay(time.Second),
			want:    "click delay:1s",
		},
		{
			name:    "Throttle",
			trigger: trigger.NewEvent("click").Throttle(500 * time.Millisecond),
			want:    "click throttle:500ms",
		},
		{
			name:    "From",
			trigger: trigger.NewEvent("click").From("#element"),
			want:    "click from:#element",
		},
		{
			name:    "From with spaces",
			trigger: trigger.NewEvent("click").From("parent > child"),
			want:    "click from:(parent > child)",
		},
		{
			name:    "FromNonStandard",
			trigger: trigger.NewEvent("click").FromNonStandard(trigger.FromDocument),
			want:    "click from:document",
		},
		{
			name:    "FromRelative",
			trigger: trigger.NewEvent("click").FromRelative(trigger.FromSelectorNext, "#alert"),
			want:    "click from:next #alert",
		},
		{
			name:    "FromRelative with whitespace",
			trigger: trigger.NewEvent("click").FromRelative(trigger.FromSelectorNext, "#alert > button"),
			want:    "click from:next (#alert > button)",
		},
		{
			name:    "Target",
			trigger: trigger.NewEvent("click").Target("#element"),
			want:    "click target:#element",
		},
		{
			name:    "Target with spaces",
			trigger: trigger.NewEvent("click").Target("parent > child"),
			want:    "click target:(parent > child)",
		},
		{
			name:    "Consume",
			trigger: trigger.NewEvent("click").Consume(),
			want:    "click consume",
		},
		{
			name:    "Queue",
			trigger: trigger.NewEvent("click").Queue(trigger.First),
			want:    "click queue:first",
		},
		{
			name:    "Clear",
			trigger: trigger.NewEvent("click").Consume().Clear(trigger.Consume),
			want:    "click",
		},
		{
			name:    "Ordering multiple",
			trigger: trigger.NewEvent("click").Filter("isActive").Queue(trigger.First).Consume().Target("#element").From("#parent > #child"),
			want:    "click[isActive] consume from:(#parent > #child) queue:first target:#element",
		},
		{
			name:    "Intersect",
			trigger: trigger.NewIntersectEvent(),
			want:    "intersect",
		},
		{
			name:    "Intersect.Root",
			trigger: trigger.NewIntersectEvent().Root("#element"),
			want:    "intersect root:#element",
		},
		{
			name:    "Intersect.Root with spaces",
			trigger: trigger.NewIntersectEvent().Root("#parent > #child"),
			want:    "intersect root:(#parent > #child)",
		},
		{
			name:    "Intersect.Threshold",
			trigger: trigger.NewIntersectEvent().Threshold(0.2),
			want:    "intersect threshold:0.2",
		},
		{
			name:    "Intersect supports other options",
			trigger: trigger.NewIntersectEvent().Root("#element").Delay(time.Second),
			want:    "intersect delay:1s root:#element",
		},
		{
			name:    "Poll",
			trigger: trigger.NewPoll(5 * time.Second),
			want:    "every 5s",
		},
		{
			name:    "FilteredPoll",
			trigger: trigger.NewFilteredPoll(5*time.Second, "isActive"),
			want:    "every 5s [isActive]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.trigger.String()
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
