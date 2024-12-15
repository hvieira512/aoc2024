package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

type Equation struct {
	target  int
	numbers []int
}

func getEquations(lines []string) []Equation {
	equations := []Equation{}

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		test, _ := strconv.Atoi(parts[0])
		aux := strings.Fields(parts[1])
		numbers := []int{}
		for _, numStr := range aux {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
		equations = append(equations, Equation{target: test, numbers: numbers})
	}

	return equations
}

func main() {
	utils.RenderDayHeader(7)
	lines, _ := utils.Strings("cmd/day7/input.txt")
	equations := getEquations(lines)

	fmt.Printf("Part 1: %v\n", partOne(equations))
	fmt.Printf("Part 2: %v\n", partTwo(equations))
}

func partOne(equations []Equation) int {
	result := 0
	operators := []rune{'+', '*'}

	for _, equation := range equations {
		solved := false
		combos := getOpsCombo(operators, len(equation.numbers)-1)

		for _, combo := range combos {
			if testEquation(equation, combo) {
				solved = true
			}
		}
		if solved {
			result += equation.target
			continue
		}
	}

	return result
}

func getOpsCombo(operators []rune, length int) [][]rune {
	var result [][]rune

	var backtrack func(curr []rune)
	backtrack = func(curr []rune) {
		if len(curr) == length {
			comboCopy := make([]rune, len(curr))
			copy(comboCopy, curr)
			result = append(result, comboCopy)
			return
		}

		for _, op := range operators {
			backtrack(append(curr, op))
		}
	}

	backtrack([]rune{})
	return result
}

func testEquation(equation Equation, combo []rune) bool {
	result := equation.numbers[0]

	for i, op := range combo {
		switch op {
		case '+':
			result += equation.numbers[i+1]
		case '*':
			result *= equation.numbers[i+1]
		}
	}

	return result == equation.target
}

// part two
func partTwo(equations []Equation) int {
	result := 0

	for _, equation := range equations {
		if canObtain(equation.target, equation.numbers) {
			result += equation.target
		}
	}

	return result
}

func canObtain(target int, array []int) bool {
	if len(array) == 1 {
		return target == array[0]
	}

	if target%array[len(array)-1] == 0 && canObtain(target/array[len(array)-1], array[:len(array)-1]) {
		return true
	}

	if target > array[len(array)-1] && canObtain(target-array[len(array)-1], array[:len(array)-1]) {
		return true
	}

	sTarget := strconv.Itoa(target)
	sLast := strconv.Itoa(array[len(array)-1])

	if len(sTarget) > len(sLast) && strings.HasSuffix(sTarget, sLast) {
		newTarget, _ := strconv.Atoi(sTarget[:len(sTarget)-len(sLast)])
		if canObtain(newTarget, array[:len(array)-1]) {
			return true
		}
	}

	return false
}
