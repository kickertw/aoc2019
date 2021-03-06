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

func runProgram(inputs []int, userInput int) {
	i := 0
	for i < len(inputs) {
		op := inputs[i]
		if op == 99 {
			return
		}
		// fmt.Printf("@ Index[%v] - Opcode = %v\n", i, op)

		opCode := op % 10
		p1Mode := (op / 100) % 10
		p2Mode := (op / 1000) % 10

		p1 := 0
		p2 := 0
		if opCode < 3 {
			p1 = inputs[i+1]
			if p1Mode == 0 {
				p1 = inputs[p1]
			}

			p2 = inputs[i+2]
			if p2Mode == 0 {
				p2 = inputs[p2]
			}
		}

		outputIndex := inputs[i+3]

		switch opCode {
		case 1:
			inputs[outputIndex] = p1 + p2
			// fmt.Printf("%v + %v = %v (Stored @ %v)\n", p1, p2, inputs[outputIndex], outputIndex)
		case 2:
			inputs[outputIndex] = p1 * p2
			// fmt.Printf("%v * %v = %v (Stored @ %v)\n", p1, p2, inputs[outputIndex], outputIndex)
		case 3:
			inputs[inputs[i+1]] = userInput
			// fmt.Printf("input %v (Stored @ %v)\n", userInput, inputs[i+1])
		case 4:
			fmt.Println(inputs[inputs[i+1]])
			// fmt.Printf("output %v (Stored @ %v)\n", inputs[inputs[i+1]], inputs[i+1])
		}

		if opCode == 1 || opCode == 2 {
			i += 4
		} else {
			i += 2
		}
	}
}

func runProgramP2(inputs []int, userInput int) {
	i := 0
	for i < len(inputs) {
		op := inputs[i]
		if op == 99 {
			return
		}
		// fmt.Printf("@ Index[%v] - Opcode = %v\n", i, op)

		opCode := op % 10
		p1Mode := (op / 100) % 10
		p2Mode := (op / 1000) % 10

		p1 := inputs[i+1]
		p2 := inputs[i+2]
		outputIndex := 0
		if opCode == 1 || opCode == 2 || opCode == 7 || opCode == 8 {
			if p1Mode == 0 {
				p1 = inputs[p1]
			}

			if p2Mode == 0 {
				p2 = inputs[p2]
			}

			outputIndex = inputs[i+3]
		}

		jumpIndex := 0
		switch opCode {
		case 1:
			inputs[outputIndex] = p1 + p2
			// fmt.Printf("%v + %v = %v (Stored @ %v)\n", p1, p2, inputs[outputIndex], outputIndex)
		case 2:
			inputs[outputIndex] = p1 * p2
			// fmt.Printf("%v * %v = %v (Stored @ %v)\n", p1, p2, inputs[outputIndex], outputIndex)
		case 3:
			inputs[inputs[i+1]] = userInput
			// fmt.Printf("input %v (Stored @ %v)\n", userInput, inputs[i+1])
		case 4:
			fmt.Println(inputs[inputs[i+1]])
			// fmt.Printf("output %v (Stored @ %v)\n", inputs[inputs[i+1]], inputs[i+1])
		case 5, 6:
			if (opCode == 5 && p1 != 0) || (opCode == 6 && p1 == 0) {
				jumpIndex = p2
			} else {
				jumpIndex = i + 3
			}
		case 7:
			inputs[outputIndex] = 0
			if p1 < p2 {
				inputs[outputIndex] = 1
			}
		case 8:
			inputs[outputIndex] = 0
			if p1 == p2 {
				inputs[outputIndex] = 1
			}
		}

		// move the instruction pointer
		if opCode == 1 || opCode == 2 || opCode == 7 || opCode == 8 {
			i += 4
		} else if opCode == 3 || opCode == 4 {
			i += 2
		} else if opCode == 5 || opCode == 6 {
			i = jumpIndex
		}
	}
}

func main() {

	// Printing out file contents
	inputfile := "day5input.txt"

	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	inputs := stringToIntArr(input)
	runProgram(inputs, 1)
	fmt.Println("P1 Done!")

	inputs = stringToIntArr(input)
	runProgramP2(inputs, 8)
	fmt.Println("P2 Done!")
}
