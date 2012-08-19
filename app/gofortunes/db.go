package gofortunes

import (
	"appengine"
	"appengine/datastore"
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func LoadDB(c appengine.Context, reader io.Reader) {
	gzr, err := gzip.NewReader(reader)
	if err != nil {
		c.Errorf("error : %s\n", err.Error())
		return
	}
	defer gzr.Close()
	tr := tar.NewReader(gzr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if hdr.Typeflag != tar.TypeReg {
			continue
		}
		if err != nil {
			c.Errorf("error when reading backup: %v\n", err)
			continue
		}
		categoryName := hdr.Name[strings.Index(hdr.Name, "/") + 1:]
		fmt.Printf("reading %s\n", categoryName)
		catKey := datastore.NewKey(c, "Category", categoryName, 0, nil)
		category := Category{
			Name: categoryName,
		}
		_, err = datastore.Put(c, catKey, &category)
		if err != nil {
			c.Errorf("error when adding category: %v\n", err)
			continue
		}
		size := hdr.Size
		data := make([]byte, size)
		_, err = tr.Read(data)
		if len(data) < 2 {
			c.Errorf("Short read for category: %v\n", err)
			continue
		}
		if err != nil {
			c.Errorf("Error when reading category: %v\n", err)
			continue
		}
		// Remove last '%\n'
		if data[len(data)-2] == '%' && data[len(data)-1] == '\n' {
			data = data[:len(data)-2]
		}
		text := string(data)
		fortunes := strings.Split(text, "%\n")
		for i := range fortunes {
			fortune := Fortune{
				Text:     fortunes[i],
				Category: categoryName,
			}
			fortuneKey := datastore.NewIncompleteKey(c, "Fortune", nil)
			_, err := datastore.Put(c, fortuneKey, &fortune)
			if err != nil {
				c.Errorf("error when adding new fortune: %v\n", err)
				continue
			}
		}
	}
}
