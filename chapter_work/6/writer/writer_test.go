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
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
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

func TestPrePopPermsOverwrite(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/prepop_test.txt"
	err := os.WriteFile(path, []byte{1}, 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = writer.WriteToFile(path, []byte{1})
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
