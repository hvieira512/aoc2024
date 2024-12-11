package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func getStones(line string) []int64 {
	stones := []int64{}
	stonesStr := strings.Split(line, " ")
	for _, str := range stonesStr {
		stone, _ := strconv.ParseInt(str, 10, 64)
		stones = append(stones, stone)
	}
	return stones
}

func blink(stones []int64) []int64 {
	newStones := make([]int64, 0, len(stones)*2)

	for _, stone := range stones {
		numDigits := lenNum(stone)

		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}

		if numDigits%2 == 0 {
			mid := numDigits / 2
			divisor := int64(math.Pow10(mid))
			left := stone / divisor
			right := stone % divisor
			newStones = append(newStones, left, right)
			continue
		}
		newStones = append(newStones, stone*2024)
	}

	return newStones
}

func lenNum(i int64) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func solve(stones []int64, blinks int) int {
	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}
	return len(stones)
}

func main() {
	utils.RenderDayHeader(11)
	lines, _ := utils.ReadLines("cmd/day11/input.txt")
	stones := getStones(lines[0])

	// fmt.Printf("Part 1: %v\n", solve(stones, 25))
	fmt.Printf("Part 2: %v\n", solve(stones, 75))
}
