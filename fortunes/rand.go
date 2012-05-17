package fortunes

import (
	"errors"
	"math/rand"
)

func RandChoice(choices []string) (string, error) {
	if len(choices) == 0 {
		return "", errors.New("empty list of choices")
	}
	if len(choices) == 1 {
		return choices[0], nil
	}
	index := rand.Intn(len(choices))
	return choices[index], nil
}
