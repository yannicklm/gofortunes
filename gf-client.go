package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"net/http"
	"net/url"
)

func usage() {
	fmt.Println("Usage: %s URL [CATEGORY]", os.Args[0])
	os.Exit(2)
}

func die(err string) {
	fmt.Fprintf(os.Stderr, err + "\n")
	os.Exit(1)
}

func main() {
	var srvUrl, category string
	var resp *http.Response
	var err error

	if len(os.Args) < 2 || len(os.Args) > 3 {
		usage()
	}

	srvUrl = os.Args[1]
	if len(os.Args) == 3 {
		category = os.Args[2]
	}

	if category == "" {
		resp, err = http.Get(srvUrl + "/fortune")
	} else {
		resp, err = http.PostForm(srvUrl +"/fortune",
			url.Values{"category" : {}})
	}

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

	fmt.Print(string(body))
}
