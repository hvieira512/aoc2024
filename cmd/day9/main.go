package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

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

func isSPComplete(diskmap string) bool {
	idx := strings.Index(diskmap, ".")
	firstHalf := diskmap[0:idx]
	secondHalf := diskmap[idx:]

	if strings.Contains(firstHalf, ".") || strings.ContainsAny(secondHalf, "0123456789") {
		return false
	}
	return true
}

func getBlocks(diskmap string) []string {
	auxBlocks := strings.Split(diskmap, ".")
	blocks := []string{}
	for _, auxBlock := range auxBlocks {
		if len(auxBlock) != 0 { // ignore empty lines
			blocks = append(blocks, auxBlock)
		}
	}
	return blocks
}

func secondParse(diskmap string) string {
	for i := 0; i < 10; i++ {
		if isSPComplete(diskmap) {
			break
		}

		// get the first and last block of this iteration
		blocks := getBlocks(diskmap)

		fBlock := blocks[0]
		lBlock := blocks[len(blocks)-1]

		fLastDigitIndex := findDigitIdxDiskmap(fBlock, diskmap)
		lLastDigitIndex := findDigitIdxDiskmap(lBlock, diskmap)

		fmt.Println(i)
		fmt.Println(fBlock, lBlock)
		fmt.Println(fLastDigitIndex, lLastDigitIndex)
		fmt.Println()

		if fLastDigitIndex != -1 && lLastDigitIndex != -1 && fLastDigitIndex+1 < len(diskmap) {
			swapCharAtIndex(&diskmap, fLastDigitIndex+1, rune(diskmap[lLastDigitIndex]))
			swapCharAtIndex(&diskmap, lLastDigitIndex, '.')
		}
	}

	return diskmap
}

func swapCharAtIndex(s *string, index int, newChar rune) {
	runes := []rune(*s)
	if index >= 0 && index < len(runes) {
		runes[index] = newChar
	}
	*s = string(runes)
}

func findDigitIdxDiskmap(block, diskmap string) int {
	blockIndex := strings.Index(diskmap, block)
	if blockIndex == -1 {
		return -1
	}

	lastDigitIndex := -1
	// search from the end of the block
	for i := len(block) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(block[i])) {
			lastDigitIndex = blockIndex + i
			break
		}
	}

	return lastDigitIndex
}

func partOne(diskmap string) int {
	result := 0

	diskmap = "12345"
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
