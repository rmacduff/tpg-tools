package search_test

import (
	"bytes"
	"search"
	"testing"
)

func TestSearchMatch(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("This string has a match")
	fakeTerminal := &bytes.Buffer{}
	matchString := "match"
	s, err := search.NewSearch(
		search.WithInput(inputBuf),
		search.WithOutput(fakeTerminal),
	)
	if err != nil {
		t.Fatal(err)
	}
	s.Search(matchString)
	want := "This string has a match"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestSearchNoMatch(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("This string has no match")
	fakeTerminal := &bytes.Buffer{}
	matchString := "bad"
	s, err := search.NewSearch(
		search.WithInput(inputBuf),
		search.WithOutput(fakeTerminal),
	)
	if err != nil {
		t.Fatal(err)
	}
	s.Search(matchString)
	want := ""
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
