package remember_test

import (
	"bytes"
	"os"
	"remember"
	"testing"
)

func TestPrint(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	r, err := remember.Setup(
		remember.WithOutput(fakeTerminal),
		remember.WithDbFile("testdata/simple_list.txt"),
	)
	if err != nil {
		t.Fatal(err)
	}
	r.Print()
	want := "milk\ncall mom\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestCreateFile(t *testing.T) {
	t.Parallel()
	_, err := remember.Setup(
		remember.WithDbFile("testdata/create_file_test.txt"),
	)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat("testdata/create_file_test.txt"); os.IsNotExist(err) {
		t.Errorf("got error creating file: %v", err)
	}
	os.Remove("testdata/create_file_test.txt")
}

func TestWrite(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	r, err := remember.Setup(
		remember.WithOutput(fakeTerminal),
		remember.WithDbFile("testdata/write_list.txt"),
	)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add("pay bills")
	if err != nil {
		t.Fatal(err)
	}
	r.Print()
	os.Remove("testdata/write_list.txt")
	want := "pay bills\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestWriteTwice(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	r, err := remember.Setup(
		remember.WithOutput(fakeTerminal),
		remember.WithDbFile("testdata/write_twice_list.txt"),
	)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add("first item")
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add("second item")
	if err != nil {
		t.Fatal(err)
	}
	r.Print()
	os.Remove("testdata/write_twice_list.txt")
	want := "first item\nsecond item\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
