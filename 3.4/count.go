package count

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type counter struct {
	input            io.Reader
	output           io.Writer
	matchString      string
	matchInsensitive bool
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
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

func WithMatchString(matchString string) option {
	return func(c *counter) error {
		c.matchString = matchString
		return nil
	}
}

func WithMatchInsensitive(matchInsensitive bool) option {
	return func(c *counter) error {
		c.matchInsensitive = matchInsensitive
		return nil
	}
}

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		input:            os.Stdin,
		output:           os.Stdout,
		matchString:      "",
		matchInsensitive: false,
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
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		if c.matchInsensitive {
			if strings.Contains(
				strings.ToLower(scanner.Text()),
				strings.ToLower(c.matchString)) {
				lines++
			}

		} else {
			if strings.Contains(scanner.Text(), c.matchString) {
				lines++
			}
		}
	}
	return lines
}

func Lines() int {
	c, err := NewCounter()
	if err != nil {
		panic("internal error")
	}
	return c.Lines()
}
