package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

type Equation struct {
	test    int
	numbers []int
}

func parseInput(lines []string) []Equation {
	equations := []Equation{}

	for i := range lines {
		parts := strings.Split(lines[i], ": ")

		test, _ := strconv.Atoi(parts[0])

		numbersStr := strings.Fields(parts[1])
		numbers := []int{}
		for _, numberStr := range numbersStr {
			number, _ := strconv.Atoi(numberStr)
			numbers = append(numbers, number)
		}
		equations = append(equations, Equation{test: test, numbers: numbers})
	}

	return equations
}

func getOperators(operators []string, length int) []string {
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

func testEquation(equation Equation, ops []string) bool {
	for _, opAttempt := range ops {
		result := equation.numbers[0]
		for i := 0; i < len(equation.numbers)-1; i++ {
			if string(opAttempt[i]) == "+" {
				result += equation.numbers[i+1]
			} else if string(opAttempt[i]) == "*" {
				result *= equation.numbers[i+1]
			}
		}
		if result == equation.test {
			return true
		}
	}

	return false
}

func partOne(equations []Equation) int {
	result := 0
	opsChars := []string{"+", "*"}

	for _, equation := range equations {
		ops := getOperators(opsChars, len(equation.numbers)-1)

		if testEquation(equation, ops) {
			result += equation.test
		}
	}

	return result
}

// ------------------
// ---- PART TWO ----
// ------------------
func partTwo(equations []Equation) int {
	result := 0
	opsChars := []string{"+", "*", "||"}

	for _, equation := range equations {
		if equation.test != 7290 {
			continue
		}
		fmt.Println(equation)
		ops := getOperators(opsChars, len(equation.numbers)-1)

		if testEquation(equation, ops) {
			result += equation.test
		} else if testEquationV2(equation, ops) {
			result += equation.test
		}
	}

	return result
}

func testEquationV2(equation Equation, ops []string) bool {
	for _, opAttempt := range ops {
		equation.numbers = updateEquation(equation.numbers, opAttempt)
		result := equation.numbers[0]
		for i := 0; i < len(equation.numbers)-1; i++ {
			if string(opAttempt[i]) == "+" {
				result += equation.numbers[i+1]
			} else if string(opAttempt[i]) == "*" {
				result *= equation.numbers[i+1]
			}
		}
		if result == equation.test {
			return true
		}
	}
	return false
}

func updateEquation(numbers []int, op string) []int {
	for i := 0; i < len(numbers)-1; i++ {
		fmt.Println(op, i)
		if op == "||" {
			fmt.Println(numbers[i], numbers[i+1])
			concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", numbers[i], numbers[i+1]))
			numbers[i] = concatenated
			numbers = append(numbers[:i+1], numbers[i+2:]...)
			i--
		}
	}
	return numbers
}

func main() {
	utils.RenderDayHeader(7)
	lines, _ := utils.ReadLines("cmd/day7/example.txt")
	equations := parseInput(lines)

	// fmt.Printf("Part 1: %v\n", partOne(equations))
	fmt.Printf("Part 2: %v\n", partTwo(equations))
}
