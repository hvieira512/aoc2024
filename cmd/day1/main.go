package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	u "github.com/hvieira512/aoc2024/cmd/utils"
)

func partOne(lines []string) int {
	leftValues, rightValues := getNumbers(lines)
	result := 0

	for len(leftValues) != 0 {
		lPos, leftie := getSmallest(leftValues)
		rPos, rightie := getSmallest(rightValues)

		result += int(math.Abs(float64(leftie - rightie)))

		leftValues = u.DeleteAtIndex(leftValues, lPos)
		rightValues = u.DeleteAtIndex(rightValues, rPos)
	}

	return result
}

func partTwo(lines []string) int {
	leftValues, rightValues := getNumbers(lines)
	result := 0

	for _, leftValue := range leftValues {
		count := 0
		for _, rightValue := range rightValues {
			if leftValue == rightValue {
				count++
			}
		}
		result += leftValue * count
	}

	return result
}

func getSmallest(numbers []int) (int, int) {
	pos := 0
	smallest := numbers[pos]
	for i, number := range numbers {
		if number < smallest {
			smallest = number
			pos = i
		}
	}

	return pos, smallest
}

func getNumbers(lines []string) ([]int, []int) {
	leftValues, rightValues := []int{}, []int{}
	for _, line := range lines {
		numbers := strings.Fields(line)

		left, _ := strconv.Atoi(numbers[0])
		right, _ := strconv.Atoi(numbers[1])

		leftValues = append(leftValues, left)
		rightValues = append(rightValues, right)
	}

	return leftValues, rightValues
}

func main() {
	u.RenderDayHeader(1)
	lines, err := u.Strings("cmd/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %v\n", partOne(lines))
	fmt.Printf("Part 2: %v\n", partTwo(lines))
}
