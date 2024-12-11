package main

import (
	"fmt"
	"strconv"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func main() {
	utils.RenderDayHeader(9)
	lines, _ := utils.ReadLines("cmd/day9/input.txt")
	diskmap := lines[0]

	fmt.Printf("\nPart 1: %v\n", partOne(diskmap))
}

func partOne(diskmap string) int {
	disk := getFilesystem(diskmap)
	blanks := getBlanks(disk)
	disk = ampiphod(disk, blanks)

	return checksum(disk)
}

func checksum(disk []int) int {
	sum := 0

	for i, x := range disk {
		sum += i * x
	}

	return sum
}

func ampiphod(disk []int, blanks []int) []int {
	for _, i := range blanks {
		for len(disk) > 0 && disk[len(disk)-1] == -1 {
			disk = disk[:len(disk)-1]
		}
		if len(disk) <= i {
			break
		}
		disk[i] = disk[len(disk)-1]
		disk = disk[:len(disk)-1]
	}
	return disk
}

func getBlanks(filesystem []int) []int {
	blanks := []int{}
	for i, x := range filesystem {
		if x == -1 {
			blanks = append(blanks, i)
		}
	}
	return blanks
}

func getFilesystem(diskmap string) []int {
	filesystem := []int{}
	fileID := 0

	for i, char := range diskmap {
		x, _ := strconv.Atoi(string(char))
		if i%2 == 0 {
			for range x {
				filesystem = append(filesystem, fileID)
			}
			fileID++
		} else {
			for range x {
				filesystem = append(filesystem, -1)
			}
		}
	}
	return filesystem
}
