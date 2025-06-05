package hello_test

import (
	"bytes"
	"testing"

	"github.com/Zipeer2/hello"
)

func TestPrintTo_PrintsHelloMessageToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	hello.PrintTo(buf)
	want := "Hello, World\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
