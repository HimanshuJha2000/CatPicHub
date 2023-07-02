package utils

import (
	"github.com/DevtronLabs/CatPicHub/internal/catpichub/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateUniqueID(t *testing.T) {
	val := utils.GenerateUniqueID()
	assert.Equal(t, len(val), 6)
}

func TestIsSpecialCharacter(t *testing.T) {
	// Test with special characters
	specialChars := "!@#$%^&*()-_=+[]{};:'\"\\|,<.>/?`~"
	for _, char := range specialChars {
		if !utils.IsSpecialCharacter(char) {
			t.Errorf("Expected IsSpecialCharacter to return true for special character %c", char)
		}
	}

	// Test with non-special characters
	nonSpecialChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, char := range nonSpecialChars {
		if utils.IsSpecialCharacter(char) {
			t.Errorf("Expected IsSpecialCharacter to return false for non-special character %c", char)
		}
	}
}

func TestIsAlphabet(t *testing.T) {
	// Test with alphabet characters
	alphabetChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, char := range alphabetChars {
		if !utils.IsAlphabet(char) {
			t.Errorf("Expected IsAlphabet to return true for alphabet character %c", char)
		}
	}

	// Test with non-alphabet characters
	nonAlphabetChars := "!@#$%^&*()-_=+[]{};:'\"\\|,<.>/?`~0123456789"
	for _, char := range nonAlphabetChars {
		if utils.IsAlphabet(char) {
			t.Errorf("Expected IsAlphabet to return false for non-alphabet character %c", char)
		}
	}
}
