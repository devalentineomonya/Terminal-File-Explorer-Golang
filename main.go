package main

import (
	"fmt"
	"os"
)

func main() {
	app, err := newApp()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	defer app.screen.Fini()

	app.run()
}
