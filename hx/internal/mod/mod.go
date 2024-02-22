// package mod providers utilities for working with modifiers.
package mod

import (
	"slices"
	"strings"
)

// Join is a port of strings.Join, with support for a key:value map.
func Join[T ~string](main string, modifiers map[T]string) string {
	if len(modifiers) == 0 {
		return main
	}

	// Start with the length of all the separators
	n := len(modifiers) - 1
	for modifier, value := range modifiers {
		// Add the length of every modifier:value pair
		n += len(modifier) + 1 + len(value)
	}

	var b strings.Builder

	// First always render the swap strategy.
	_, _ = b.WriteString(main)
	_ = b.WriteByte(' ')

	// Sort the modifiers to ensure a consistent order
	mods := make([]T, len(modifiers))
	i := 0
	for modifier := range modifiers {
		mods[i] = modifier
		i++
	}
	slices.Sort(mods)

	// Then render each modifier:value pair
	for i, modifier := range mods {
		value := modifiers[modifier]

		// Write the modifier
		_, _ = b.WriteString(string(modifier))

		// Write the value if it exists
		if value != "" {
			_ = b.WriteByte(':')
			_, _ = b.WriteString(value)
		}

		if i < len(modifiers)-1 {
			_ = b.WriteByte(' ')
		}
	}

	return b.String()
}
