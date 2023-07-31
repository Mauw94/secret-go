package utils

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz!@#$%^")

func MakePassphrase() []byte {

	b := make([]rune, 32)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return []byte(string(b))
}
