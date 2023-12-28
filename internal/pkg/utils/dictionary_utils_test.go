package utils

import (
	"log"
	"os"
	"strings"
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

func TestIsWordInDictionary1(t *testing.T) {
	word := "dummy"
	dictionary := make(map[string]bool)
	dictionary[word] = false
	assert.True(t, IsWordInDictionary(dictionary, word))
}

func TestIsWordInDictionary2(t *testing.T) {
	word := "DUMMY"
	dictionary := make(map[string]bool)
	dictionary[word] = false
	assert.True(t, IsWordInDictionary(dictionary, strings.ToLower(word)))
}

func TestIsWordInDictionary3(t *testing.T) {
	word := "Dummy"
	dictionary := make(map[string]bool)
	dictionary[word] = false
	assert.True(t, IsWordInDictionary(dictionary, strings.ToLower(word)))
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
