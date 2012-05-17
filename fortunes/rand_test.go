package fortunes

import (
	"testing"
)

func Test_RandChoice(t *testing.T) {
	choices := []string{"a", "b"}
	picked, err := RandChoice(choices)
	if err != nil {
		t.Error(err)
	}
	CheckIsIn(t, picked, choices)
}
