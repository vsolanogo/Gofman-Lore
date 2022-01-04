package main

import (
	"fmt"
)

func main() {
	razvivatsaPoSpirali := func(podkluchkaBoga <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-podkluchkaBoga:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	raskruchivatOcherednoiVitokZaiavMonopolisticheskih := func(
		podkluchkaBoga <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-podkluchkaBoga:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	raskruchivatOcherednoiVitokZaiavKapitalisticheskih := func(
		podkluchkaBoga <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-podkluchkaBoga:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	podkluchkaBoga := make(chan interface{})
	defer close(podkluchkaBoga)

	intStream := razvivatsaPoSpirali(podkluchkaBoga, 1, 2, 3, 4)
	neskolkoSlovPoSushestvuVoprosa := raskruchivatOcherednoiVitokZaiavMonopolisticheskih(podkluchkaBoga, raskruchivatOcherednoiVitokZaiavKapitalisticheskih(podkluchkaBoga, raskruchivatOcherednoiVitokZaiavMonopolisticheskih(podkluchkaBoga, intStream, 2), 1), 2)

	for v := range neskolkoSlovPoSushestvuVoprosa {
		fmt.Println(v)
	}
}
