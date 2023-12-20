package utils

import (
	"log"
	"os"
	"testing"

	"github.com/test-go/testify/assert"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TestCases(t *testing.T) {
	word := "dummy"
	titleWord := cases.Title(language.AmericanEnglish).String(word)
	assert.Equal(t, "Dummy", titleWord)
}

func TestAddToDictionary(t *testing.T) {
	word := "1234"
	dictionary := make(map[string]bool)
	AddToDictionary(dictionary, word)
	_, exists := dictionary[word]
	assert.Equal(t, true, exists)
}

func TestAddToFile(t *testing.T) {
	file, err := os.OpenFile("tmp", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	word := "dummy123"
	n, err := AddToFile(file, word)
	assert.Nil(t, err)
	assert.Equal(t, len(word)+1, n)

	os.Remove("tmp")
}
