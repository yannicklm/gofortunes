package gofortunes

import (
	"net/http"
)

type Fortune struct {
	Text     string
	Category string
}

type Category struct {
	Name         string
}

func init() {
	http.HandleFunc("/", Root)
	http.HandleFunc("/getForm", GetForm)
	http.HandleFunc("/get", Get)
	http.HandleFunc("/addForm", AddForm)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/restoreForm", RestoreForm)
	http.HandleFunc("/restore", Restore)
	http.HandleFunc("/restoreTask", RestoreTask)
}
