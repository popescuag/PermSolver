package solver

import (
	"sort"

	"github.com/popescuag/PermSolver/internal/pkg/utils"
)

type treeNode struct {
	pathToRoot     string
	remainingChars string
	children       []treeNode
}

func SolvePuzzle(letters string, dictionary map[string]bool, minLetterCount int, maxLetterCount int) (int, map[int][]string) {
	var results map[int][]string = make(map[int][]string)
	//Create root element
	root := treeNode{
		pathToRoot:     "",
		remainingChars: letters,
	}

	nodes := buildTree(&root, &results, dictionary)
	for i := minLetterCount; i <= maxLetterCount; i++ {
		sort.Strings(results[i])
	}
	return nodes, results
}

func addChild(currentNode *treeNode, child treeNode, nodes *int) {
	currentNode.children = append(currentNode.children, child)
	*nodes++
}

/*
Expands a tree node by creating a child for each of the remaining chars.
*/
func expandNode(currentNode *treeNode, results *map[int][]string, dictionary map[string]bool, nodes *int) {
	if len(currentNode.remainingChars) == 0 {
		return
	}
	for i := 0; i < len(currentNode.remainingChars); i++ {
		currChar := currentNode.remainingChars[i]
		newChild := treeNode{
			pathToRoot:     currentNode.pathToRoot + string(currChar),
			remainingChars: string(currentNode.remainingChars[0:i]) + string(currentNode.remainingChars[(i+1):]),
		}
		addChild(currentNode, newChild, nodes)
		wordLen := len(newChild.pathToRoot)
		if wordLen >= 3 {
			//check if the word is in the dictionary
			if utils.IsWordInDictionary(dictionary, newChild.pathToRoot) {
				var arr = (*results)[wordLen]
				//add to results map if not already there
				if !utils.Contains(arr, newChild.pathToRoot) {
					arr = append(arr, newChild.pathToRoot)
				}
				(*results)[wordLen] = arr
			}
		}
	}
}

func buildTree(currNode *treeNode, results *map[int][]string, dictionary map[string]bool) int {
	nodes := 0
	expandNode(currNode, results, dictionary, &nodes)
	for _, child := range currNode.children {
		buildTree(&child, results, dictionary)
	}
	return nodes
}
