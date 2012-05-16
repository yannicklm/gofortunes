package fortunes

import (
	"io/ioutil"
	"os"
	"testing"
)

func setupDB() *FortunesDB {
	tmp, _ := ioutil.TempDir("", "gf-db-test")
	db := new(FortunesDB)
	db.BaseDir = tmp
	return db
}

func Test_EmptyDB(t *testing.T) {
	var db = setupDB()
	defer os.RemoveAll(db.BaseDir)
	var s = db.GetFortune("jokes")
	if s != "" {
		t.Errorf("Fortune returned by new db not empty\nGot %s", s)
	}

	var categories = db.GetCategories()
	if len(categories) != 0 {
		t.Errorf("New db should be empty, but contains the following categories %s", categories)
	}
}

func Test_AddOneQuote(t *testing.T) {
	var db = setupDB()
	defer os.RemoveAll(db.BaseDir)
	var quote = "This is a quote"
	db.AddFortune(quote, "quote")
	var got = db.GetFortune("quote")
	if got != quote {
		t.Errorf("got '%s', expecting '%s'", got, quote)
	}
}

func test_AddJokeAndQuote(t *testing.T) {
	var db = setupDB()
	defer os.RemoveAll(db.BaseDir)
	var quote = "This is a quote"
	db.AddFortune(quote, "quote")
	var joke = "This is a joke"
	db.AddFortune(joke, "joke")
	var got = db.GetFortune("quote")
	if got != quote {
		t.Errorf("got '%s', expecting '%s'", got, quote)
	}
	got = db.GetFortune("joke")
	if got != joke {
		t.Errorf("got '%s', expecting '%s'", got, joke)
	}
}
