package gofortunes

import (
	"appengine"
	"appengine/datastore"
	"net/http"
	"fmt"
)

func addNewCategory(c appengine.Context, name string) error {
	// Add the category to the list if necessary,
	// and update the fortune count
	catKey := datastore.NewKey(c, "Category", name, 0, nil)
	category := Category{
		Name:         name,
	}
	_, err := datastore.Put(c, catKey, &category)
	return err
}

func Add(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	categoryName := r.FormValue("category")
	fortune := Fortune{
		Text:     text,
		Category: categoryName,
	}
	c := appengine.NewContext(r)
	// Add the full fortune to the list,
	fortuneKey := datastore.NewIncompleteKey(c, "Fortune", nil)
	_, err := datastore.Put(c, fortuneKey, &fortune)
	if err != nil {
		http.Error(w, "Could no add fortune: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = addNewCategory(c, categoryName)
	if err != nil {
		http.Error(w, "Could not add new category: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func AddForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, addFormHTML)
}

const addFormHTML = `
<html>
 <body>
   <form action="/add" method="post">
	<div><textarea name="category" rows="1" cols="60"></textarea></div>
	<div><textarea name="text" rows="15" cols="80"></textarea></div>
	<div><input type="submit" value="Add"></div>
  </form>
 </body>
</html>
`
