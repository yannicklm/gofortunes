package fortunes

func GetFortune(db FortunesManager) (fortune string, category string, err error) {
	return db.GetFortune()
}

func GetFortuneByCategory(db FortunesManager, category string) (string, error) {
	return db.GetFortuneByCategory(category)
}

func AddFortune(db FortunesManager, text string, category string) error {
	return db.AddFortune(text, category)
}


