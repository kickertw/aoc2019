package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func p1FuelCalc(x int) int {
	return x/3 - 2
}

func p2FuelCalc(x int) int {
	tempFuelVal := p1FuelCalc(x)
	retVal := 0
	for tempFuelVal > 0 {
		retVal += tempFuelVal
		tempFuelVal = p1FuelCalc(tempFuelVal)
	}

	return retVal
}

func main() {

	// Printing out file contents
	inputfile := "day1input.txt"

	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p1FinalAnswer, p2FinalAnswer := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, _ := strconv.Atoi(scanner.Text())
		p1FinalAnswer += p1FuelCalc(input)
		p2FinalAnswer += p2FuelCalc(input)
	}

	fmt.Printf("P1 Final Answer = %v \n", p1FinalAnswer)
	fmt.Printf("P2 Final Answer = %v", p2FinalAnswer)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
