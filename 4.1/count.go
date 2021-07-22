package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct{
	input io.Reader
	output io.Writer
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func (c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func (c *counter) error {
		if len(args) < 1 {
			return errors.New("no args supplied")
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = f
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func (c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		input: os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return counter{}, err
		}
	}
	return c, nil
}

func (c counter) Lines() {
	fmt.Fprintln(c.output, c.NumberOfLines())
}

func (c counter) NumberOfLines() int {
	lines := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		lines++
	}
	return lines
}

func Lines() {
	c, err := NewCounter()
	if err != nil {
		panic("internal error calling NewCounter")
	}
	c.Lines()
}