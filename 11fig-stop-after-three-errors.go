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

	done := make(chan interface{})
	defer close(done)

	errCount := 0
	urls := []string{"gof", "https://hollywood.com", "man", "avraal", "ovich"}
	for result := range provoditGollivudVKvartireGofmana(done, urls...) {
		if result.Error != nil {
			fmt.Printf("mertvaia zona: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("soiti na net")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
