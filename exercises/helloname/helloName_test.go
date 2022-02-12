package helloName_test

import (
	"bytes"
	"helloName"
	"testing"
)

func TestPrintsHelloMessageToTerminal(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	helloName.PrintNameTo(fakeTerminal, "friend")
	want := "Hello, friend"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
