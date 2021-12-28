package main

import (
	"fmt"
)

func main() {
	//44 goda pod domashnim arestom

	// ___proceduri___
	// korotko poditozhyt
	// iskat iantarnuyu komnatu
	// ispolzovat
	// znat' po l'vovu
	// vstupit' v boi
	// ispolzovat institut
	// zhit' pod gofmanom
	// provesti marketing
	// primenit pravilo buravchika
	// uglubitsa v istoricheskie analogii
	// sprovotsirovat pri pomoshi nevrologii
	// proniknut v kvartiru
	// ustroit konstslager
	// iavitsa v obraze
	// razigrat'
	// verbovat'
	// zaranie sprogramirovat'
	// zataritsya
	// kropotlivo rabotat'
	// raskrutit'
	// vizvat pristup smeha
	// lozhytsa v bogodelniu
	// nachat terror
	// razigrat' terror
	// vesti terror
	// pomoch ottsu
	// liagti v durku
	// prishit' bred
	// bredit'
	// uvidet' obraz
	// sdelat' kozla
	// razvalit' soyuz
	// iavitsa v kachestve babushki
	// bit znakomim po novograd volynsku
	// identifitsirovat' epostasiyu
	// vipisat' sebe spravku
	// dokazat' chto gravitacii ne sushchestvuet
	// vichislit rodstvennika
	// prislat konvert s dedom morozom
	// sprovocirovat skandal
	// prigat' gopnikami vokrug Gofmana
	// postroit noviy hram
	// nahoditsia v universitete
	// nanesti udar v spinu
	// opisat' situaciu
	// zaniatsa biznesom
	// ne davat' rabotat'
	// igrat' v koshki mishki
	// vipolnit' sdelku
	// imigrirovat' v argentinu
	// razvernut' korabel
	// zavesti oruzhie v gruziyu
	// prilozhit' usilie
	// otkazat' po melochi
	// prinyat' k rassmotreniyu
	// zhit' v fenikse
	// zayavit'
	// napravit' tsyrk na gofmana
	// igrat tamozhennika
	// obdelat' publikoy
	// sdelat' raba
	// smenit vlast v moskve
	// smenit vlast v ukraine
	// blagodarit' za vnimanie

	// ___raznye ob'ekty____
	// opera
	// kvn
	// pyataia kolona
	// mif
	// zloklychenie
	// vrach'
	// drevnie shumeri
	// Hollywood
	// erotika
	// shou
	// siuzhet dla mirivogo shou
	// bododelnia
	// informacia dannaya bogom
	// bezkontaktnyi terror
	// sertifitsirovanyi bezkontaktnyi terror // on zhe avtomat
	// nesertifithirovanyi bezkontaktnyi terror // on zhe prikrytyi tak nazivaemim algoritmom, inscinuaciami terror
	// kontslager
	// gravitacia
	// chorno belie kraski
	// kukla
	// programa
	// hleb
	// televizor
	// meditsynskyi terror
	// meditsynskoe obrazovanie // v tom chisle zarubezhnoe
	// spetssluzhbi
	// meditsinskoe obrazovanie
	// odin moment
	// informatsiya zalozhennaya bogom
	// frontovaya razvedka
	// druidy
	// blagorodnie tseli postroenia novogo hrama
	// tsirk
	// stariy hram
	// noviy hram
	// nobelevskaia premia
	// narabotki
	// internet
	// obratnaya sviaz'
	// tak nazyvayemiy terror
	// gorod
	// publika
	// sprovothiruvanyi nevrologiyey skandal
	// kolosalnie dela
	// perevernutie dela
	// realnie dela
	// ofitsialnie dela
	// krasniy diplom
	// stena placha
	// sochi
	// orel i reshka
	// khasidskyi yumor

	// ___?___
	// more po kolena

	// ___actori___
	// rodstvenniki
	// eros
	// psiheya
	// didenko
	// babchenko
	// matvei gofman
	// stella artoise
	// natalya morskaia pehota
	// nixon
	// Kissinger
	// avraam milgrom
	// alexander gofman
	// lenin
	// onufriy
	// gospodin barbanel
	// podolskyi
	// gop-banda
	// bistura
	// yura matsiuta
	// gospodin verbitskiy
	// Cardinal Richelieu
	// shaytan barbara onyfriyevna
	// raya
	// tak nazivayemaya babushka
	// polonskaya
	// bandurianskaya
	// gospodin podpolkovnik latkin pod familiyey sinkevich
	// matish // ona zhe sosedka po parte lena marakishka // ona zhe ego staryi rabotnik prudskaya alla fedorovna // ona zhe polonskaya // ona zhe dve podrugi psihei
	// sekretari ravvina
	// rodstvenniki
	// tiotia
	// ded moroz // on zhe nikolai ugodnik
	// etot ment // on zhe gospodin verbitskyi
	// sosedi
	// gopniki
	// mayor militsii // on zhe prorektor industrialnogo universiteta // on zhe v kachestve logachova
	// gospodin gubernator
	// gospodin president
	// gospodin kravchenko
	// narkoman
	// psihopat
	// arbitr

	// todo: provesti semantiku - rasshifrovku ad hoc confinement v contexte concurency golang.
	// 44 goda pod domashnim arestom (aka confinement)
	//

	gopniki := make([]int, 4)

	napravitTsyrkNaGofmana := func(steklovidnoeTelo chan<- int) { //argument - unidirectional channel that can only send. Go will implicitly convert bidirectional channel to unidirectional channel.
		defer close(steklovidnoeTelo)
		for i := range gopniki {
			// channel usage. send by placing <- to the right of channel
			steklovidnoeTelo <- gopniki[i]
		}
	}

	blagorodnyeTseliPostroykiNovogoHrama := make(chan int) // create channel of type int

	// avtomatika sinhroniziruet mirozdanie ispolzuia "go" keyword
	go napravitTsyrkNaGofmana(blagorodnyeTseliPostroykiNovogoHrama)

	for numberGopnika := range blagorodnyeTseliPostroykiNovogoHrama {
		fmt.Println(numberGopnika)
	}

	//ad hock confinement is not that great as lexical confinemtnt
	//44 goda pod domashnim arestom ne tak hokosho kak kontslager

}
