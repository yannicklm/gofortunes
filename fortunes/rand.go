package fortunes

import (
	"crypto/rand"
	"errors"
	"io"
)

// Rand only returns bit Ints
func randUInt64(max uint64) (uint64, error) {
	// Read 8 byts from the rand reader and convert it
	// to a uint64
	bytes := make([]byte, 8)
	nn, err := rand.Read(bytes)
	if nn < 8 {
		return 0, io.ErrShortBuffer
	}
	if err != nil {
		return 0, err
	}
	var res uint64 = 0
	for i, b := range bytes {
		res |= uint64(int(b) << 8 * i)
	}
	res = res % max
	return res, nil
}

func RandChoice(choices []string) (string, error) {
	if len(choices) == 0 {
		return "", errors.New("empty list of choices")
	}
	if len(choices) == 1 {
		return choices[0], nil
	}
	index, err := randUInt64(uint64(len(choices)))
	if err != nil {
		return "", err
	}
	return choices[index], nil
}
