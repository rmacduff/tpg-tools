package writer

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
)

type writer struct {
	outFile  string
	numZeros int
	buffSize int
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
		w.outFile = args[0]
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

func NewWriter(opts ...option) (writer, error) {
	w := writer{
		buffSize: 1024 * 1024,
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

// func appendToFile(path string, data []byte) error {
// 	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	_, err = f.Write(data)
// 	return err
// }

func (w writer) WriteZeros(outFile string, numZeros int) error {
	w.outFile = outFile
	w.numZeros = numZeros
	return w.WriteZerosWithConfig()
}

func (w writer) WriteZerosWithConfig() error {
	f, err := os.OpenFile(w.outFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	numChunks := int(math.Ceil(float64(w.numZeros) / float64(w.buffSize)))
	var numWritten int
	zeroArray := make([]byte, w.buffSize)

	for i := 0; i < numChunks; i++ {
		// adjust number of zeros to be written if we don't have a full buffer's worth
		if w.numZeros-numWritten < w.buffSize {
			zeroArray = make([]byte, w.numZeros-numWritten)
		}
		_, err := f.Write(zeroArray)
		if err != nil {
			return err
		}
		numWritten += w.buffSize
	}
	return os.Chmod(w.outFile, 0600)
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
