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
	intCode        []int
	phaseSetting   int
	currentInput   int
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
func runProgram(pointerIndex int, intCode []int, phaseSetting int, currentInput int, isLastAmp bool) (int, int, bool) {
	i := 0
	phaseIsSet := false
	var ampOutput int

	if pointerIndex >= 0 {
		phaseIsSet = true
		i = pointerIndex
	}

	for i < len(intCode) {
		op := intCode[i]
		if op == 99 {
			keepSwimming := true

			if isLastAmp {
				keepSwimming = false
			}

			return ampOutput, -1, keepSwimming
		}

		opCode := op % 10
		p1Mode := (op / 100) % 10
		p2Mode := (op / 1000) % 10

		p1 := intCode[i+1]
		p2 := intCode[i+2]
		outputIndex := 0

		if p1Mode == 0 {
			p1 = intCode[p1]
		}

		if opCode != 3 && opCode != 4 {
			if p2Mode == 0 {
				p2 = intCode[p2]
			}

			if opCode == 1 || opCode == 2 || opCode == 7 || opCode == 8 {
				outputIndex = intCode[i+3]
			}
		}

		switch opCode {
		case 1:
			intCode[outputIndex] = p1 + p2
			i += 4
		case 2:
			intCode[outputIndex] = p1 * p2
			i += 4
		case 3:
			if phaseIsSet && currentInput >= 0 {
				intCode[intCode[i+1]] = currentInput
				currentInput = -1
			} else if !phaseIsSet {
				intCode[intCode[i+1]] = phaseSetting
				phaseIsSet = true
			} else {
				return ampOutput, i, true
			}
			i += 2
		case 4:
			ampOutput = p1
			i += 2
		case 5, 6:
			if (opCode == 5 && p1 != 0) || (opCode == 6 && p1 == 0) {
				i = p2
			} else {
				i += 3
			}
		case 7:
			intCode[outputIndex] = 0
			if p1 < p2 {
				intCode[outputIndex] = 1
			}
			i += 4
		case 8:
			intCode[outputIndex] = 0
			if p1 == p2 {
				intCode[outputIndex] = 1
			}
			i += 4
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

	intCode := stringToIntArr(input)

	var thrusterOutput int
	winningCombo := []int{}
	largestOutput := 0
	seq := []int{5, 6, 7, 8, 9}
	perms := permutations(seq)

	for ii := 0; ii < len(perms); ii++ {
		for jj := 0; jj < len(perms[ii]); jj++ {
			ampOutput := 0
			ampInput := perms[ii]
			thrusterOutput = 0
			amplifiers := []Amp{
				Amp{id: "A", phaseSetting: ampInput[0], currentInput: 0, lastPointerIdx: -1},
				Amp{id: "B", phaseSetting: ampInput[1], lastPointerIdx: -1},
				Amp{id: "C", phaseSetting: ampInput[2], lastPointerIdx: -1},
				Amp{id: "D", phaseSetting: ampInput[3], lastPointerIdx: -1},
				Amp{id: "E", phaseSetting: ampInput[4], lastPointerIdx: -1},
			}

			for i := 0; i < len(amplifiers); i++ {
				amplifiers[i].intCode = make([]int, len(intCode))
				copy(amplifiers[i].intCode, intCode)
			}

			pointerIndex := -1
			runningAmpIdx := 0
			keepSwimming := true
			for keepSwimming {
				ampOutput, pointerIndex, keepSwimming = runProgram(amplifiers[runningAmpIdx].lastPointerIdx, amplifiers[runningAmpIdx].intCode, amplifiers[runningAmpIdx].phaseSetting, amplifiers[runningAmpIdx].currentInput, runningAmpIdx == 4)
				amplifiers[runningAmpIdx].lastPointerIdx = pointerIndex

				//fmt.Printf("Loop %v - Amp[%v] - input[%v] - output[%v] - ended @ index[%v]\n", loop, runningAmpIdx, amplifiers[runningAmpIdx].currentInput, ampOutput, pointerIndex)

				runningAmpIdx++
				if runningAmpIdx == 5 {
					runningAmpIdx = 0
				}

				if keepSwimming {
					amplifiers[runningAmpIdx].currentInput = ampOutput
				} else {
					thrusterOutput = ampOutput
				}
			}
		}

		if thrusterOutput > largestOutput {
			largestOutput = thrusterOutput
			winningCombo = perms[ii]
			// fmt.Printf("New Largest Output = %v / Amp Seq = %v\n", thrusterOutput, winningCombo)
		}
	}

	fmt.Printf("Largest Output = %v / Amp Seq = %v", largestOutput, winningCombo)
}
