package main

import (
	"fmt"
	"net/http"
	"gofortunes/fortunes"
)

type FortuneHandler struct {
	db *fortunes.FortunesDB
}

func NewFH(path string) *FortuneHandler {
	res := new(FortuneHandler)
	db := fortunes.NewDB(path)
	res.db = db
	return res
}

func (fh *FortuneHandler) ServeHTTP(w http.ResponseWriter,
									 r *http.Request) {
	var err error;
	var category, fortune string;
	var ok bool;
	err = r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "error: %s", err)
		return
	}
	form := r.Form
	categories, ok := form["category"]
	if ok {
		// Not sure why there is a list in the form key ...
		category = categories[0]
		fortune, err = fh.db.GetFortuneByCategory(category)
	} else {
		fortune, category, err = fh.db.GetFortune()
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error when getting fortune: %s", err)
		return
	}
	fmt.Fprintf(w, "%s\n[%s]\n", fortune, category)
}


func main() {
	fh := NewFH("/usr/share/fortune")
	http.Handle("/fortune", fh)
	http.ListenAndServe("localhost:3000", nil)
}

