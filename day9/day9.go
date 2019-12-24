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

func stringToIntMap(input string) map[int]int {
	s := strings.Split(input, ",")

	var retVal = make(map[int]int)

	for ii, val := range s {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		retVal[ii] = intVal
	}

	return retVal
}

func runProgram(inputs map[int]int, userInput int) map[int]int {
	i := 0
	relativeBase := 0

	for {
		op := 0
		if val, ok := inputs[i]; ok {
			op = val
		} else {
			inputs[i] = 0
		}

		if op == 99 {
			return inputs
		}
		// fmt.Printf("@ Index[%v] - Opcode = %v\n", i, op)

		opCode := op % 10
		p1Mode := (op / 100) % 10
		p2Mode := (op / 1000) % 10

		p1 := inputs[i+1]
		p2 := inputs[i+2]
		outputIndex := 0

		if p1Mode == 0 {
			p1 = inputs[p1]
		} else if p1Mode == 2 {
			p1 = inputs[relativeBase+p1]
		}

		if opCode != 3 && opCode != 4 && opCode != 9 {
			if p2Mode == 0 {
				p2 = inputs[p2]
			} else if p2Mode == 2 {
				p2 = inputs[relativeBase+p2]
			}

			if opCode == 1 || opCode == 2 || opCode == 7 || opCode == 8 {
				outputMode := (op / 10000) % 10

				outputIndex = inputs[i+3]
				if outputMode == 2 {
					outputIndex += relativeBase
				}
			}
		}

		jumpIndex := 0
		switch opCode {
		case 1:
			inputs[outputIndex] = p1 + p2
			fmt.Printf("%v + %v = %v (Stored @ %v)\n", p1, p2, inputs[outputIndex], outputIndex)
		case 2:
			inputs[outputIndex] = p1 * p2
			fmt.Printf("%v * %v = %v (Stored @ %v)\n", p1, p2, inputs[outputIndex], outputIndex)
		case 3:
			storageIdx := inputs[i+1]
			if p1Mode == 2 {
				storageIdx = storageIdx + relativeBase
			}

			inputs[storageIdx] = userInput
			fmt.Printf("input %v (Stored @ %v)\n", userInput, storageIdx)
		case 4:
			fmt.Printf("Output - %v\n", p1)
		case 5, 6:
			if (opCode == 5 && p1 != 0) || (opCode == 6 && p1 == 0) {
				fmt.Print("Conditional met! - ")
				jumpIndex = p2
			} else {
				fmt.Print("Conditional not met! - ")
				jumpIndex = i + 3
			}
			fmt.Printf("Jumping to index %v\n", jumpIndex)
		case 7:
			inputs[outputIndex] = 0
			if p1 < p2 {
				inputs[outputIndex] = 1
			}
			fmt.Printf("(%v < %v)Setting %v @ index %v\n", p1, p2, inputs[outputIndex], outputIndex)
		case 8:
			inputs[outputIndex] = 0
			if p1 == p2 {
				inputs[outputIndex] = 1
			}
			fmt.Printf("(%v == %v)Setting %v @ index %v\n", p1, p2, inputs[outputIndex], outputIndex)
		case 9:
			relativeBase += p1
			fmt.Printf("relativeBase is now %v\n", relativeBase)
		}

		// move the instruction pointer
		if opCode == 1 || opCode == 2 || opCode == 7 || opCode == 8 {
			i += 4
		} else if opCode == 3 || opCode == 4 || opCode == 9 {
			i += 2
		} else if opCode == 5 || opCode == 6 {
			i = jumpIndex
		}
	}
}

func main() {

	// Get User Input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')

	userInput, err := strconv.Atoi(strings.TrimSuffix(text, "\r\n"))
	if err != nil {
		userInput = 5
	}
	// userInput := 1 // debug only.
	fmt.Printf("Input = %v\n", userInput)

	// Printing out file contents
	inputfile := "input.txt"

	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	inputs := stringToIntMap(input)
	inputs = runProgram(inputs, userInput)
	// fmt.Println(inputs)
}
