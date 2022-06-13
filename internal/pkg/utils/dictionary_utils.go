package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
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

func AddToDictionaryFile(dictionary map[string]bool, word string) {
	_, found1 := dictionary[word]
	_, found2 := dictionary[strings.Title(word)]
	_, found3 := dictionary[strings.ToUpper(word)]

	if !found1 && !found2 && !found3 {
		file, err := os.OpenFile("words.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		_, err = file.Write([]byte(word))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		log.Printf("Added %v to dictionary", word)
	} else {
		log.Printf("Word %v already in dictionary", word)
	}
}
