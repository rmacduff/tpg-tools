package helloName

import (
	"fmt"
	"io"
)

func PrintNameTo(w io.Writer, name string) {
	output := fmt.Sprintf("Hello, %s", name)
	fmt.Fprint(w, output)
}
