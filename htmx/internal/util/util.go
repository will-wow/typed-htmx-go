package util

import (
	"fmt"
	"strings"
)

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// joinStringLikes joins a slice of string-like values into a single string.
func JoinStringLikes[T ~string](elems []T, sep string) string {
	var stringElems = make([]string, len(elems))
	for i, x := range elems {
		stringElems[i] = string(x)
	}
	return strings.Join(stringElems, sep)
}

// makeRelativeSelector creates a function that combines an allowed relative modifier with a CSS selector and returns a typed result.
func MakeRelativeSelector[Modifier ~string, Selector ~string]() func(Modifier, string) Selector {
	return func(modifier Modifier, selector string) Selector {
		return Selector(fmt.Sprintf("%s %s", modifier, selector))
	}
}
