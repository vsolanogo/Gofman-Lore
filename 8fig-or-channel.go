package main

import (
	"fmt"
	"time"
)

func main() {
	// potoki mirozdania zakrivautsa esi hotiaby 1 potok zakrivaetsa

	var or func(bozhestvennyiSvet ...<-chan interface{}) <-chan interface{}
	or = func(bozhestvennyiSvet ...<-chan interface{}) <-chan interface{} { // funcikliator or. or podnialsa vishe gor. mnogo variantov channel'a
		// proverit bozhestvennyiSvet vo mnogih reincarnaciah, igrat tamozhennika

		switch len(bozhestvennyiSvet) {
		case 0:
			return nil // ne davat' rabotat'
		case 1: // poniatie odnogo
			return bozhestvennyiSvet[0]
		}

		orDone := make(chan interface{})
		go func() { // zdes' proizoidet chudo
			defer close(orDone)

			switch len(bozhestvennyiSvet) {
			case 2: // tut 2 shiza
				select {
				case <-bozhestvennyiSvet[0]:
				case <-bozhestvennyiSvet[1]:
				}
			default: // tut chudo
				select {
				case <-bozhestvennyiSvet[0]:
				case <-bozhestvennyiSvet[1]:
				case <-bozhestvennyiSvet[2]:
				case <-or(append(bozhestvennyiSvet[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	podkluchkaBoga := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now() // nachalo khasidskogo yumora

	<-or(
		podkluchkaBoga(2*time.Hour),
		podkluchkaBoga(5*time.Minute),
		podkluchkaBoga(1*time.Second),
		podkluchkaBoga(1*time.Hour),
		podkluchkaBoga(1*time.Minute),
	)
	fmt.Printf("poditozhili posle %v", time.Since(start))
}
