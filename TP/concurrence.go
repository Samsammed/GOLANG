package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go maFonction(&wg)
	fmt.Println("Fin du programme")
	wg.Wait()
}
func maFonction(wg *sync.WaitGroup) {
	fmt.Println("j'ai  fini  !")
	wg.Done()
}
