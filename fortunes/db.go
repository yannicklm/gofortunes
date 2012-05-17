package fortunes

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"time"
	"path"
	"math/rand"
	"strings"
)

type FortunesDB struct {
	baseDir string
}


func NewDB(baseDir string) *FortunesDB {
	res := new(FortunesDB)
	res.baseDir = baseDir
	now := time.Now()
	rand.Seed(int64(now.Nanosecond()))
	return res
}


func (db *FortunesDB) GetFortune() (fortune string, category string, err error) {
	categories, err := db.GetCategories()
	if err != nil {
		return "", "", err
	}
	cat, err := RandChoice(categories)
	if err != nil {
		return "", "", err
	}
	f, err := db.GetFortuneByCategory(cat)
	if err != nil {
		return "", "", err
	}
	return f, cat, nil
}

func (db *FortunesDB) GetFortuneByCategory(category string) (string, error) {
	toRead := path.Join(db.baseDir, category)
	fi, err := os.Stat(toRead)
	if err != nil {
		return "", errors.New("No such category: " + category)
	}
	size := fi.Size()

	f, err := os.OpenFile(toRead, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		return "", err
	}
	data := make([]byte, size)
	nn, err := f.Read(data)
	n := int64(nn)
	if err != nil {
		return "", err
	}
	if n < size {
		return "", io.ErrShortBuffer
	}

	if len(data) < 2 {
		return "", errors.New("No fortune in this category: " + category)
	}

	// Remove last '%\n'
	if data[len(data)-2] == '%' && data[len(data)-1] == '\n' {
		data = data[:len(data)-2]
	}

	text := string(data)
	fortunes := strings.Split(text, "%\n")
	fortune, err := RandChoice(fortunes)
	if err != nil {
		return "", err
	}
	fortune = strings.TrimRight(fortune, "\n")
	if fortune == "" {
		return "", errors.New("fortune was empty (in category: " + category + ")")
	}
	return fortune, nil
}

func (db *FortunesDB) AddFortune(text string, category string) error {
	toWrite := path.Join(db.baseDir, category)
	f, err := os.OpenFile(toWrite, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	data := []byte(text)
	if len(data) == 0 {
		return errors.New("Empty string")
	}
	if data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}
	data = append(data, '%', '\n')
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	f.Close()
	return err
}

func (db *FortunesDB) GetCategories() ([]string, error) {
	res := make([]string, 0)
	list, err := ioutil.ReadDir(db.baseDir)
	if err != nil {
		return []string{}, err
	}
	for _, fileInfo := range list {
		if fileInfo.IsDir() {
			continue
		}
		res = append(res, fileInfo.Name())
	}
	return res, nil
}