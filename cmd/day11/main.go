package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func main() {
	utils.RenderDayHeader(11)
	lines, _ := utils.ReadLines("cmd/day11/input.txt")
	stones := getStones(lines)

	fmt.Printf("Part 1: %v\n", solve(stones, 25))
	fmt.Printf("Part 2: %v\n", solve(stones, 75))
}

func getStones(lines []string) []int {
	stones := []int{}
	stonesStr := strings.Split(lines[0], " ")
	for _, str := range stonesStr {
		stone, _ := strconv.Atoi(str)
		stones = append(stones, stone)
	}
	return stones
}

var memo = map[string]int{}

func count(stone, steps int) int {
	key := fmt.Sprintf("%d-%d", stone, steps)

	if val, exists := memo[key]; exists {
		return val
	}

	if steps == 0 {
		return 1
	}

	var result int
	if stone == 0 {
		result = count(1, steps-1)
	} else {
		numStr := strconv.Itoa(stone)
		length := len(numStr)
		if length%2 == 0 {
			mid := length / 2
			left, _ := strconv.Atoi(numStr[:mid])
			right, _ := strconv.Atoi(numStr[mid:])
			result = count(left, steps-1) + count(right, steps-1)
		} else {
			result = count(stone*2024, steps-1)
		}
	}

	memo[key] = result
	return result
}

func solve(stones []int, blinks int) int {
	total := 0
	for _, stone := range stones {
		total += count(stone, blinks)
	}

	return total
}
