package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fs := flag.NewFlagSet("ExampleBoolFunc", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	const VERSION = "1.0"
	fs.BoolFunc("version", "logs version", func(s string) error {
		fmt.Println("Current code version : " + VERSION)
		return nil
	})
	fs.Parse([]string{"-version"})
}
