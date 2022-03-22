package writer_test

import (
	"os"
	"testing"
	"writer"

	"github.com/google/go-cmp/cmp"
)

func TestWriteToFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/write_test.txt"
	want := []byte{1, 2, 3}
	err := writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode().Perm()
	if perm != 0600 {
		t.Errorf("want file mode 0600, got 0%o", perm)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestWriteToFileClobbers(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/clobber_test.txt"
	err := os.WriteFile(path, []byte{4, 5, 6}, 0600)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{1, 2, 3}
	err = writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestPermsClosed(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/perms_test.txt"
	// Pre-create empty file with open perms
	err := os.WriteFile(path, []byte{}, 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = writer.WriteToFile(path, []byte{})
	if err != nil {
		t.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode().Perm()
	if perm != 0600 {
		t.Errorf("want file mode 0600, got 0%o", perm)
	}
}

func TestWriteZeros(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/zeros_test.txt"
	want := make([]byte, 1000)
	w, err := writer.NewWriter(
	// 	writer.WithBuffSize(3),
	)
	if err != nil {
		t.Fatal(err)
	}
	err = w.WriteZeros(path, 1000)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestWriteZerosWithSmallBuffer(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/zeros_test.txt"
	want := make([]byte, 100)
	w, err := writer.NewWriter(
		writer.WithBuffSize(3),
	)
	if err != nil {
		t.Fatal(err)
	}
	err = w.WriteZeros(path, 100)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestWithArgsMissingFile(t *testing.T) {
	t.Parallel()
	args := []string{"-size", "1000"}
	_, err := writer.NewWriter(
		writer.FromArgs(args),
	)
	if err == nil {
		t.Fatal("want err on missing size option")
	}
}

func TestWithArgs(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/output_test.txt"
	args := []string{"-size", "1000", path}
	w, err := writer.NewWriter(
		writer.FromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	err = w.WriteZerosWithConfig()
	if err != nil {
		t.Fatal(err)
	}
	want := make([]byte, 1000)
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
