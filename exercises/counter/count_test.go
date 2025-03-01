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

func TestWords(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2 and 3\n4")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
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
func TestWithInputFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.WithInputFromArgs(args),
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

func TestWithInputFromArgsEmpty(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
		count.WithInputFromArgs([]string{}),
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

func TestWithInputMultipleFiles(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt", "testdata/another_file.txt"}
	c, err := count.NewCounter(
		count.WithInputFromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
