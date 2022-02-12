package search

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type search struct {
	input  io.Reader
	output io.Writer
}

type option func(*search) error

func WithInput(input io.Reader) option {
	return func(s *search) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		s.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(s *search) error {
		if output == nil {
			return errors.New("nil output reader")
		}
		s.output = output
		return nil
	}
}

func NewSearch(opts ...option) (search, error) {
	s := search{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(&s)
		if err != nil {
			return search{}, err
		}
	}
	return s, nil
}

func (s search) Search(w string) {
	scanner := bufio.NewScanner(s.input)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), w) {
			fmt.Fprint(s.output, scanner.Text())
		}
	}
}
