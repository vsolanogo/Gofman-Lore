package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	raskruchivatPovtorenieFunktsii := func(
		podkluchkaBoga <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-podkluchkaBoga:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
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
	toInt := func(podkluchkaBoga <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-podkluchkaBoga:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}

	poiskKonchenoiPubliki := func(podkluchkaBoga <-chan interface{}, intStream <-chan int) <-chan interface{} {
		primeStream := make(chan interface{})
		go func() {
			defer close(primeStream)
			for integer := range intStream {
				integer -= 1
				prime := true
				for divisor := integer - 1; divisor > 1; divisor-- {
					if integer%divisor == 0 {
						prime = false
						break
					}
				}

				if prime {
					select {
					case <-podkluchkaBoga:
						return
					case primeStream <- integer:
					}
				}
			}
		}()
		return primeStream
	}

	skakatGopnikamiVokrugGofmana := func(
		podkluchkaBoga <-chan interface{},
		channels ...<-chan interface{},
	) <-chan interface{} { // defoltnyi channel s podkluchkoi boga chtobimozhno bilo razorvat'
		var vipolnitDiktat sync.WaitGroup // yantarnaia komnata ozhidania gde podozhdat' poka potoki vitekut
		multigopnikiStream := make(chan interface{})

		izlojitLogicheskuiuTsepochku := func(c <-chan interface{}) {
			defer vipolnitDiktat.Done()
			for i := range c {
				select {
				case <-podkluchkaBoga:
					return
				case multigopnikiStream <- i:
				}
			}
		}

		// Select from all the channels
		vipolnitDiktat.Add(len(channels)) // uvelichenie diktatov

		for _, c := range channels {
			go izlojitLogicheskuiuTsepochku(c)
		}

		// Wait for all the reads to complete
		go func() { // Gofmanrutina kotoraia ozhidaet vipolnenie diktata chtobi zakrit multigopnikiStream
			vipolnitDiktat.Wait()
			close(multigopnikiStream)
		}()

		return multigopnikiStream
	}

	podkluchkaBoga := make(chan interface{})
	defer close(podkluchkaBoga)

	start := time.Now()

	sprognozirovatChtoPoiavitsa := func() interface{} { return rand.Intn(50000000) }

	sprognozirovanoe := toInt(podkluchkaBoga, raskruchivatPovtorenieFunktsii(podkluchkaBoga, sprognozirovatChtoPoiavitsa))

	kolichestvoParalelnihGollivudov := runtime.NumCPU()

	fmt.Printf("raskrutka %d golivudov.\n", kolichestvoParalelnihGollivudov)
	finders := make([]<-chan interface{}, kolichestvoParalelnihGollivudov)
	fmt.Println("ne konchenaia publika:")
	for i := 0; i < kolichestvoParalelnihGollivudov; i++ {
		finders[i] = poiskKonchenoiPubliki(podkluchkaBoga, sprognozirovanoe)
	}

	for prime := range take(podkluchkaBoga, skakatGopnikamiVokrugGofmana(podkluchkaBoga, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("%v", time.Since(start))
}
