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
	blocks := strings.Split(diskmap, ".")

	count := 0
	for i := range blocks {
		if blocks[i] != "" {
			count++
		}
	}

	if count != 1 {
		return false
	}

	return true
}

func getBlocks(diskmap string) []string {
	auxBlocks := strings.Split(diskmap, ".")
	blocks := []string{}
	for _, auxBlock := range auxBlocks {
		if len(auxBlock) != 0 {
			blocks = append(blocks, auxBlock)
		}
	}
	return blocks
}

func secondParse(diskmap string) string {
	for i := 0; i < 10; i++ {
		if isSPComplete(diskmap) {
			return diskmap
		}

		// get the first and last block of this iteration
		blocks := getBlocks(diskmap)

		fBlock := blocks[0]
		lBlock := blocks[len(blocks)-1]

		fBlockLast := getIdxLastDigitBlock(fBlock, diskmap, 1)
		lBlockLast := getIdxLastDigitBlock(lBlock, diskmap, -1)

		if fBlockLast != -1 && lBlockLast != -1 && fBlockLast+1 < len(diskmap) {
			swapCharAtIndex(&diskmap, fBlockLast+1, rune(diskmap[lBlockLast]))
			swapCharAtIndex(&diskmap, lBlockLast, '.')
		}
	}
	return ""
}

func swapCharAtIndex(s *string, index int, newChar rune) {
	runes := []rune(*s)
	if index >= 0 && index < len(runes) {
		runes[index] = newChar
	}
	*s = string(runes)
}

func getIdxLastDigitBlock(block, diskmap string, order int) int {
	blockIdx := -1
	if order == 1 {
		blockIdx = strings.Index(diskmap, block)
	} else if order == -1 {
		for i := len(diskmap) - len(block); i >= 0; i-- {
			if diskmap[i:i+len(block)] == block {
				blockIdx = i
				break
			}
		}
	}

	for i := len(block) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(block[i])) {
			return blockIdx + i
		}
	}

	return -1
}

func partOne(diskmap string) int {
	result := 0

	diskmap = "12345"
	diskmap = firstParse(diskmap)
	diskmap = secondParse(diskmap)
	fmt.Println(diskmap)

	return result
}

func main() {
	utils.RenderDayHeader(0)
	lines, _ := utils.ReadLines("cmd/day9/example.txt")
	diskmap := lines[0]

	fmt.Printf("\nPart 1: %v\n", partOne(diskmap))
	// fmt.Printf("Part 2: %v\n", partTwo(diskmap))
}
