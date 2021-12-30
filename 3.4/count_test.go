package count_test

import (
	"bytes"
	"count"
	"testing"
)

func TestLines(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLinesStringMatch(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("no\nmatch\nnada")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
		count.WithMatchString("match"),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 1
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLinesStringMatchInsensitive(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("no\nmatch\nMatch\nnada")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
		count.WithMatchString("Match"),
		count.WithMatchInsensitive(true),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 2
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
