package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"net/http"
	"os"
	"path"
	"strings"
)

func die(err string) {
	fmt.Fprintf(os.Stderr, err+"\n")
	os.Exit(1)
}

func usage() {
	fmt.Printf("Usage: %s DB_PATH URL\n", os.Args[0])
	os.Exit(2)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	dbPath := os.Args[1]
	srvUrl := os.Args[2]

	categories := getCategories(dbPath)

	for i := range categories {
		category := categories[i]
		fmt.Print("::  " + category + "\n")
		fortunes := getFortunes(dbPath, category)
		for j := range fortunes {
			fortune := fortunes[j]
			percent := float64(j) / float64(len(fortunes)) * 100.0
			fmt.Printf("%2.0f%%\r", percent)
			addFortune(srvUrl, category, fortune)
		}
		fmt.Print("\n")
	}
}

func getCategories(dbPath string) []string {
	res := make([]string, 0)
	list, err := ioutil.ReadDir(dbPath)
	if err != nil {
		return res
	}
	for _, fileInfo := range list {
		if fileInfo.IsDir() {
			continue
		}
		res = append(res, fileInfo.Name())
	}
	return res
}

func getFortunes(dbPath string, category string) []string {
	res := make([]string, 0)
	toRead := path.Join(dbPath, category)
	fi, err := os.Stat(toRead)
	if err != nil {
		return res
	}
	size := fi.Size()
	f, err := os.OpenFile(toRead, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		return res
	}
	data := make([]byte, size)
	nn, err := f.Read(data)
	n := int64(nn)
	if err != nil {
		return res
	}
	if n < size {
		return res
	}

	if len(data) < 2 {
		return res
	}

	// Remove last '%\n'
	if data[len(data)-2] == '%' && data[len(data)-1] == '\n' {
		data = data[:len(data)-2]
	}

	text := string(data)
	return strings.Split(text, "%\n")
}

func addFortune(srvUrl string, category string, fortune string) {
	if len(fortune) > 400 {
		return
	}
	resp, err := http.PostForm(srvUrl+"/add", url.Values{
		"text":     {fortune},
		"category": {category},
	})
	if err != nil {
		die(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		die(err.Error())
	}
	if resp.StatusCode != 302 {
		die(resp.Status + "\n" + string(body))
	}
}
