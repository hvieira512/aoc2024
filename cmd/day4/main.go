package main

import (
	"fmt"
	"log"

	u "github.com/hvieira512/aoc2024/cmd/utils"
)

func partOne(lines []string) int {
	result := 0

	for _, line := range lines {
		fmt.Println(line)
	}

	return result
}

func partTwo(lines []string) int {
	panic("unimplemented")
}

func main() {
	u.RenderDayHeader(4)
	lines, err := u.ReadLines("cmd/day4/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %v\n", partOne(lines))
	// fmt.Printf("Part 2: %v\n", partTwo(lines))
}
