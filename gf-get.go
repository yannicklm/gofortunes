package main

import (
	"flag"
	"fmt"
	"gofortunes/fortunes"
	"os"
)

func main() {
	var baseDir = flag.String("db-path", "", "fortune base dir")
	flag.Parse()
	if *baseDir == "" {
		fmt.Println("--db-path flag is required")
		os.Exit(2)
	}
	db := new(fortunes.FortunesDB)
	db.BaseDir = *baseDir
	var s = fortunes.GetFortune(db, "jokes")
	fmt.Printf("Got this fortune: %s", s)
}
