package solver

import (
	"strings"

	"github.com/popescuag/PermSolver/internal/pkg/utils"
)

type treeNode struct {
	pathToRoot     string
	remainingChars string
	children       []treeNode
}

var (
	results map[int][]string = make(map[int][]string)
	nodes                    = 0
)

func SolvePuzzle(letters string, dictionary map[string]bool) (int, map[int][]string) {
	//Create root element
	root := treeNode{
		pathToRoot:     "",
		remainingChars: letters,
	}

	buildTree(&root, &results, dictionary)
	return nodes, results
}

func addChildren(currentNode *treeNode, child treeNode) {
	currentNode.children = append(currentNode.children, child)
	nodes++
}

/*
Expands a tree node by creating a child for each of the remaining chars.
*/
func expandNode(currentNode *treeNode, results *map[int][]string, dictionary map[string]bool) {
	if len(currentNode.remainingChars) == 0 {
		return
	}
	for i := 0; i < len(currentNode.remainingChars); i++ {
		currChar := currentNode.remainingChars[i]
		newChild := treeNode{
			pathToRoot:     currentNode.pathToRoot + string(currChar),
			remainingChars: string(currentNode.remainingChars[0:i]) + string(currentNode.remainingChars[(i+1):]),
		}
		addChildren(currentNode, newChild)
		wordLen := len(newChild.pathToRoot)
		if wordLen >= 3 {
			//check if the word is in the dictionary
			_, found1 := dictionary[newChild.pathToRoot]
			_, found2 := dictionary[strings.Title(newChild.pathToRoot)]
			_, found3 := dictionary[strings.ToUpper(newChild.pathToRoot)]
			if found1 || found2 || found3 {
				var arr = (*results)[wordLen]
				//add to results map if not already there
				utils.AddToSortedArray(&arr, newChild.pathToRoot)
				(*results)[wordLen] = arr
			}
		}
	}
}

func buildTree(currNode *treeNode, results *map[int][]string, dictionary map[string]bool) {
	expandNode(currNode, results, dictionary)
	for _, child := range currNode.children {
		buildTree(&child, results, dictionary)
	}
}
