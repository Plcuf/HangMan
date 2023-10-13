package main

import (
	"Hangman/fonctions"
	"fmt"
	"math/rand"
)

func main() {
	fonctions.Clear()
	compteur := 10
	word := GetWord()
	free := word[rand.Intn(len(word)-1)]
	tofind := [len(word)]string{}
	for i := 0; i < len(word)-1; i++ {
		if word[i] == free {
			tofind[i] = word[i]
		} else {
			tofind[i] = "_"
		}
	}

	for compteur > 0 {
		Affichage(compteur)
		fmt.Println(tofind)
		scan := Scan()
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
