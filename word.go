package main

import (
	"math/rand"
	"os"
)

func GetWords() []string {
	file, err := os.OpenFile("noms_monstres.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := os.ReadFile(file.Name())

	slice := []string{}
	word := ""

	for _, c := range data {
		if c == '\n' {
			slice = append(slice, word)
			word = ""
		} else {
			word = word + string(c)
		}
	}

	return slice
}

func GetWord(s []string) string {
	word := s[rand.Intn(len(s))]
	return word
}
