package utils

import (
	"math/rand"
	"time"
	"unicode"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateUniqueID generates unique 6 character ID for primary key in cat_pics table
func GenerateUniqueID() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// To check if the received character is a special character
func IsSpecialCharacter(value rune) bool {
	return unicode.IsPunct(value) || unicode.IsSymbol(value)
}

// To check if the received character is an alphabet
func IsAlphabet(value rune) bool {
	return unicode.IsLetter(value)
}
