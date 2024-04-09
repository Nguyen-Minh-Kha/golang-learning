package main

import (
    "fmt"
)

// Empty interface can be anything
type emptyInterface interface {}

func PrintIt (input interface{}) {
	switch v := input.(type) {
		case string:
            fmt.Println("The type is string")
        case int:
            fmt.Println("The type is int")
		default:
            fmt.Printf("idk this type %T\n", v)
	}
}
func main() {
	x := 123
    PrintIt(x)
}