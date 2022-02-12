package itstime

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Clock struct {
	Now    time.Time
	Output io.Writer
}

func NewClock() Clock {
	return Clock{
		Now:    time.Now(),
		Output: os.Stdout,
	}
}

func (c Clock) CurrentTime() {
	fmt.Fprintf(c.Output, "It's %d past %d", c.Now.Minute(), c.Now.Hour())
}

func CurrentTime() {
	NewClock().CurrentTime()
}
