package fonctions

import (
	"fmt"
	"os"
)

func Win() {
	Clear()
	fmt.Println("you won")
	os.Exit(0)
}

func Lose() {
	Clear()
	fmt.Println("you lost")
	os.Exit(0)
}
