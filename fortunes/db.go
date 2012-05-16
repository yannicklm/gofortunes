package fortunes

import (
	"io/ioutil"
)

type FortunesDB struct {
	BaseDir string
}

func (db *FortunesDB) GetFortune(category string) string {
	var toRead = db.BaseDir + "/" + category
	contents, err := ioutil.ReadFile(toRead)
	if err != nil {
		return ""
	}
	return string(contents)
}

func (db *FortunesDB) AddFortune(text string, category string) {
	var toWrite = db.BaseDir + "/" + category
	ioutil.WriteFile(toWrite, []byte(text), 0644)
}

func (db *FortunesDB) GetCategories() []string {
	return []string{}
}
