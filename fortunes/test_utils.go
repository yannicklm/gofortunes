package fortunes

import (
	"testing"
)

func CompareSlice(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CheckIsIn(t *testing.T, needle string, haystack []string) {
	for _, x := range haystack {
		if x == needle {
			return
		}
	}
	t.Errorf("%s not in %s", needle, haystack)
}
