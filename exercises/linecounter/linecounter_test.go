package linecounter_test

import (
	"bytes"
	"linecounter"
	"testing"
)

func TestCountLines(t *testing.T) {
	t.Parallel()
	// fakeTerminal := &bytes.Buffer{}
	// c := linecounter.Counter{
	// 	Output: os.Stdout,
	// }
	c := linecounter.NewCounter()
	c.Input = bytes.NewBufferString("1\n2\n3")
	got := c.Lines()
	want := 3
	if want != got {
		t.Errorf("for input %s: want line count %d, got %d", c.Input, want, got)
	}
}
