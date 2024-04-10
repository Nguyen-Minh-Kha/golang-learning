package main

import (
	"fmt"
	"sync"
)

func maFonction(wg *sync.WaitGroup) {
    fmt.Println("j'ai fini")
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
    go maFonction(&wg)
	fmt.Println("fin du programme")
}