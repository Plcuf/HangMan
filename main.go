package main

import (
	"Hangman/fonctions"
	"fmt"
	"math/rand"
)

func main() {
	compteur := 10
	word := fonctions.GetWord(fonctions.GetWords())
	free := word[rand.Intn(len(word)-1)]
	tofind := []string{}
	for i := 0; i < len(word)-1; i++ {
		if word[i] == free {
			tofind = append(tofind, string(word[i]))
		} else {
			tofind = append(tofind, "_")
		}
	}

	for compteur > 0 {
		//fonctions.Clear()
		//Affichage(compteur)
		fmt.Println(tofind)
		fmt.Println(word)

		scan := fonctions.Scan()
		exist := false
		remaining := 0

		if len(scan) < 2 {
			for i := 0; i < len(word)-1; i++ {
				if scan == string(word[i]) {
					tofind[i] = scan
					exist = true
				} else if tofind[i] == "_" {
					remaining++
				}
			}
			if exist == false {
				compteur--
			}
			if remaining == 0 {
				fonctions.Win()
			}
		} else {
			fmt.Println(scan)
			fmt.Println(len(scan))
			fmt.Print(word)
			fmt.Println("AB")
			if scan == word {
				fmt.Println("sinon okkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk")
				fonctions.Win()
			} else {
				compteur += 2
			}
		}
	}
	fonctions.Lose()
}
