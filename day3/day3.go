package main

import (
	"bufio"
	"fmt"
	"image"
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

func readInput() ([]string, []string) {
	inputfile := "day3input.txt"

	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	retVal1 := strings.Split(scanner.Text(), ",")
	scanner.Scan()
	retVal2 := strings.Split(scanner.Text(), ",")

	return retVal1, retVal2
}

func pointToString(point image.Point) string {
	return strconv.Itoa(point.X) + "," + strconv.Itoa(point.Y)
}

func moveWire(visitedPoints map[string]image.Point, location image.Point, input string) image.Point {
	dir := input[:1]
	steps, _ := strconv.Atoi(input[1:])

	switch dir {
	case "U":
		newY := location.Y + steps
		for location.Y < newY {
			location.Y++
			_, ok := visitedPoints[pointToString(location)]
			if !ok {
				visitedPoints[pointToString(location)] = location
			}
			// fmt.Printf("  Moved U - (%v,%v)", location.X, location.Y)
		}
	case "D":
		newY := location.Y - steps
		for location.Y > newY {
			location.Y--
			_, ok := visitedPoints[pointToString(location)]
			if !ok {
				visitedPoints[pointToString(location)] = location
			}
			// fmt.Printf("  Moved D - (%v,%v)", location.X, location.Y)
		}
	case "L":
		newX := location.X - steps
		for location.X > newX {
			location.X--
			_, ok := visitedPoints[pointToString(location)]
			if !ok {
				visitedPoints[pointToString(location)] = location
			}
			// fmt.Printf("  Moved L - (%v,%v)", location.X, location.Y)
		}
	case "R":
		newX := location.X + steps
		for location.X < newX {
			location.X++
			_, ok := visitedPoints[pointToString(location)]
			if !ok {
				visitedPoints[pointToString(location)] = location
			}
			// fmt.Printf("  Moved R - (%v,%v)", location.X, location.Y)
		}
	default:
		fmt.Println("Invalid Direction! - " + dir)
	}

	return location
}

func intAbs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}

func findMinManhattenDist(wire1 map[string]image.Point, wire2 map[string]image.Point, originX int, originY int) {
	minDist := -1
	for key, coord := range wire1 {
		_, ok := wire2[key]

		// if the key/coordinate from wire1 exists in wire2, we have an overlap
		if ok {
			tempDist := intAbs(originX-coord.X) + intAbs(originY-coord.Y)
			if minDist == -1 || tempDist < minDist {
				minDist = tempDist
			}
		}
	}

	fmt.Printf("P1 - The shortest Manhatten distance is %v", minDist)
}

func drawWire(input []string) map[string]image.Point {
	retVal := make(map[string]image.Point)
	location := image.Pt(0, 0)
	for _, step := range input {
		location = moveWire(retVal, location, step)
	}

	return retVal
}

func main() {
	wire1input, wire2input := readInput()

	wire1 := drawWire(wire1input)
	wire2 := drawWire(wire2input)

	findMinManhattenDist(wire1, wire2, 0, 0)
}
