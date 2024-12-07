package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/hvieira512/aoc2024/cmd/utils"
)

func getOperatorsComb(operators []string, length int) []string {
	var result []string
	var helper func(current string, level int)

	helper = func(current string, level int) {
		if level == length {
			result = append(result, current)
			return
		}
		for _, op := range operators {
			helper(current+op, level+1)
		}
	}

	helper("", 0)
	return result
}

func getEquations(lines []string) map[int][]int {
	equations := map[int][]int{}

	for _, line := range lines {
		sepIdx := strings.Index(line, ":")
		result, _ := strconv.Atoi(line[0:sepIdx])

		numbersStr := strings.Fields(line[sepIdx+1:])
		intNumbers := []int{}
		for i := range numbersStr {
			intNumber, _ := strconv.Atoi(numbersStr[i])
			intNumbers = append(intNumbers, intNumber)
		}
		equations[result] = intNumbers
	}

	return equations
}

func partOne(lines []string) int {
	total := 0
	equations := getEquations(lines)
	allOps := []string{"+", "*"}

	for k, v := range equations {
		ops := getOperatorsComb(allOps, len(v)-1)
		found := false

		for i := range ops {
			tmpResult := v[0]
			for j := 0; j < len(v)-1; j++ {
				switch string(ops[i][j]) {
				case "+":
					tmpResult += v[j+1]
				case "*":
					tmpResult *= v[j+1]
				}
			}
			if k == tmpResult && !found {
				fmt.Printf("Testing equation: %d with numbers %v\n", k, v)
				fmt.Printf("  Combination %s => Result: %d\n", ops[i], tmpResult)
				total += k
				found = true
				break
			}
		}
	}

	return total
}

func main() {
	u.RenderDayHeader(7)
	lines, _ := u.ReadLines("cmd/day7/example.txt")

	fmt.Printf("Part 1: %v\n", partOne(lines))
	// fmt.Printf("Part 2: %v\n", partTwo(lines))
}
