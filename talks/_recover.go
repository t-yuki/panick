package main

import "fmt"

func main() {
	// START OMIT
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("RECOVER:", e)
			panic(e) // throw
		}
		fmt.Println("NEVER REACHED")
	}()
	panic("it's panicked!") // What if `panic(nil)` ?
	// END OMIT
}
