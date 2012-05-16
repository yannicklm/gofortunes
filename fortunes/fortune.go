package fortunes

func GetFortune(db FortunesManager, category string) string {
	return db.GetFortune(category)
}

func AddFortune(db FortunesManager, text string, category string) {
	db.AddFortune(text, category)
}


