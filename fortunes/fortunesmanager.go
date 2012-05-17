package fortunes

type FortunesManager interface {
	GetFortune() (fortune string, category string, err error)
	GetFortuneByCategory(category string) (string, error)
	AddFortune(text string, category string) error
	GetCategories() ([]string, error)
}
