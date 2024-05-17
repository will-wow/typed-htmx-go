package util_test

import (
	"testing"

	"github.com/will-wow/typed-htmx-go/htmx/internal/util"
)

func TestBoolToString(t *testing.T) {
	tests := []struct {
		input bool
		want  string
	}{
		{input: true, want: "true"},
		{input: false, want: "false"},
	}
	for _, test := range tests {
		got := util.BoolToString(test.input)
		if got != test.want {
			t.Errorf("BoolToString(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestJoinStringLikes(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		want := "a,b,c"
		got := util.JoinStringLikes([]string{"a", "b", "c"}, ",")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("string likes", func(t *testing.T) {
		type TestString string
		want := "a,b,c"
		got := util.JoinStringLikes([]TestString{"a", "b", "c"}, ",")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestMakeRelativeSelector(t *testing.T) {
	type TestModifier string
	var next TestModifier = "next"
	type TestSelector string

	selector := util.MakeRelativeSelector[TestModifier, TestSelector]()
	got := selector(next, "div")
	var want TestSelector = "next div"

	if got != want {
		t.Errorf(`got "%v", want "%v"`, got, want)
	}
}
