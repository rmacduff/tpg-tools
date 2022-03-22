package writer

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
)

type writer struct {
	fileName   string
	numZeros   int
	buffSize   int
	output     io.Writer
	retryCount int
	File       os.File
}

type option func(*writer) error

func FromArgs(args []string) option {
	return func(w *writer) error {
		fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		numZeros := fset.Int("size", 1024, "Size of zero file to create")
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		w.numZeros = *numZeros
		args = fset.Args()
		if len(args) != 1 {
			return errors.New("need to specify file to write to")
		}
		w.fileName = args[0]
		w.output, err = os.OpenFile(w.fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		return nil
	}
}

func WithBuffSize(buffSize int) option {
	return func(w *writer) error {
		if buffSize < 1 {
			return errors.New("buffSize cannot be less than 1")
		}
		w.buffSize = buffSize
		return nil
	}
}

func WithZeros(numZeros int) option {
	return func(w *writer) error {
		if numZeros < 0 {
			return errors.New("number of zeros cannot be less than 0")
		}
		w.numZeros = numZeros
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(w *writer) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		w.output = output
		return nil
	}
}

func WithRetryCount(retry int) option {
	return func(w *writer) error {
		w.retryCount = retry
		return nil
	}
}

func NewWriter(opts ...option) (writer, error) {
	w := writer{
		buffSize:   1024 * 1024,
		retryCount: 3,
	}
	for _, opt := range opts {
		err := opt(&w)
		if err != nil {
			return writer{}, err
		}
	}
	return w, nil
}

func WriteToFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}
	return os.Chmod(path, 0600)
}

func (w *writer) WriteZerosToFile(outFile string, numZeros int) error {
	var err error
	w.fileName = outFile
	w.output, err = os.OpenFile(outFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	w.numZeros = numZeros
	return w.WriteZerosWithConfig()
}

func (w *writer) WriteZeros(output io.Writer, numZeros int) error {
	w.output = output
	w.numZeros = numZeros
	return w.WriteZerosWithConfig()
}
func (w writer) WriteZerosWithConfig() error {
	// f, err := os.OpenFile(w.outFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()

	var failedWrites int
	numChunks := int(math.Ceil(float64(w.numZeros) / float64(w.buffSize)))
	var numWritten int
	zeroArray := make([]byte, w.buffSize)

	for i := 0; i < numChunks; i++ {
		// adjust number of zeros to be written if we don't have a full buffer's worth
		if w.numZeros-numWritten < w.buffSize {
			zeroArray = make([]byte, w.numZeros-numWritten)
		}
		// fmt.Fprintf(os.Stdout, "retryCount: %d\n", w.retryCount)
		for failedWrites < w.retryCount {
			_, err := w.output.Write(zeroArray)
			if err != nil {
				failedWrites++
				fmt.Fprintf(os.Stdout, "failedWrites: %d\n", failedWrites)
			} else {
				break
			}
			if failedWrites == w.retryCount {
				return err
			}
		}
		numWritten += w.buffSize
		failedWrites = 0
	}
	// Do we actually want to be setting the file mode?
	// Would only need to be done if we're working with os.File
	// return os.Chmod(w.fileName, 0600)
	return nil
}

func RunCLI() {
	w, err := NewWriter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	w.WriteZerosWithConfig()
}
