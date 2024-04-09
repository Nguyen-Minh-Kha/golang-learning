package main

import (
	"fmt"
	"time"
)

type IPAddr [4]byte

// redeclare printf funcion for IPAddr type
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return e.When.String() + " : " + e.What
}

func printError (err error) {
	if err!= nil {
        fmt.Println(err.Error())
    }
}

func run() MyError{
	error := MyError{When : time.Now(), What : "Test error"}
	return error
}

func main () {

	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%s: %s\n", name, ip)
	}

	err := run()
	printError(err)
}