package gofortunes

import (
	"appengine"
	"appengine/datastore"
	"html/template"
	"net/http"
	"math/rand"
	"fmt"
)

func getRandomCategory(c appengine.Context) (string, error) {
	categories := make([]Category, 0, 1000)
	q := datastore.NewQuery("Category").Limit(1000)
	_, err := q.GetAll(c, &categories)
	if err != nil {
		return "", err;
	}
	randIndex := rand.Intn(len(categories))
	return categories[randIndex].Name, nil
}

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	categoryName := r.FormValue("category")
	if categoryName == "" {
		randomCat, err := getRandomCategory(c)
		if err != nil {
			http.Error(w, "Could not get random category" + err.Error(), http.StatusInternalServerError)
			return
		}
		categoryName = randomCat
	}
	q := datastore.NewQuery("Fortune").Filter("Category =", categoryName)
	count, err := q.Count(c)
	if err != nil {
		mess := fmt.Sprint("Could not count fortunes matching: %s. Error was: %s", categoryName, err.Error())
		http.Error(w, mess, http.StatusInternalServerError)
		return
	}
	if count == 0 {
		mess := fmt.Sprintf("No fortune matching %s", categoryName)
		http.Error(w, mess, http.StatusInternalServerError)
		return
	}
	fortunes := make([]Fortune, 0, count)
	_, err = q.GetAll(c, &fortunes)
	if err != nil {
		mess := fmt.Sprint("Could not get fortunes matching: %s. Error was: %s", categoryName, err.Error())
		http.Error(w, mess, http.StatusInternalServerError)
		return
	}
	randIndex := rand.Intn(count)
	fortune := fortunes[randIndex]
	fmt.Fprintf(w, "%s\n[%s]\n", fortune.Text, fortune.Category)
}


func GetForm(w http.ResponseWriter, r *http.Request) {
	categories := make([]Category, 0, 1000)
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Category").Order("Name").Limit(1000)
	_, err := q.GetAll(c, &categories)
	if err != nil {
		http.Error(w, "Could not get categories: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = getFormTemplate.Execute(w, categories)
	if err != nil {
		http.Error(w, "Could not execute template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

const getFormTemplateHTML = `
<html>
 <body>
   <form action="/get" method="post">
	{{range . }}
	  <input type="radio" name="category" value="{{.Name}}" /> {{.Name}} <br />
	{{end}}
	  <input type="radio" name="category" value="" /> Any <br />
	<div><input type="submit" value="Submit!"></div>
  </form>
 </body>
</html>
`

var getFormTemplate = template.Must(template.New("getForm").Parse(getFormTemplateHTML))
