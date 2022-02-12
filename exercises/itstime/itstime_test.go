package itstime_test

import (
	"bytes"
	"itstime"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	c := itstime.Clock{
		Now:    time.Date(2000, 1, 1, 2, 10, 0, 0, time.UTC),
		Output: fakeTerminal,
	}
	want := "It's 10 past 2"
	c.CurrentTime()
	got := fakeTerminal.String()

	if want != got {
		t.Errorf("wanted time %q, got %q", want, got)
	}
}
