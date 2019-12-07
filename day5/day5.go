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

func findAnswer(inputs []int, userInput int) int {
	i := 0
	for i < len(inputs) {
		op := inputs[i]
		switch op {
		case 99:
			break
		case 3, 4:
			// 3 = input, 4 = output
			if op == 3 {
				inputs[inputs[i+1]] = userInput
			} else {
				fmt.Println(inputs[inputs[i+1]])
			}
			i += 2
		case 1, 2:
			operand1 := inputs[inputs[i+1]]
			operand2 := inputs[inputs[i+2]]
			outputIndex := inputs[i+3]

			if op == 1 {
				inputs[outputIndex] = operand1 + operand2
			} else if op == 2 {
				inputs[outputIndex] = operand1 * operand2
			}

			i += 4
		default:
			// parameter mode
		}
	}

	// Part 1 answer
	if inputs[1] == 12 && inputs[2] == 2 {
		fmt.Printf("P1 Answer = %v\n", inputs[0])
		return 1
	}

	return 0
}

func main() {

	// Printing out file contents
	inputfile := "day5input.txt"
	initialInput := 1

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

			retVal := findAnswer(inputs)
			foundAnswers += retVal
		}
	}
}
