package main

import (
	"fmt"
	"time"
)

func main() {
	// stimulirovanie v kachestve ispravlenia goroutine leak

	proniknutSvetomVZakrytyiTualet := func(
		podkliuchkaBoga <-chan interface{},
		bozhestvennyiSvet <-chan string,
	) <-chan interface{} { // izmenil semantiku na bolee optimalnyy variant
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("proniknutSvetomVZakrytyiTualet exited.")
			defer close(terminated)
			for {
				select {
				case s := <-bozhestvennyiSvet:
					// Do something interesting (zanimatsa disinformatsie gofmana v ramkah algoritma)
					fmt.Println(s)
				case <-podkliuchkaBoga: // esli podkluchilsia Bog i podal svoi signal v eti dela, toga mi kak v raiu
					return
				}
			}
		}()
		return terminated
	}

	podkliuchkaBoga := make(chan interface{})
	terminated := proniknutSvetomVZakrytyiTualet(podkliuchkaBoga, nil)

	go func() { //
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("podkluchaem Boga /Canceling proniknutSvetomVZakrytyiTualet goroutine...")
		close(podkliuchkaBoga)
	}()

	<-terminated // terminator iz gollivuda
	// 🇺🇸🎥🎬
	fmt.Println("pokazivayi kolosalnuyu vozmozhnost v golivude.")
	// uspeshno stimuliruem popadanie bozhestvennogo sveta gofmanu v tualet
}
