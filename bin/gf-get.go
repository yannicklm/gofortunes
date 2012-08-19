package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"os"
)

func usage() {
	fmt.Printf("Usage: %s URL [CATEGORY]\n", os.Args[0])
	os.Exit(2)
}

func die(err string) {
	fmt.Fprintf(os.Stderr, err + "\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		usage()
	}
	srvUrl := os.Args[1]
	var category string = ""
	if len(os.Args) == 3 {
		category = os.Args[2]
	}

	resp, err := http.PostForm(srvUrl+ "/get",
		url.Values{
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

	fmt.Print(string(body))
}
