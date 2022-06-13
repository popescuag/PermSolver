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
		utils.AddToDictionaryFile(dictionary, os.Args[2])
		return
	}

	letters := os.Args[1]

	nodes, results := solver.SolvePuzzle(letters, dictionary)

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
	//print results
	for i := from; i <= to; i++ {
		fmt.Printf("%d %s \n", i, results[i])
	}
	fmt.Printf("Total nodes: %d. Duration: %d ms", nodes, time.Since(start)/time.Millisecond)
}
