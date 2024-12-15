package main

import (
	"fmt"
	"log"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

const (
	Filename = "example"
)

func main() {
	utils.RenderDayHeader(15)
	lines, err := utils.Strings("cmd/day15/" + Filename + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %v\n", partOne(lines))
	// fmt.Printf("Part 2: %v\n", partTwo(lines))
}

func partOne(lines []string) int {
	result := 0

	for i := range lines {
		fmt.Println(lines[i])
	}

	return result
}

func partTwo(lines []string) int {
	panic("unimplemented")
}
