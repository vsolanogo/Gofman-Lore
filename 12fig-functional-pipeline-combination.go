package main

import (
	"fmt"
)

// Несколько слов по существу вопроса.

// organizovat popadanie v ihnuiu bogodelniu
// bit ne dla nashego rila
// rasskazat hohmu
// bistupat
// rasskazat o stole
// prikinit'
// doprashivat sovetskuiu radistku
// uzurpirovat mesto tsaria solomona
// dobista smeshchenie angela
// razvivatsa po spirali
// morozit menia
// bit pohozhim na shtirlitza
// idti domoi
// ne protivorechit
// perevesti proces s nee na menia
// razigrat provocatsiu
// poiavliatsa ievreiskim licam
// otvesti po publike
// raskruchivat ocherednoi vitok zaiav kompartiinyh
// pet pesniu
// govorit' cherez gubu
// bit iz toi zhe operi
// iskat duh ottsa
// iskat kto sozdal sibir'
// iskat predsedatela nabludatelnogo soveta
// naiti togo knto reshat nichego ne mozhet a tolko nabludat'
// devat' knigi
// osdtavit' stol
// bit iz sibiri
// soderzhat knigi
// idti v evropu
// privit' bolee horoshi vkus
// prevoshodit' sovetskie dela
// bit' sviazanym s psihilogiei i oblasti vkusa
// videt' shtirlitza v crimu
// sorvat' mehovuiu shapku porivom vetra
// raskruchivat' v kino

// otel Watergate
// hohma
// film o shtirlice
// 17 mgnovenii moskvi
// ocherednaia hohma
// stop
// sibir'
// mirozdanie
// avtomatika
// kapitalisticheskie monopolisticheskie dela
// dedovshchina
// ruminskaia mebel
// simvol

// shtirlitz
// tihonov
// gospodin gavrliuk
// otec avraam
// agent gestapo

func main() {
	raskruchivatOcherednoiVitokZaiavMonopolisticheskih := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}
	raskruchivatOcherednoiVitokZaiavKapitalisticheskih := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range raskruchivatOcherednoiVitokZaiavKapitalisticheskih(raskruchivatOcherednoiVitokZaiavMonopolisticheskih(ints, 2), 1) {
		fmt.Println(v)
	}
}
