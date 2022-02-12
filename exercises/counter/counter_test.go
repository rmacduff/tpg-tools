package counter_test

import (
	"bytes"
	"counter"
	"testing"
)

func TestCounterFirstReturnZero(t *testing.T) {
	t.Parallel()
	want := 0
	c := counter.NewCounter()
	got := c.Next()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCounterNextReturnOne(t *testing.T) {
	t.Parallel()
	want := 1
	c := counter.NewCounter()
	c.Next()
	got := c.Next()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCounterNextReturnTwo(t *testing.T) {
	t.Parallel()
	want := 2
	c := counter.NewCounter()
	c.Next()
	c.Next()
	got := c.Next()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCounterSetOne(t *testing.T) {
	t.Parallel()
	want := 2
	c := counter.NewCounter()
	c.Count = 2
	got := c.Next()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCounterRun(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	c := counter.Counter{
		Count:    0,
		Output:   fakeTerminal,
		RunLimit: true,
		RunIter:  4,
		Delay:    1,
	}
	c.Run()
	want := "0123"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestCounterRunTime(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	c := counter.Counter{
		Count:    0,
		Output:   fakeTerminal,
		RunLimit: true,
		RunIter:  4,
		Delay:    1,
	}
	c.Run()
	want := "0123"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
