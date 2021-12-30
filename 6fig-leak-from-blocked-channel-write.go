package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// razbor problemi semantiki kogda nado poslat' bozhestvennyi svet iz tualeta nazad v avtomatiku no ne poluchaetsa

	poslatBozhestvennyiSvetIzTualeta := func() <-chan int {
		luchiSveta := make(chan int)
		go func() {
			defer fmt.Println("poslatBozhestvennyiSvetIzTualeta closure exited.") // ne rabotaet
			defer close(luchiSveta)
			for {
				luchiSveta <- rand.Int()
			}
		}()

		return luchiSveta
	}

	luchiSveta := poslatBozhestvennyiSvetIzTualeta()
	fmt.Println("3 lucha sveta:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-luchiSveta)
	}
}
