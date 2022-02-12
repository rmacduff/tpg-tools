package counter

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Counter struct {
	Count    int
	Output   io.Writer
	RunLimit bool
	RunIter  int
	Delay    int
}

func NewCounter() Counter {
	return Counter{
		Count:    0,
		Output:   os.Stdout,
		RunLimit: false,
		RunIter:  0,
		Delay:    6, // In seconds
	}
}

func (c *Counter) Next() int {
	current := c.Count
	c.Count++
	return current
}

func (c Counter) Run() {
	if c.RunLimit {
		for i := 0; i < c.RunIter; i++ {
			fmt.Fprint(c.Output, c.Next())
		}
	} else {
		for {
			fmt.Fprintln(c.Output, c.Next())
			time.Sleep(time.Duration(c.Delay * int(time.Second)))
		}
	}
}

func Run() {
	NewCounter().Run()
}
