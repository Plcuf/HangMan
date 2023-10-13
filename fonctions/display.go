package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Write(s string) {
	fmt.Print(s)
}

func Clear() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}
