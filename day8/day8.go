package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stringToIntArr(input string) []int {
	s := strings.Split(input, ",")

	var retVal = []int{}

	for _, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		retVal = append(retVal, j)
	}

	return retVal
}

func main() {
	inputfile := "input.txt"

	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	inputs := stringToIntArr(input)

	layerWidth, layerHeight := 25, 6
	maxLen := layerWidth * layerHeight
	layers := [][]int{}
	layer := make([]int, maxLen)

	counter := 0
	for _, val := range inputs {
		layer[counter] = val
		counter++
		if counter == maxLen {
			layers = append(layers, layer)
			counter = 0
		}
	}

}
