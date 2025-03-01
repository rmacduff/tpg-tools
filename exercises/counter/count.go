package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input  []io.Reader
	output io.Writer
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = append(c.input, input)
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) < 1 {
			return nil
		}
		for _, file := range args {
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			c.input = append(c.input, f)
		}
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		input:  make([]io.Reader, 0), // why does this have to be 0??
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

func (c counter) Lines() int {
	lines := 0
	for _, input := range c.input {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			lines++
		}
	}
	return lines
}

func (c counter) Words() int {
	words := 0
	for _, input := range c.input {
		scanner := bufio.NewScanner(input)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}
	return words
}

func Lines() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.Lines()
}

func Words() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.Words()

}
