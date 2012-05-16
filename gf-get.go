package main

import (
	"fmt"
	"gofortunes/fortunes"
)

func main() {
	var s = fortunes.GetFortune()
	fmt.Print(s)
}
