package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// raskruchivatPovtorenie := func(
	// 	podkluchkaBoga <-chan interface{},
	// 	values ...interface{},
	// ) <-chan interface{} {
	// 	valueStream := make(chan interface{})
	// 	go func() {
	// 		defer close(valueStream)
	// 		for {
	// 			for _, v := range values {
	// 				select {
	// 				case <-podkluchkaBoga:
	// 					return
	// 				case valueStream <- v:
	// 				}
	// 			}
	// 		}
	// 	}()
	// 	return valueStream
	// }

	take := func(
		podkluchkaBoga <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-podkluchkaBoga:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	podkluchkaBoga := make(chan interface{})
	defer close(podkluchkaBoga)

	rand := func() interface{} {
		return rand.Int()
	}

	for num := range take(podkluchkaBoga, repeatFn(podkluchkaBoga, rand), 10) {
		fmt.Println(num)
	}
}
