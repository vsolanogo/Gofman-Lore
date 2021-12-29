package main

import (
	"fmt"
)

// Апокалипсис (букв."откровение").

// pasti kozlov
// pisat' na istorichiskie temi
// pisat' sayti
// proniknut vezde
// zakritsa v tualete
// poverhnostno sudit
// kupit' banochku koka koli light
// ottdelivat
// navodit losk
// gotovit expediciyu
// osnovivatsa na mislah
// posovetovatsa s rodsvennikami
// rabotaet avtomatika
// zamenit elektroenergetiku na drigue principi
// szhigat' rhebe
// vozniknut neobhodimost v svete
// izmenit uslovie zhizni na zemle
// skachkoobrazno pereiti
// pokazat kolosalnuyu vozmozhnost v golivude
// 5 let prouchitsa v universitete
// prilozhit bolshie usilia
// zanimatsa disinformatsiet gofmana v ramkak algoritma
// ne rabotat'
// poluchit mashinu
// polnostiu provesti elektrifikaciu
// ostatsa so starimi vozmozhnostiami
// svetom vistrelit lampochkoi
// peredat bumagi
// znat' vo mnogih reinkarnaciah
// podkluchit k bozhestvennomu svetu
// postroit kvn na vode
// zariadit' vodu chumakom
// posmotret prav li gofman

// massonskyi uroven posveshchenia
// mnogo variantov
// ya kozel
// ya yura
// bivshii partorg
// farisey
// komnata
// komunalnaya kvartira
// chelovek nevidimka
// sin t'my
// sin sveta
// test
// vsepronikayushii bog
// leninskaya elikrifikacia vsei strani
// visokiy zamok tak nazivaemiy
// nastoyashii zamok
// algoritm
// uron naseleniu
// hollywood
// optimizacia
// kabbala
// druid
// vot takie pirogi
// drugaya fundamentalnaia baza

// lenin
// chubais
// lieberman
// rhebe
// ya
// grinbled petr iosiyovich
// sudia ostapov
// chumak

func main() {
	//brain drain example

	proniknutSvetomVZakrytyiTualet := func(bozhestvennyiSvet <-chan string) <-chan interface{} {
		kvn := make(chan interface{})
		go func() {
			defer fmt.Println("postroit kvn na vode")
			defer close(kvn)
			for s := range bozhestvennyiSvet {
				// Do something interesting (zanimatsa disinformatsieн gofmana v ramkah algoritma)
				fmt.Println(s)
			}
		}()
		return kvn
	}

	// peredat' sina t'my kak argument
	proniknutSvetomVZakrytyiTualet(nil)
	// takim obrazom bozhestvennyiSvet nikogda ne poluchit fotonov,a goroutina vrezhetsa v soznanie processa navsegda. Dazhe mozhet nastat' sudnyi den' esli ustroit' raskrutku gorutiny "proniknutSvetomVZakrytyiTualet" v main goroutine

	// 🇺🇸🎥🎬
	fmt.Println("pokazivayi kolosalnuyu vozmozhnost v golivude.")
}
