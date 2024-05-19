package loadingstates_test

import (
	"fmt"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
	"github.com/will-wow/typed-htmx-go/htmx/ext/loadingstates"
)

var hx = htmx.NewStringAttrs()

func ExampleDataLoading() {
	attr := loadingstates.DataLoading(hx)
	fmt.Println(attr)
	// Output: data-loading
}

func ExampleDataLoadingStyle() {
	attr := loadingstates.DataLoadingStyle(hx, "flex")
	fmt.Println(attr)
	// Output: data-loading='flex'
}

func ExampleDataLoadingClass() {
	attr := loadingstates.DataLoadingClass(hx, "bg-gray-100 opacity-80")
	fmt.Println(attr)
	// Output: data-loading-class='bg-gray-100 opacity-80'
}

func ExampleDataLoadingClassRemove() {
	attr := loadingstates.DataLoadingClassRemove(hx, "bg-gray-100")
	fmt.Println(attr)
	// Output: data-loading-class-remove='bg-gray-100'
}

func ExampleDataLoadingDisable() {
	attr := loadingstates.DataLoadingDisable(hx)
	fmt.Println(attr)
	// Output: data-loading-disable
}

func ExampleDataLoadingAriaBusy() {
	attr := loadingstates.DataLoadingAriaBusy(hx)
	fmt.Println(attr)
	// Output: data-loading-aria-busy
}

func ExampleDataLoadingDelay() {
	attr := loadingstates.DataLoadingDelay(hx)
	fmt.Println(attr)
	// Output: data-loading-delay
}

func ExampleDataLoadingDelayBy() {
	attr := loadingstates.DataLoadingDelayBy(hx, time.Second)
	fmt.Println(attr)
	// Output: data-loading-delay='1000'
}

func ExampleDataLoadingTarget() {
	attr := loadingstates.DataLoadingTarget(hx, "#target")
	fmt.Println(attr)
	// Output: data-loading-target='#target'
}

func ExampleDataLoadingPath() {
	attr := loadingstates.DataLoadingPath(hx, "/save")
	fmt.Println(attr)
	// Output: data-loading-path='/save'
}

func ExampleDataLoadingStates() {
	attr := loadingstates.DataLoadingStates(hx)
	fmt.Println(attr)
	// Output: data-loading-states
}
