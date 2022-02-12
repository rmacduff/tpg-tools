package hackertext_test

import (
	"bytes"
	"hackertext"
	"testing"
)

func TestPrintStdin(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	inputBuf := bytes.NewBufferString("Hello, hackers!")
	ht, err := hackertext.NewText(
		hackertext.WithInput(inputBuf),
		hackertext.WithOutput(fakeTerminal),
		hackertext.WithDelay(1),
	)
	if err != nil {
		t.Fatal(err)
	}
	err = ht.Print()
	if err != nil {
		t.Fatal(err)
	}
	want := "Hello, hackers!"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestPrintFile(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	args := []string{"testdata/one_line.txt"}
	ht, err := hackertext.NewText(
		hackertext.WithInputFromArgs(args),
		hackertext.WithOutput(fakeTerminal),
		hackertext.WithDelay(1),
	)
	if err != nil {
		t.Fatal(err)
	}
	err = ht.Print()
	if err != nil {
		t.Fatal(err)
	}
	want := "this is a test\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
