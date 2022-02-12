package hackertext

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type hackertext struct {
	output io.Writer
	input  io.Reader
	delay  int
}

type option func(*hackertext) error

func WithOutput(output io.Writer) option {
	return func(ht *hackertext) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		ht.output = output
		return nil
	}
}

func WithInput(input io.Reader) option {
	return func(ht *hackertext) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		ht.input = input
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(ht *hackertext) error {
		if len(args) < 1 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		ht.input = f
		return nil
	}
}

func WithDelay(delay int) option {
	return func(ht *hackertext) error {
		ht.delay = delay
		return nil
	}
}

func NewText(opts ...option) (hackertext, error) {
	ht := hackertext{
		output: os.Stdout,
		input:  os.Stdin,
		delay:  100,
	}

	for _, opt := range opts {
		err := opt(&ht)
		if err != nil {
			return hackertext{}, err
		}
	}

	return ht, nil
}

func (ht hackertext) Print() error {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, ht.input)
	if err != nil {
		return errors.New("could not read from input")
	}
	for _, c := range buf.String() {
		fmt.Fprintf(ht.output, "%c", c)
		time.Sleep(time.Duration(ht.delay) * time.Millisecond)
	}
	return nil
}

func Print() {
	ht, err := NewText(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		panic("internal error")
	}
	ht.Print()
}
