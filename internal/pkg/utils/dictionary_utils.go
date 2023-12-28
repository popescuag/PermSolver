package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ReadDictionaryFile() map[string]bool {
	file, err := os.Open("words.txt")
	var dictionary map[string]bool = make(map[string]bool)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dictionary[line] = false
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return dictionary
}

func IsWordInDictionary(dictionary map[string]bool, word string) bool {
	_, found1 := dictionary[word]
	_, found2 := dictionary[cases.Title(language.AmericanEnglish).String(word)]
	_, found3 := dictionary[strings.ToUpper(word)]

	return found1 || found2 || found3
}

func AddToDictionary(dictionary map[string]bool, word string) {
	dictionary[word] = false
}

func GetDictionaryFile() *os.File {
	file, err := os.OpenFile("words.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func AddToFile(file *os.File, word string) (int, error) {
	n, err := file.Write([]byte(fmt.Sprintf("\n%v", word)))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Printf("Added %v to file", word)
	return n, err
}
