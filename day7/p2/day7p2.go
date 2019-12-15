package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Amp - representing an amplifier
type Amp struct {
	id             string
	phaseSetting   int
	lastOutput     int
	lastPointerIdx int
}

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

// Run's the intCode.
// 1st output = result
// 2nd output = last pointer location.
// 3rd output = Keep the program going!
func runProgram(pointerIndex int, inputs []int, phaseSetting int, ampOutput int, isLastAmp bool) (int, int, bool) {
	i := 0
	phaseIsSet := false
	if pointerIndex >= 0 {
		phaseIsSet = true
		i = pointerIndex
	}

	for i < len(inputs) {
		op := inputs[i]
		if op == 99 {
			keepSwimming := true

			if isLastAmp {
				keepSwimming = false
			}

			return ampOutput, i, keepSwimming
		}

		opCode := op % 10
		p1Mode := (op / 100) % 10
		p2Mode := (op / 1000) % 10

		p1 := inputs[i+1]
		p2 := inputs[i+2]
		outputIndex := 0

		if p1Mode == 0 {
			p1 = inputs[p1]
		}

		if opCode != 3 && opCode != 4 {
			if p2Mode == 0 {
				p2 = inputs[p2]
			}

			if opCode == 1 || opCode == 2 || opCode == 7 || opCode == 8 {
				outputIndex = inputs[i+3]
			}
		}

		jumpIndex := 0
		switch opCode {
		case 1:
			inputs[outputIndex] = p1 + p2
			//fmt.Printf("%v + %v = %v (Stored @ %v)\n", p1, p2, inputs[outputIndex], outputIndex)
		case 2:
			inputs[outputIndex] = p1 * p2
			//fmt.Printf("%v * %v = %v (Stored @ %v)\n", p1, p2, inputs[outputIndex], outputIndex)
		case 3:
			if phaseIsSet {
				inputs[inputs[i+1]] = ampOutput
			} else {
				inputs[inputs[i+1]] = phaseSetting
				phaseIsSet = true
			}

			//fmt.Printf("input %v (Stored @ %v)\n", inputs[inputs[i+1]], inputs[i+1])
		case 4:
			ampOutput = p1
			return ampOutput, i + 2, true
			//fmt.Printf("Amp Output = %v\n", p1)
		case 5, 6:
			if (opCode == 5 && p1 != 0) || (opCode == 6 && p1 == 0) {
				//fmt.Print("Conditional met! - ")
				jumpIndex = p2
			} else {
				//fmt.Print("Conditional not met! - ")
				jumpIndex = i + 3
			}
			//fmt.Printf("Jumping to index %v\n", jumpIndex)
		case 7:
			inputs[outputIndex] = 0
			if p1 < p2 {
				inputs[outputIndex] = 1
			}
			//fmt.Printf("Setting %v @ index %v\n", inputs[outputIndex], outputIndex)
		case 8:
			inputs[outputIndex] = 0
			if p1 == p2 {
				inputs[outputIndex] = 1
			}
			//fmt.Printf("Setting %v @ index %v\n", inputs[outputIndex], outputIndex)
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

	return ampOutput, i, true
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
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

	inputs := stringToIntArr(input)

	winningCombo := []int{}
	//largestOutput := 0
	ampInput := []int{9, 8, 7, 6, 5}
	//perms := permutations(ampInput)

	amplifiers := []Amp{
		Amp{id: "A", lastOutput: 0, phaseSetting: ampInput[0], lastPointerIdx: -1},
		Amp{id: "B", lastOutput: 0, phaseSetting: ampInput[1], lastPointerIdx: -1},
		Amp{id: "C", lastOutput: 0, phaseSetting: ampInput[2], lastPointerIdx: -1},
		Amp{id: "D", lastOutput: 0, phaseSetting: ampInput[3], lastPointerIdx: -1},
		Amp{id: "E", lastOutput: 0, phaseSetting: ampInput[4], lastPointerIdx: -1},
	}

	pointerIndex := -1
	runningAmpIdx := 0
	nextInput := 0
	keepSwimming := true
	for keepSwimming {
		fmt.Printf("Amp[%v] starting at index [%v]\n", amplifiers[runningAmpIdx].id, amplifiers[runningAmpIdx].lastPointerIdx)
		nextInput, pointerIndex, keepSwimming = runProgram(amplifiers[runningAmpIdx].lastPointerIdx, inputs, amplifiers[runningAmpIdx].phaseSetting, nextInput, runningAmpIdx == 4)

		amplifiers[runningAmpIdx].lastPointerIdx = pointerIndex + 1
		amplifiers[runningAmpIdx].lastOutput = nextInput

		fmt.Printf("  Ran w/ output %v - stopped at %v\n", nextInput, amplifiers[runningAmpIdx].lastPointerIdx)

		runningAmpIdx++
		if runningAmpIdx == 5 {
			runningAmpIdx = 0
		}
	}
	// for ii := 0; ii < len(perms); ii++ {
	// 	ampOutput := 0
	// 	for jj := 0; jj < len(perms[ii]); jj++ {
	// 		ampOutput = runProgram(inputs, perms[ii][jj], ampOutput)
	// 	}

	// 	if ampOutput > largestOutput {
	// 		largestOutput = ampOutput
	// 		winningCombo = perms[ii]
	// 	}
	// }

	fmt.Printf("Largest Output = %v / Amp Seq = %v", amplifiers[4].lastOutput, winningCombo)
}
