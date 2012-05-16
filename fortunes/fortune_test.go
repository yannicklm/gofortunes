package fortunes

import (
	"testing"
)

func Test_GetFortune(t *testing.T) {
	var s = GetFortune()
	if s == "" {
		t.Fail()
	}
}
