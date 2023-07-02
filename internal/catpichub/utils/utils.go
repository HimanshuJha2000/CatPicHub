package utils

import (
	"math/rand"
	"time"
	"unicode"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateUniqueID() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func IsSpecialCharacter(value rune) bool {
	return unicode.IsPunct(value) || unicode.IsSymbol(value)
}

func IsAlphabet(value rune) bool {
	return unicode.IsLetter(value)
}
