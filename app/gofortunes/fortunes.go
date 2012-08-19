package gofortunes

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"html/template"
	"math/rand"
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
	http.HandleFunc("/", root)
	http.HandleFunc("/getForm", getForm)
	http.HandleFunc("/get", get)
	http.HandleFunc("/addForm", addForm)
	http.HandleFunc("/add", add)
}

const rootHTML = `
<html>
 <body>
   <h1> Welcome to GoFortunes </h1>
   <ul>
	<li> <a href="addForm"> Add a new fortune </a> </li>
	<li> <a href="getForm"> Get fortune </a> </li>
  </ul>
 </body>
</html>
`

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

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootHTML)
}

func get(w http.ResponseWriter, r *http.Request) {
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

func getForm(w http.ResponseWriter, r *http.Request) {
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

func add(w http.ResponseWriter, r *http.Request) {
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

func addForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, addFormHTML)
}
