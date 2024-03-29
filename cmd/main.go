package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/popescuag/PermSolver/internal/pkg/solver"
	"github.com/popescuag/PermSolver/internal/pkg/utils"
)

func main() {
	start := time.Now()

	if len(os.Args) == 1 {
		log.Fatal("At least one argument required, letters. Second and third arguments are optional")
		return
	}

	dictionary := utils.ReadDictionaryFile()

	if os.Args[1] == "add" {
		word := os.Args[2]
		if !utils.IsWordInDictionary(dictionary, word) {
			utils.AddToDictionary(dictionary, word)
			file := utils.GetDictionaryFile()
			utils.AddToFile(file, word)
		} else {
			log.Printf("Word %v already in dictionary", word)
		}
		return
	}

	letters := os.Args[1]

	from := 3
	to := len(letters)
	if len(os.Args) >= 3 {
		from, _ = strconv.Atoi(os.Args[2])
	}
	if len(os.Args) == 4 {
		to, _ = strconv.Atoi(os.Args[3])
	}
	if to > len(letters) {
		to = len(letters)
	}
	nodes, results := solver.SolvePuzzle(letters, dictionary, from, to)
	//print results
	for i := from; i <= to; i++ {
		fmt.Printf("%d %s \n", i, results[i])
	}
	fmt.Printf("Total nodes: %d. Duration: %d ms", nodes, time.Since(start)/time.Millisecond)
}
