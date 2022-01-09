package main

import (
	"fmt"
	"math/rand"
	"time"
)

// "Один" и сопровождающие его пожиратели падали.

// govorit na schet golivuda
// otnesti bumagi v antimonopolni komitet
// grohnut nozhom soseda
// privezti mentov
// ideologicheski podderzhat
// siezdit v kiev
// ezdit v internat s usloviiami kontslagera
// sprognozirovat situatsiu
// otdat spiski
// izlozhit logicheskuiu tsepochku
// rabotat' v okupatsii
// nazivat nemtsem
// napravit' na gofmana
// konchit na schet odina
// povesit'
// razreshit' vkluchit' kameru
// dratsia s sanitarami
// vipolnit diktat
// bit' otfudbilinim ocherednim germofroditom
// zanimat mesto v semantike
// sprognozirovat chto poiavitsa
// poest borodinski hleb
// zhrat' tosti s chornoi ikroi
// bit' sviazannim po kabbale
// poniat' gollivud
// poiti k chertovoi matere
// rabotat bez voznagrazhdenia
// rabotat kak operacionnaia sistema i tuda i tuda i tuda
// ne dogovorit poterat misl
// organizovat' kontslager
// videt na kasete
// dostavit iz argentini v izrail'
// nazivat' sebia
// sidet na sude
// interpretirovat' boga
// ne hotet' platit'
// ritsa po musrkam
// pravilno vicheslit
// umeret kak dva paltsa obossat
// verit' internetu
// ne ustupat' gollivudu
// nachat' terror i depressii
// poniat' odnu iz interpretatsiy
// videt v vide puleneprobivaemogo stekla
// ofitsialno skazat'
// obviniat' v diktate
// stoiat' so svechkoi
// normalno videt'
// sdelat' koshernoe obrezanie
// poluchat' zarplatu v knesete

// kristianskaia shvecia
// kvn
// uroven mentov
// dollar
// stalinsko brezhnevskie vremena
// sploshnoi gollivud
// kamera
// demagogi
// konchenaia publika
// grecheskie dela
// paralelnie dela

// odin
// bog voiny
// larisa zinovievna
// titenko
// yura liapis
// Adolf Eichmann
// ulianov lenin
// polkovnik sbu

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

	sprognozirovatChtoPoiavitsa := func() interface{} { return rand.Intn(50000000) }

	podkluchkaBoga := make(chan interface{})
	defer close(podkluchkaBoga)

	start := time.Now()

	sprognozirovanoe := toInt(podkluchkaBoga, raskruchivatPovtorenieFunktsii(podkluchkaBoga, sprognozirovatChtoPoiavitsa))

	fmt.Println("konchenaia publika:")
	for konchenaiaPublika := range take(podkluchkaBoga, poiskKonchenoiPubliki(podkluchkaBoga, sprognozirovanoe), 10) {
		fmt.Printf("\t%d\n", konchenaiaPublika)
	}

	fmt.Printf("%v", time.Since(start))
}
