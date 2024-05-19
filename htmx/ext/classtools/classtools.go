// package classtools allows you to specify CSS classes that will be swapped onto or off of the elements by using a classes or data-classes attribute. This functionality allows you to apply CSS Transitions to your HTML without resorting to javascript.
package classtools

import (
	"strings"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
)

// Extension allows you to specify CSS classes that will be swapped onto or off of the elements by using a classes or data-classes attribute. This functionality allows you to apply CSS Transitions to your HTML without resorting to javascript.
//
// # Install
//
//	<script src="https://unpkg.com/htmx.org@1.9.12/dist/ext/class-tools.js"></script>
//
// Extension: [class-tools]
//
// [class-tools]: https://htmx.org/extensions/class-tools/
const Extension htmx.Extension = "class-tools"

// An operation represents the type of class operation to perform after the specified delay.
type operation string

const (
	operationAdd    operation = "add"
	operationRemove operation = "remove"
	operationToggle operation = "toggle"
)

// A classOperation represents a single operation to be performed on a class, after a delay.
type classOperation struct {
	operation operation
	class     string
	delay     time.Duration
}

// A run represents a sequence of class operations to be performed on an element.
// construct classOperations using the [Add], [Remove], and [Toggle] functions.
type Run []classOperation

// Classes allows you to specify CSS classes that will be swapped onto or off of the elements by using a classes or data-classes attribute. This functionality allows you to apply CSS Transitions to your HTML without resorting to javascript.
// A classes attribute value consists one or more [Run]s of class operations. All class operations within a given Run will be applied sequentially, with the delay specified.
// Use [ClassesParallel] to specify multiple parallel runs of class operations.
// A class operation is an operation ([Add], [Remove], or [Toggle]), a CSS class name, and an optional time delay. If the delay is not specified, the default delay is 100ms.
//
// # Usage
//
//	<div { hx.Ext(classtools.Extension)... }>
//		<div { classtools.Classes(classtools.Add("foo", time.Millisecond*100))... } /> <!-- adds the class "foo" after 100ms -->
//		<div class="bar" { classtools.Classes(hx, classtools.Remove(hx, "foo", time.Second*2))... } /> <!-- removes the class "bar" after 1s -->
//		<div class="bar" { classtools.Classes(hx, classtools.Remove(hx, "bar", time.Second), classtools.Add("foo", time.Second))... } /> <!-- removes the class "bar" after 1s then adds the class "foo" 1s after that -->
//		<div { classtools.Classes(classtools.Toggle(hx, "bar", time.Second))... } /> <!-- toggles the class "foo" every 1s -->
//	</div>
//
// Extension: [class-tools]
//
// [class-tools]: https://htmx.org/extensions/class-tools/
func Classes[T any](hx htmx.HX[T], operations ...classOperation) T {
	return ClassesParallel(hx, []Run{operations})
}

// ClassesParallel allows you to specify multiple runs of CSS classes that will be swapped onto or off of the elements by using a classes or data-classes attribute. This functionality allows you to apply CSS Transitions to your HTML without resorting to javascript.
// A classes attribute value consists one or more [Run]s of class operations. All class operations within a given Run will be applied sequentially, with the delay specified.
// Use [Classes] to more concisely specify a single runs of class operations.
// A class operation is an operation ([Add], [Remove], or [Toggle]), a CSS class name, and a time delay (which can be 0).
//
// # Usage
//
//	<div { hx.Ext(classtools.Extension)... }>
//		<div class="bar" { classtools.ClassesParallel(hx, []classtools.Run{
//			{classtools.Remove("bar", time.Second)},
//			{classtools.Add("foo", time.Second)},
//	})... } /> <!-- removes the class "bar" and adds class "foo" after 1s  -->
//	</div>
//
// Extension: [class-tools]
//
// [class-tools]: https://htmx.org/extensions/class-tools/
func ClassesParallel[T any](hx htmx.HX[T], runs []Run) T {
	classes := strings.Builder{}

	for i, run := range runs {
		for j, op := range run {
			classes.WriteString(string(op.operation))
			classes.WriteRune(' ')
			classes.WriteString(op.class)
			classes.WriteRune(':')
			classes.WriteString(op.delay.String())
			if j < len(run)-1 {
				classes.WriteString(", ")
			}
		}
		if i < len(runs)-1 {
			classes.WriteString(" & ")
		}
	}

	return hx.Attr("classes", classes.String())
}

// Add will add a class to the element after the specified delay.
// Only the first delay value will be used.
func Add(className string, delay time.Duration) classOperation {
	return makeOperation(operationAdd, className, delay)
}

// Remove will remove a class from the element after the specified delay.
func Remove(className string, delay time.Duration) classOperation {
	return makeOperation(operationRemove, className, delay)
}

// Toggle will toggle a class on the element on and off, every time the delay elapses.
func Toggle(className string, delay time.Duration) classOperation {
	return makeOperation(operationToggle, className, delay)
}

func makeOperation(op operation, className string, delay time.Duration) classOperation {
	return classOperation{
		operation: op,
		class:     className,
		delay:     delay,
	}
}
