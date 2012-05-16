package main

import (
	"gofortunes/fortunes"
)

func main() {
	db := new(fortunes.FortunesDB)
	db.BaseDir = "/usr/share/fortune"
	fortunes.AddFortune(db, "This is a joke\n", "jokes")
}
