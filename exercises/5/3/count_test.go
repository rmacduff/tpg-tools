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
	want := "\t3"
	got := c.Count()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
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
	want := "\t3"
	got := c.Count()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
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

func TestWordCount(t *testing.T) {
	t.Parallel()
	args := []string{"-w", "testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.FromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := "\t6"
	got := c.Count()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestByteCount(t *testing.T) {
	t.Parallel()
	args := []string{"-b", "testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.FromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := "\t21"
	got := c.Count()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestLineAndWord(t *testing.T) {
	t.Parallel()
	args := []string{"-l", "-w", "testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.FromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := "\t3\t6"
	got := c.Count()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestLinwAndByte(t *testing.T) {
	t.Parallel()
	args := []string{"-l", "-b", "testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.FromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := "\t3\t21"
	got := c.Count()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestWordAndByte(t *testing.T) {
	t.Parallel()
	args := []string{"-b", "-w", "testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.FromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := "\t6\t21"
	got := c.Count()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
