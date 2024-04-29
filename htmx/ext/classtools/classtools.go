package classtools

import (
	"strings"
	"time"

	"github.com/will-wow/typed-htmx-go/htmx"
)

const Extension htmx.Extension = "class-tools"

type operation string

const (
	operationAdd    operation = "add"
	operationRemove operation = "remove"
	operationToggle operation = "toggle"
)

type classOperation struct {
	Operation operation
	Class     string
	Delay     time.Duration
}

type Run []classOperation

func Classes[T any](hx htmx.HX[T], runs []Run) T {
	classes := strings.Builder{}

	for i, run := range runs {
		for j, class := range run {
			classes.WriteString(string(class.Operation))
			classes.WriteRune(' ')
			classes.WriteString(class.Class)
			if class.Delay > 0 {
				classes.WriteRune(':')
				classes.WriteString(class.Delay.String())
			}
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

func Add(className string, delay ...time.Duration) classOperation {
	return makeOperation(operationAdd, className, delay)
}

func Remove(className string, delay ...time.Duration) classOperation {
	return makeOperation(operationRemove, className, delay)
}

func Toggle(className string, delay ...time.Duration) classOperation {
	return makeOperation(operationToggle, className, delay)
}

func makeOperation(op operation, className string, delay []time.Duration) classOperation {
	var delayValue time.Duration
	if len(delay) > 0 {
		delayValue = delay[0]
	}

	return classOperation{
		Operation: op,
		Class:     className,
		Delay:     delayValue,
	}
}
