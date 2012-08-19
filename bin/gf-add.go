package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"net/http"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

func die(err string) {
	fmt.Fprintf(os.Stderr, err + "\n")
	os.Exit(1)
}

func usage() {
	fmt.Printf("Usage: %s URL CATEGORY\n", os.Args[0])
	os.Exit(2)
}

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


func main() {
	if len(os.Args) != 3 {
		usage()
	}
	srvUrl  := os.Args[1]
	category := os.Args[2]

	text, err := readFortune()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resp, err := http.PostForm(srvUrl+ "/add",
		url.Values{
			"text" : {text},
			"category" : {category},
		})

	if err != nil {
		die(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		die(err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		die(resp.Status + "\n" + string(body))
	}

	fmt.Printf("Added new quote:\n%s\n[%s]\n", text, category)
}
