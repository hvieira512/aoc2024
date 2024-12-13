package main

import (
	"fmt"
	"log"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func main() {
	utils.RenderDayHeader(13)
	lines, err := utils.ReadLines("cmd/day13/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %v\n", partOne(lines))
	// fmt.Printf("Part 2: %v\n", partTwo(lines))
}

func partOne(lines []string) int {
	panic("unimplemented")
}
