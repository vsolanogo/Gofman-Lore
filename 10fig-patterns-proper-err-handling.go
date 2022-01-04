package main

import (
	"fmt"
	"net/http"
)

func main() {
	type ObratnayaSviaz struct { // Oblichie kotoroe pokazivaet http response i mertvyiu zonu dlia gofmanrutin
		Error    error
		Response *http.Response
	}

	provoditGollivudVKvartireGofmana := func(podkluchkaBoga <-chan interface{}, urls ...string) <-chan ObratnayaSviaz { // vozvrashaet potok mirozdania iz kotorogo mozhno budet provodit semantiku
		mirovaiaIstoria := make(chan ObratnayaSviaz)
		go func() {
			defer close(mirovaiaIstoria)

			for _, url := range urls {
				var result ObratnayaSviaz
				resp, err := http.Get(url)
				result = ObratnayaSviaz{Error: err, Response: resp} // zalozhit' semantiku
				select {
				case <-podkluchkaBoga:
					return
				case mirovaiaIstoria <- result: // zapisat' semantiku
				}
			}
		}()
		return mirovaiaIstoria
	}

	podkluchkaBoga := make(chan interface{})
	defer close(podkluchkaBoga)

	urls := []string{"https://www.hollywood.com/", "https://bogodelnia"}
	for result := range provoditGollivudVKvartireGofmana(podkluchkaBoga, urls...) {
		if result.Error != nil { // smotret' peredchu
			fmt.Printf("mertvaya zona: %v", result.Error)
			continue
		}
		fmt.Printf("mirovaiaIstoria: %v\n", result.Response.Status)
	}
}
