package main

import (
	"log"
	"os"
)

func main() {

    log.SetPrefix("ERROR: ")
    // output 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
    log.SetOutput(os.Stdout)
}
