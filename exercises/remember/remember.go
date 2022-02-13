package remember

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type remember struct {
	output io.Writer
	dbFile string
}

type option func(*remember) error

func WithOutput(output io.Writer) option {
	return func(r *remember) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		r.output = output
		return nil
	}
}

func WithDbFile(fileName string) option {
	return func(r *remember) error {
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			if _, err := os.Create(fileName); err != nil {
				return fmt.Errorf("could not create db file: %s", fileName)
			}
		}
		r.dbFile = fileName
		return nil
	}
}

func Setup(opts ...option) (remember, error) {
	usr, err := user.Current()
	if err != nil {
		return remember{}, err
	}
	file := filepath.Join(usr.HomeDir, ".remember.db")
	r := remember{
		output: os.Stdout,
		dbFile: file,
	}
	for _, opt := range opts {
		err := opt(&r)
		if err != nil {
			return remember{}, err
		}
	}

	return r, nil
}

func (r remember) Print() error {
	f, err := os.Open(r.dbFile)
	if err != nil {
		return fmt.Errorf("could not open file: %s", r.dbFile)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Fprintln(r.output, scanner.Text())
	}

	return nil
}

func (r remember) Add(item string) error {
	f, err := os.OpenFile(r.dbFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %q", err)
	}
	defer f.Close()

	if _, err := f.WriteString(item + "\n"); err != nil {
		return fmt.Errorf("could not write item to file: %q", err)
	}

	return nil
}

func Remember() {
	r, err := Setup()
	if err != nil {
		panic("internal error")
	}
	if len(os.Args) > 1 {
		r.Add(strings.Join(os.Args[1:], " "))
	} else {
		r.Print()
	}
}
