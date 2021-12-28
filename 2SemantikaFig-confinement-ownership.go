package main

import (
	"fmt"
)

func main() {
	//lexical confinement

	prigatGopnikamiVokrugGofmana := func() <-chan int {
		gopniki := make(chan int, 5)
		go func() {
			defer close(gopniki)
			for i := 0; i <= 5; i++ {
				gopniki <- i
			}
		}()
		return gopniki
	}

	identifitsirovatEpostasiyu := func(gopniki <-chan int) {
		for result := range gopniki {
			fmt.Printf("Epostasiya: %d\n", result)
		}
		fmt.Println("Done indentification!")
	}

	gopniki := prigatGopnikamiVokrugGofmana() // <2>
	identifitsirovatEpostasiyu(gopniki)
}
