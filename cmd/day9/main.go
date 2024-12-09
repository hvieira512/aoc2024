package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func firstParse(diskmap string) (blocks string) {
	position, fileID := 0, 0
	for i := range diskmap {
		digit := string(diskmap[i])
		aux, _ := strconv.Atoi(digit)

		if i%2 == 0 {
			for j := 0; j < aux; j++ {
				blocks += fmt.Sprintf("%d", fileID)
				position++
			}
			fileID++
		} else {
			aux, _ := strconv.Atoi(digit)
			for j := 0; j < aux; j++ {
				blocks += fmt.Sprintf(".")
			}
		}
	}
	return blocks
}

func helper(diskmap string) bool {
	idx := strings.Index(diskmap, ".")
	firstHalf := diskmap[0:idx]
	secondHalf := diskmap[idx:]

	if strings.Contains(firstHalf, ".") || strings.ContainsAny(secondHalf, "0123456789") {
		return false
	}
	return true
}

func secondParse(diskmap string) string {
	for i := 0; i < 5; i++ {
		if helper(diskmap) {
			break
		}

		fmt.Println(diskmap)
	}

	return diskmap
}

func partOne(diskmap string) int {
	result := 0

	diskmap = firstParse(diskmap)
	diskmap = secondParse(diskmap)

	return result
}

func main() {
	utils.RenderDayHeader(0)
	lines, _ := utils.ReadLines("cmd/day9/example.txt")
	diskmap := lines[0]

	fmt.Printf("\nPart 1: %v\n", partOne(diskmap))
	// fmt.Printf("Part 2: %v\n", partTwo(diskmap))
}
