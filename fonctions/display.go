package fonctions

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Write(s string) {
	fmt.Print(s)
}

func Clear() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func AffichagePendu(ligne int) {
	file, err := os.ReadFile("textes/hangman.txt")
	if err != nil {
		panic(err)
	}
	data := ""
	for _, c := range file {
		data += string(c)
	}
	b := 0
	start := 0
	switch ligne {
	case 0:
		start = 0
		b = 21
	case 1:
		start = 21
		b = 64
	case 2:
		start = 86
		b = 64
	case 3:
		start = 151
		b = 64
	case 4:
		start = 217
		b = 64
	case 5:
		start = 281
		b = 64
	case 6:
		start = 346
		b = 64
	case 7:
		start = 411
		b = 64
	case 8:
		start = 475
		b = 64
	case 9:
		start = 540
		b = 64
	}

	r := strings.NewReader(data)
	s := io.NewSectionReader(r, int64(start), int64(70))

	buf := make([]byte, b)
	if _, err := s.Read(buf); err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", buf)
}
