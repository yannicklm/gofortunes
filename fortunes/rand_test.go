package fortunes

import (
	"testing"
)

func Benchmark_RandChoice3(bench *testing.B) {
	var a, b, c int
	choices := []string{"a", "b", "c"}
	for i := 0; i < 3*bench.N; i++ {
		picked, err := RandChoice(choices)
		if err != nil {
			bench.Error(err)
		}
		switch picked {
		case "a":
			a++
		case "b":
			b++
		case "c":
			c++
		default:
			bench.Error("weird picked: ", picked)
		}
	}
	bench.Log("a ", a, "b ", b, "c", c)
}

func Benchmark_RandChoice2(bench *testing.B) {
	var a, b int
	choices := []string{"a", "b"}
	for i := 0; i < 2*bench.N; i++ {
		picked, err := RandChoice(choices)
		if err != nil {
			bench.Error(err)
		}
		switch picked {
		case "a":
			a++
		case "b":
			b++
		default:
			bench.Error("weird picked: ", picked)
		}
	}
	bench.Log("a ", a, "b ", b)
}
