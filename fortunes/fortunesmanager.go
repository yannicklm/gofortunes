package fortunes

type FortunesManager interface {
	GetFortune(category string) string
	AddFortune(text string, category string)
	GetCategories() []string
}
