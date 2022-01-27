package main

import (
	"fmt"
	"runtime"
)

// main is entry point of any Go application.
// It must be in package main.
func main() {
	fmt.Printf("Hello World! Compiled with Go Version %s.\n",
		runtime.Version()[2:])
}
