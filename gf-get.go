package main

import (
	"fmt"
	"gofortunes/fortunes"
	"os"
)

func usage() {
	fmt.Println("Usage: %s DB_PATH [CATEGORY]", os.Args[0])
	os.Exit(2)
}

func main() {
	var fortune, baseDir, category string
	var err error
	if len(os.Args) < 2 || len(os.Args) > 3 {
		usage()
	}
	baseDir = os.Args[1]
	if len(os.Args) == 3 {
		category = os.Args[2]
	}
	db := new(fortunes.FortunesDB)
	db.BaseDir = baseDir
	if category == "" {
		fortune, category, err = fortunes.GetFortune(db)
	} else {
		fortune, err = fortunes.GetFortuneByCategory(db, category)
	}
	if err != nil {
		fmt.Printf("Error occured: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n[%s]\n", fortune, category)
}
