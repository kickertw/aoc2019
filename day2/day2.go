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

func findAnswer(inputs []int, keyOutput int) int {
	i := 0
	for i < len(inputs) {
		op := inputs[i]
		if op == 99 {
			break
		} else {
			operand1 := inputs[inputs[i+1]]
			operand2 := inputs[inputs[i+2]]
			outputIndex := inputs[i+3]

			if op == 1 {
				inputs[outputIndex] = operand1 + operand2
			} else if op == 2 {
				inputs[outputIndex] = operand1 * operand2
			}
		}

		i += 4
	}

	// Part 1 answer
	if inputs[1] == 12 && inputs[2] == 2 {
		fmt.Printf("P1 Answer = %v\n", inputs[0])
		return 1
	}

	// Part 2 answer
	if inputs[0] == keyOutput {
		fmt.Printf("P2 Answer = 100 * noun[%v] + verb[%v] = %v \n", inputs[1], inputs[2], (100*inputs[1] + inputs[2]))
		return 2
	}

	return 0
}

func main() {

	// Printing out file contents
	inputfile := "day2input.txt"
	keyOutput := 19690720

	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	foundAnswers := 0
	for noun := 0; noun < 100 && foundAnswers < 3; noun++ {
		for verb := 0; verb < 100 && foundAnswers < 3; verb++ {
			inputs := stringToIntArr(input)
			inputs[1] = noun
			inputs[2] = verb

			retVal := findAnswer(inputs, keyOutput)
			foundAnswers += retVal
		}
	}
}
