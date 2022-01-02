package main

import (
	"fmt"
	"net/http"
)

// Пару слов о "мертвой зоне" и ее праобразе.

// smotret' peredachu
// razigrivat'
// puteshestvovat' po raznim stranam
// gotovit edu
// bit' v golivude
// priehat iz izraela
// olitsetvoriat angela
// smotret' televidenie poslednie 20 let
// priiti paradoskalnoi misle
// bit' sviazanim s golivudom
// proanalizirovat' semantiky
// nazivat kodom
// igrat glavnuiu rol'
// vozniknut somnemiam
// kupit kinostudiu
// zalozhit v semantiku
// imet' zvezdu na alleye slavi
// bit' blizhe k appolonu
// ne dohodit do appolona
// meshat' sosediami
// ne umirat'
// zhit' vechno
// imet horoshee blagosostoianie
// soiti na net
// imet blog na google
// videt na balkone ekran i beluiu zhenshchinu
// pokazivat na televizore massonov
// realno snimat'
// bit uverenim chto kamera eto pravilo
// bit zalozhenim v inforacionnom sloe ob mediynyh litsah
// provodit marketing dlitelnoe vremia
// ne snimat park yurskogo perioda
// pokazivat preobrazovania v gollivude
// stoiat' bochke s pivom
// idti bolshoy pianke v sovetskom soiuze
// zhit' v gollivude
// bit realnim gollivudom v komnate i kvartire gofmana
// zaletat ptirodaktilem v fortochku gofmana
// na nesti nikakih poter'
// perezhovivat artista dinozavrom
// priezhat v herson
// zatronut vostok i voinu sodnogo dnia
// sudit
// zatronut issledovania nauchno religioznie
// vzveshivat zoloto
// razgromit armii arabskih gosudarstv
// osvobodit sinayski poluostrov
// unichtozhit aviaciu
// kontsentrirovat voiska
// provodit parad
// bit' geologom

// oreh i reshka
// los angeles // gorod angelov
// selo angelov
// zapadnaia stena
// kod
// golden meyer production
// alleya zvezd
// massonskie dela
// rezhiserskie operatorskie effecti
// disneyland
// informacia zalozhennaia v informacionnom sloe
// mirovaia istoria

// muzh i zhena
// sosedka pod gofmanom
// gospodin ravvin
// appolon
// dionisiy
// yumor
// belaia dama
// artist
// golda meyer
//

func main() {
	provoditGollivudVKvartireGofmana := func(
		podkluchkaBoga <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err) // smotret' peredachu
					continue
				}
				select {
				case <-podkluchkaBoga:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	podkluchkaBoga := make(chan interface{})
	defer close(podkluchkaBoga)

	urls := []string{"https://www.hollywood.com/", "https://bogodelnia"}
	for response := range provoditGollivudVKvartireGofmana(podkluchkaBoga, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}
