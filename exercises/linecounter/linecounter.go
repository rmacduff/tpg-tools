package linecounter

import (
	"bufio"
	"io"
	"os"
)

type Counter struct {
	Input io.Reader
}

func NewCounter() Counter {
	return Counter{
		Input: os.Stdin,
	}
}

func (c Counter) Lines() int {
	scanner := bufio.NewScanner(c.Input)
	var count int
	for scanner.Scan() {
		count++
	}
	return count
}

func Lines() int {
	return NewCounter().Lines()
}
