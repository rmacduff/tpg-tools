package count

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type counter struct {
	countLines bool
	countWords bool
	countBytes bool
	input      io.Reader
	output     io.Writer
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

func FromArgs(args []string) option {
	return func(c *counter) error {
		fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		countLines := fset.Bool("l", false, "Count lines")
		countWords := fset.Bool("w", false, "Count words")
		countBytes := fset.Bool("b", false, "Count bytes")
		fset.SetOutput(c.output)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		c.countLines = *countLines
		c.countWords = *countWords
		c.countBytes = *countBytes
		if !*countLines && !*countWords && !*countBytes {
			c.countLines = true
		}
		args = fset.Args()
		if len(args) < 1 {
			return nil
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
		input:  os.Stdin,
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

// Wrap io.Reader so bytes can be counted
// Shamelessly taken from https://benjamincongdon.me/blog/2018/04/10/Counting-Scanned-Bytes-in-Go/
type byteCountReader struct {
	reader    io.Reader
	bytesRead int
}

func (r *byteCountReader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	r.bytesRead += n
	return n, err
}

func (c counter) Count() string {
	var lineCount int
	var wordCount int
	reader := &byteCountReader{reader: c.input}
	lineScanner := bufio.NewScanner(reader)
	for lineScanner.Scan() {
		lineCount++
		if c.countWords {
			wordScanner := bufio.NewScanner(bytes.NewBufferString(lineScanner.Text()))
			wordScanner.Split(bufio.ScanWords)
			for wordScanner.Scan() {
				wordCount++
			}
		}
	}

	var output string

	if c.countLines {
		output += fmt.Sprintf("\t%d", lineCount)
	}
	if c.countWords {
		output += fmt.Sprintf("\t%d", wordCount)
	}
	if c.countBytes {
		output += fmt.Sprintf("\t%d", reader.bytesRead)
	}

	return output
}

func RunCLI() {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(c.Count())
}
