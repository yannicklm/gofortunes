package gofortunes

import (
	"appengine"
	"appengine/blobstore"
	"html/template"
	"io"
	"fmt"
	"net/http"
)

func RestoreForm(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	uploadURL, err := blobstore.UploadURL(c, "/restore", nil)
	if err != nil {
		c.Errorf("%s", err)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = restoreFormTemplate.Execute(w, uploadURL)
	if err != nil {
		c.Errorf("%v", err)
	}
}

func Restore(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
		c.Errorf("%v", err)
		return
	}
	file := blobs["file"]
	if len(file) == 0 {
		c.Errorf("no file uploaded")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	blobKey := file[0].BlobKey
	blobReader := blobstore.NewReader(c, blobKey)
	restoreFromReader(blobReader)
	http.Redirect(w, r, "/", http.StatusFound)
}

func restoreFromReader(r io.Reader) {
	data := make([]byte, 100)
	r.Read(data)
	fmt.Printf("read 100 bytes: %s", string(data))
}

var restoreFormTemplate = template.Must(template.New("root").Parse(restoreFormTemplateHTML))

const restoreFormTemplateHTML = `
<html><body>
<form action="{{.}}" method="POST" enctype="multipart/form-data">
Upload File: <input type="file" name="file"><br>
<input type="submit" name="submit" value="Submit">
</form></body></html>
`
