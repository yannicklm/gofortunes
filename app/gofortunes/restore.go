package gofortunes

import (
	"appengine"
	"appengine/blobstore"
	"appengine/taskqueue"
	"html/template"
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
	key := string(blobKey)
	t := taskqueue.NewPOSTTask("/restoreTask",
		map[string][]string{"blobKey": {key}})
	_, err = taskqueue.Add(c, t, "")
	if err != nil {
		c.Errorf("Could not add task. Error was: %v", err)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func RestoreTask(w http.ResponseWriter, r *http.Request) {
	blobKey := appengine.BlobKey(r.FormValue("blobKey"))
	c := appengine.NewContext(r)
	blobInfo, err := blobstore.Stat(c, blobKey)
	if err != nil {
		c.Errorf("%v", err)
		return
	}
	c.Infof("Restoring from %s", blobInfo.Filename)
	reader := blobstore.NewReader(c, blobKey)
	LoadDB(c, reader)
}

var restoreFormTemplate = template.Must(template.New("root").Parse(restoreFormTemplateHTML))

const restoreFormTemplateHTML = `
<html><body>
<form action="{{.}}" method="POST" enctype="multipart/form-data">
Upload File: <input type="file" name="file"><br>
<input type="submit" name="submit" value="Submit">
</form></body></html>
`
