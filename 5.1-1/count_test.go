package count_test

import (
	"bytes"
	"count"
	"io"
	"testing"
)

func TestFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.FromArgs(args),
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

func TestFromArgsErrorsOnBogusFlag(t *testing.T) {
	t.Parallel()
	args := []string{"-bogus"}
	_, err := count.NewCounter(
		count.WithOutput(io.Discard),
		count.FromArgs(args),
	)
	if err == nil {
		t.Fatal("want error on bogus flag, got nil")
	}
}

func TestFromArgsEmpty(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
		count.FromArgs([]string{}),
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

func TestWords(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2 words\n3 this time")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	got := c.Words()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWordCount(t *testing.T) {
	t.Parallel()
	args := []string{"-w", "testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.FromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 5
	got := c.Words()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
