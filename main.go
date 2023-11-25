package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Plcuf/HangMan/fonctions"
)

func main() {
	game()
}

func game() {
	compteur := 10
	brokenword := fonctions.GetWord(fonctions.GetWords("textes/noms_monstres.txt"))
	word := ""
	alreadysaid := []string{}

	for i := 0; i < len(brokenword)-1; i++ {
		word += string(brokenword[i])
	}

	brokenword = strings.ToLower(brokenword)
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
		fmt.Println(alreadysaid)

		scan := fonctions.Scan()
		exist := false
		remaining := 0
		insaid := true

		if len(scan) < 2 {
			for i := 0; i < len(alreadysaid); i++ {
				if scan == alreadysaid[i] {
					fmt.Println("you already said this letter")
					time.Sleep(1 * time.Second)
					insaid = false
				}
			}
			if insaid == true {
				alreadysaid = append(alreadysaid, scan)
			}
			insaid = false
			for i := 0; i < len(word); i++ {
				if scan == string(brokenword[i]) {
					tofind[i] = string(brokenword[i])
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
