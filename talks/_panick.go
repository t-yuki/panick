package main

import (
	"io"
	"os"

	"github.com/t-yuki/panick"
)

func main() {
	// START OMIT
	defer func() {
		if panick.Panicked() { // HL
			io.WriteString(os.Stdout, "it's panicked!\n")
		}
		recover()
	}()
	panic("oops!")
	// END OMIT
}
