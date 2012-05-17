package fortunes

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)


func setupDB() *FortunesDB {
	tmp, _ := ioutil.TempDir("", "gf-db-test")
	db := new(FortunesDB)
	db.BaseDir = tmp
	return db
}

func Test_EmptyDB(t *testing.T) {
	db := setupDB()
	defer os.RemoveAll(db.BaseDir)
	s, err := db.GetFortuneByCategory("jokes")
	if s != "" {
		t.Errorf("Fortune returned by new db not empty\nGot %s", s)
	}
	if err == nil {
		t.Errorf("Should have returned an error")
	}
	categories, err := db.GetCategories()
	if len(categories) != 0 {
		t.Errorf("New db should be empty, "+
			"but contains the following categories %s", categories)
	}
}

func Test_AddOneQuote(t *testing.T) {
	db := setupDB()
	defer os.RemoveAll(db.BaseDir)
	quote := "This is a quote"
	err := db.AddFortune(quote, "quote")
	if err != nil {
		t.Errorf("Error when adding quote: %s", err)
	}

	got, err := db.GetFortuneByCategory("quote")
	if err != nil {
		t.Error("Error when getting quote: %s", err)
	}
	if got != quote {
		t.Errorf("got '%s', expecting '%s'", got, quote)
	}
}

func Test_AddJokeAndQuote(t *testing.T) {
	db := setupDB()
	defer os.RemoveAll(db.BaseDir)
	quote := "This is a quote"
	db.AddFortune(quote, "quotes")
	joke := "This is a joke"
	db.AddFortune(joke, "jokes")
	got, _ := db.GetFortuneByCategory("quotes")
	if got != quote {
		t.Errorf("got '%s', expecting '%s'", got, quote)
	}
	got, _ = db.GetFortuneByCategory("jokes")
	if got != joke {
		t.Errorf("got '%s', expecting '%s'", got, joke)
	}
	expected_cat := []string{"jokes", "quotes"}
	actual_cat, err := db.GetCategories()
	if err != nil {
		t.Errorf("Error when getting categorines: %s", err)
	}
	ok := CompareSlice(expected_cat, actual_cat)
	if !ok {
		t.Errorf("got %s, expecting %s", actual_cat, expected_cat)
	}

	fortune, category, _ := db.GetFortune()

	CheckIsIn(t, fortune, []string{quote, joke})
	CheckIsIn(t, category, []string{"quotes", "jokes"})
}

func Test_AddQuotes(t *testing.T) {
	db := setupDB()
	defer os.RemoveAll(db.BaseDir)
	quotes := []string{"First quote", "Second quote"}
	for _, quote := range quotes {
		db.AddFortune(quote, "quotes")
	}
	returned, _ := db.GetFortuneByCategory("quotes")
	CheckIsIn(t, returned, quotes)
}

func Test_GetNoSuchCategory(t *testing.T) {
	db := setupDB()
	defer os.RemoveAll(db.BaseDir)
	_, err := db.GetFortuneByCategory("doesnotexists")
	if !strings.Contains(err.Error(), "No such category") {
		t.Errorf("Wrong error message:\n %s", err)
	}
}
