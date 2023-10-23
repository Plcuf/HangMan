package main

import (
	"Hangman/fonctions"
	"fmt"
	"math/rand"
)

func main() {
	compteur := 10
	brokenword := fonctions.GetWord(fonctions.GetWords())
	word := ""
	//alreadysaid := ""

	for i := 0; i < len(brokenword)-1; i++ {
		word += string(brokenword[i])
	}

	free := word[rand.Intn(len(word)-1)]
	tofind := []string{}

	for i := 0; i < len(word); i++ {
		if word[i] == free {
			tofind = append(tofind, string(word[i]))
		} else {
			tofind = append(tofind, "_")
		}
	}

	for compteur > 0 {
		fonctions.Clear()
		fonctions.AffichagePendu(10 - compteur)
		fmt.Println(tofind, " ", compteur, "ꨄ")

		scan := fonctions.Scan()
		exist := false
		remaining := 0

		if len(scan) < 2 {
			for i := 0; i < len(word); i++ {
				if scan == string(word[i]) {
					tofind[i] = scan
					exist = true
				} else if tofind[i] == "_" {
					remaining++
				}
			}
			if !exist {
				compteur--
			}
			if remaining == 0 {
				fonctions.Win()
			}
		} else {
			if scan == word {
				fonctions.Win()
			} else {
				compteur -= 2
			}
		}
	}
	fonctions.Lose()
}
