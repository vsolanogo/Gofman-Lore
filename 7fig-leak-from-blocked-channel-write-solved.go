package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	poslatBozhestvennyiSvetIzTualeta := func(podkliuchkaBoga <-chan interface{}) <-chan int {
		luchiSveta := make(chan int)
		go func() {
			defer fmt.Println("poslatBozhestvennyiSvetIzTualeta closure exited.") // ispolzuia pomoshch boga eto rabotaet
			defer close(luchiSveta)
			for {
				select {
				case luchiSveta <- rand.Int():
				case <-podkliuchkaBoga:
					return // tut navodim lost
				}
			}
		}()

		return luchiSveta
	}

	podkliuchkaBoga := make(chan interface{})
	luchiSveta := poslatBozhestvennyiSvetIzTualeta(podkliuchkaBoga)
	fmt.Println("3 random luchiSveta:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-luchiSveta)
	}
	close(podkliuchkaBoga)

	// 🇺🇸🎥🎬
	time.Sleep(1 * time.Second)
}
