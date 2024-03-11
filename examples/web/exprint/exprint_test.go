package exprint_test

import (
	"embed"
	"strings"
	"testing"

	"github.com/will-wow/typed-htmx-go/examples/web/exprint"
)

//go:embed exprint_test.go
var thisFile embed.FS

func TestNew(t *testing.T) {
	ex := exprint.New(thisFile, "//", "")

	//ex:start:test
	printed, err := ex.Print("exprint_test.go", "test")
	if err != nil {
		t.Fatalf("failed to print: %v", err)
	}
	//ex:end:test

	expected := strings.TrimSpace(`
printed, err := ex.Print("exprint_test.go", "test")
if err != nil {
	t.Fatalf("failed to print: %v", err)
}
`)

	if printed != expected {
		t.Fatalf("expected %q, got %q", expected, printed)
	}
}
