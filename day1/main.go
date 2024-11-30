package main

import (
	"fmt"

	u "github.com/hvieira512/aoc2024/utils"
)

type Example struct{}

func partOne(lines []string) int {
	for line := range lines {
		fmt.Println(line)
	}

	return 0
}

func main() {
	u.RenderDayHeader(1)

	lines, _ := u.ReadLines("input.tx")

	fmt.Printf("Part 1: %v\n", partOne(lines))
	// fmt.Printf("Part 2: %v\n", partTwo(lines))
}
