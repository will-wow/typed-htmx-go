package trigger_test

import (
	"fmt"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx/trigger"
)

func ExampleOn() {
	trig := trigger.On("click")
	fmt.Println(trig.String())
	// Output: click
}

func ExampleEvent_When() {
	trig := trigger.On("click").When("checkGlobalState()")
	fmt.Println(trig.String())
	// Output: click[checkGlobalState()]
}

func ExampleEvent_Once() {
	trig := trigger.On("click").Once()
	fmt.Println(trig.String())
	// Output: click once
}

func ExampleEvent_Changed() {
	trig := trigger.On("click").Changed()
	fmt.Println(trig.String())
	// Output: click changed
}

func ExampleEvent_Delay() {
	trig := trigger.On("click").Delay(time.Second)
	fmt.Println(trig.String())
	// Output: click delay:1s
}

func ExampleEvent_Throttle() {
	trig := trigger.On("click").Throttle(500 * time.Millisecond)
	fmt.Println(trig.String())
	// Output: click throttle:500ms
}

func ExampleEvent_From() {
	trig := trigger.On("click").From("#element")
	fmt.Println(trig.String())
	// Output: click from:(#element)
}

func ExampleEvent_From_withSpaces() {
	trig := trigger.On("click").From("parent > child")
	fmt.Println(trig.String())
	// Output: click from:(parent > child)
}

func ExampleEvent_From_nonStandard() {
	trig := trigger.On("click").From(trigger.FromDocument)
	fmt.Println(trig.String())
	// Output: click from:(document)
}

func ExampleFromRelative() {
	trig := trigger.On("click").From(trigger.FromRelative(trigger.Next, "#alert"))
	fmt.Println(trig.String())
	// Output: click from:next (#alert)
}

func ExampleFromRelative_withWhitespace() {
	trig := trigger.On("click").From(trigger.FromRelative(trigger.Next, "#alert > button"))
	fmt.Println(trig.String())
	// Output: click from:next (#alert > button)
}

func ExampleEvent_Target() {
	trig := trigger.On("click").Target("#element")
	fmt.Println(trig.String())
	// Output: click target:#element
}

func ExampleEvent_Target_withSpaces() {
	trig := trigger.On("click").Target("parent > child")
	fmt.Println(trig.String())
	// Output: click target:(parent > child)
}

func ExampleEvent_Consume() {
	trig := trigger.On("click").Consume()
	fmt.Println(trig.String())
	// Output: click consume
}

func ExampleEvent_Queue() {
	trig := trigger.On("click").Queue(trigger.First)
	fmt.Println(trig.String())
	// Output: click queue:first
}

func ExampleEvent_Clear() {
	trig := trigger.On("click").Consume().Clear(trigger.Consume)
	fmt.Println(trig.String())
	// Output: click
}

func ExampleOn_ordering_multiple() {
	trig := trigger.On("click").When("isActive").Queue(trigger.First).Consume().Target("#element").From("#parent > #child")
	fmt.Println(trig.String())
	// Output: click[isActive] consume from:(#parent > #child) queue:first target:#element
}

func ExampleIntersect() {
	trig := trigger.Intersect()
	fmt.Println(trig.String())
	// Output: intersect
}

func ExampleIntersectEvent_Root() {
	trig := trigger.Intersect().Root("#element")
	fmt.Println(trig.String())
	// Output: intersect root:#element
}

func ExampleIntersectEvent_Root_withSpaces() {
	trig := trigger.Intersect().Root("#parent > #child")
	fmt.Println(trig.String())
	// Output: intersect root:(#parent > #child)
}

func ExampleIntersectEvent_Threshold() {
	trig := trigger.Intersect().Threshold(0.2)
	fmt.Println(trig.String())
	// Output: intersect threshold:0.2
}

func ExampleIntersect_supportsOtherOptions() {
	trig := trigger.Intersect().Root("#element").Delay(time.Second)
	fmt.Println(trig.String())
	// Output: intersect delay:1s root:#element
}

func ExampleEvery() {
	trig := trigger.Every(5 * time.Second)
	fmt.Println(trig.String())
	// Output: every 5s
}

func ExamplePoll_Filter() {
	trig := trigger.Every(5 * time.Second).Filter("isActive")
	fmt.Println(trig.String())
	// Output: every 5s [isActive]
}

func ExampleEvent_OverrideEvent() {
	trig := trigger.On("keyup").Delay(time.Second).OverrideEvent("input")
	fmt.Println(trig.String())
	// Output: input delay:1s
}
