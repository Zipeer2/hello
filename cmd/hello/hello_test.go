package hello_test

import (
	"bytes"
	"testing"

	"github.com/Zipeer2/hello"
)

func TestPrintToPrintsHelloMessageToWriter(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	hello.Print(&buf)
	got := buf.String()
	want := "Hello, world\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
