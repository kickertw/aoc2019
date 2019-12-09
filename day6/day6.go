package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

//Node - representing an object that will orbit
type Node struct {
	id       string
	parentID string
	dist     int
	children []Node
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildMap(m []Node, rootNode *Node, distance int) []Node {
	indexesToRemove := make([]int, 0)

	// Find any nodes that have "rootNodeID" as the parent.
	// Make the node and add it to the child array of the "rootNode"
	for i, val := range m {
		if val.parentID == rootNode.id {
			newNode := Node{id: val.id, parentID: val.parentID, dist: distance, children: make([]Node, 0)}
			rootNode.children = append(rootNode.children, newNode)
			// fmt.Printf("Node [%v] now has [%v] children - just added id [%v]\n", rootNode.id, len(rootNode.children), newNode.id)
			indexesToRemove = append(indexesToRemove, i)
		}
	}

	// make the original array to search smaller
	sort.Sort(sort.Reverse(sort.IntSlice(indexesToRemove)))
	for _, idx := range indexesToRemove {
		if idx == len(m)-1 {
			m = m[idx+1:]
		} else {
			m = append(m[:idx], m[idx+1:]...)
		}
	}

	// Recursively call each child and keep building...
	if len(m) > 0 {
		for ci := range rootNode.children {
			m = buildMap(m, &rootNode.children[ci], distance+1)
		}
	}

	return m
}

func countOrbits(node Node) int {

	fmt.Printf("Node[%v] has dist [%v]\n", node.id, node.dist)
	count := node.dist

	if node.children == nil || len(node.children) == 0 {
		fmt.Printf("	No more kids - returning %v\n", node.dist)
		return node.dist
	}

	for _, child := range node.children {
		count += countOrbits(child)
	}

	fmt.Printf("	Total distance from Node [%v] and below is %v\n", node.id, count)
	return count
}

func main() {

	// Printing out file contents
	inputfile := "day6input.txt"

	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	allInputs := []Node{Node{id: "COM", parentID: ""}}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ")")
		tempNode := Node{id: input[1], parentID: input[0]}
		allInputs = append(allInputs, tempNode)
	}

	rootNode := Node{id: "COM", parentID: "", children: make([]Node, 0)}
	buildMap(allInputs, &rootNode, 1)

	p1Total := countOrbits(rootNode)
	fmt.Printf("Total orbits = %v", p1Total)
}
