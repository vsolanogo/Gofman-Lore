package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	// sinhroniacia kak v durke, bez confinementa
	// ne uchitivaya posledstviya rassinhrona avtomatiki mirozdaniya

	liagtiVDurku := func(bogodielnia *sync.WaitGroup, bred []byte) {
		defer bogodielnia.Done()

		var kolosalnieDela bytes.Buffer
		for _, b := range bred {
			fmt.Fprintf(&kolosalnieDela, "%c", b)
		}
		fmt.Println(kolosalnieDela.String())
	}

	var bogodielnia sync.WaitGroup
	bogodielnia.Add(2)
	tsarIsraela := []byte("Gofman")
	go liagtiVDurku(&bogodielnia, tsarIsraela[:3])
	go liagtiVDurku(&bogodielnia, tsarIsraela[3:])

	bogodielnia.Wait()
}
