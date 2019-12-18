package main

import (
	"bufio"
	"fmt"
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

func createLayers(inputs []string, maxLen int) [][]int {
	layers := [][]int{}
	layer := make([]int, maxLen)

	counter := 0
	for _, val := range inputs {
		layer[counter], _ = strconv.Atoi(val)
		counter++
		if counter == maxLen {
			layerCopy := make([]int, maxLen)
			copy(layerCopy, layer)
			layers = append(layers, layerCopy)
			counter = 0
			layer = make([]int, maxLen)
		}
	}

	return layers
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
	inputs := strings.Split(input, "")

	layerWidth, layerHeight := 25, 6
	maxLen := layerWidth * layerHeight
	layers := createLayers(inputs, maxLen)

	fewestZeros := -1
	p1Answer := 0
	for i, layer := range layers {
		zeroCounter := make(map[int]int)
		for _, cellVal := range layer {
			if cellVal == 0 || cellVal == 1 || cellVal == 2 {
				zeroCounter[cellVal]++
			}
		}

		if fewestZeros == -1 || zeroCounter[0] < fewestZeros {
			p1Answer = zeroCounter[1] * zeroCounter[2]
			fewestZeros = zeroCounter[0]
			fmt.Printf("Layer[%v] now has fewest zeros [%v] and with [%v] ones and [%v] twos\n", i, zeroCounter[0], zeroCounter[1], zeroCounter[2])
		}
	}

	fmt.Printf("P1 = %v\n", p1Answer)

	finalImage := make([]int, maxLen)
	for ii := 0; ii < maxLen; ii++ {
		for jj := 0; jj < len(layers); jj++ {
			if layers[jj][ii] < 2 {
				finalImage[ii] = layers[jj][ii]
				break
			}
		}
	}

	fmt.Println("")
	fmt.Println("Final Image:")
	for len(finalImage) > 0 {
		fmt.Println(finalImage[0:layerWidth])
		finalImage = finalImage[layerWidth:]
	}
}
