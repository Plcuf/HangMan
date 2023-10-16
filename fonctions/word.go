package fonctions

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func GetWords() []string {
	file, err := os.OpenFile("textes/noms_monstres.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := os.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}

	slice := []string{}
	word := ""

	for _, c := range data {
		c := string(c)
		if c == "\n" {
			slice = append(slice, strings.ToLower(word))
			word = ""
		} else {
			word = word + c
		}
	}
	return slice
}

func GetWord(s []string) string {
	word := s[rand.Intn(len(s))]
	fmt.Println(word)
	fmt.Println(len(word))
	return word
}
