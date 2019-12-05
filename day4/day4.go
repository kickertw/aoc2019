package main

import (
	"fmt"
	"strconv"
)

// Should probably do this and isValidP2Password at the same time
func isValidPassword(input int) bool {
	strInput := strconv.Itoa(input)
	hasAdjDupe := false

	for i := 0; i < len(strInput)-1; i++ {
		val1, _ := strconv.Atoi(string(strInput[i]))
		val2, _ := strconv.Atoi(string(strInput[i+1]))
		if val1 > val2 {
			return false
		}

		if strInput[i] == strInput[i+1] {
			hasAdjDupe = true
		}
	}

	return hasAdjDupe
}

func isValidP2Password(input int) bool {
	strInput := strconv.Itoa(input)
	charCount := make(map[int]int)

	// 012345
	// 178999
	for i := 0; i < len(strInput)-1; i++ {
		val1, _ := strconv.Atoi(string(strInput[i]))
		val2, _ := strconv.Atoi(string(strInput[i+1]))
		if val1 > val2 {
			return false
		}

		_, ok := charCount[val1]
		if !ok {
			charCount[val1] = 1
		} else {
			charCount[val1]++
		}

		if i == len(strInput)-2 {
			_, ok := charCount[val2]
			if !ok {
				charCount[val2] = 1
			} else {
				charCount[val2]++
			}
		}
	}

	for _, val := range charCount {
		if val == 2 {
			return true
		}
	}

	return false
}

func main() {
	low := 178416
	high := 676461

	p1Counter, p2Counter := 0, 0

	for i := low; i < high; i++ {
		if isValidPassword(i) {
			p1Counter++
		}

		if isValidP2Password(i) {
			p2Counter++
		}
	}

	fmt.Printf("P1 - found %v valid passwords\n", p1Counter)
	fmt.Printf("P2 - found %v valid passwords", p2Counter)
}
