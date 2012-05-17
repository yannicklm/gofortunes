package main

import (
	"bytes"
	"errors"
	"gofortunes/fortunes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

func readFortune() (string, error) {
	tmpdir, err := ioutil.TempDir("", "gf-add")
	tmpfile := path.Join(tmpdir, "fortune.txt")
	defer os.RemoveAll(tmpdir)
	if err != nil {
		return "", nil
	}
	cmd := exec.Command("vim", tmpfile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(tmpfile)
	if err != nil {
		return "", errors.New("no file was written")
	}
	size := fi.Size()
	f, err := os.OpenFile(tmpfile, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		return "", err
	}
	data := make([]byte, size)
	n, err := f.Read(data)
	if err != nil {
		return "", err
	}
	if int64(n) < size {
		return "", io.ErrShortBuffer
	}
	// Remove empty new lines:
	data = bytes.TrimRight(data, "\n")
	return string(data), nil
}

func usage() {
	fmt.Println("Usage: %s DB_PATH CATEGORY", os.Args[0])
	os.Exit(2)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	baseDir  := os.Args[1]
	category := os.Args[2]

	db := new(fortunes.FortunesDB)
	db.BaseDir = baseDir

	text, err := readFortune()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fortunes.AddFortune(db, text, category)
}
