package main

import (
	"Hangman/fonctions"
	"fmt"
	"math/rand"
)

func main() {
	fonctions.Clear()
	compteur := 10
	word := fonctions.GetWord(fonctions.GetWords())
	free := word[rand.Intn(len(word)-1)]
	tofind := []byte{}
	for i := 0; i < len(word)-1; i++ {
		if word[i] == free {
			tofind = append(tofind, word[i])
		} else {
			tofind = append(tofind, '_')
		}
	}

	for compteur > 0 {
		Affichage(compteur)
		fmt.Println(tofind)
		scan := fonctions.Scan()
		for i := 0; i < len(word)-1; i++ {
			if scan == word[i] {
				tofind[i] = scan
			}
		}

		if scan == word {
			Win()
		}
	}
}
