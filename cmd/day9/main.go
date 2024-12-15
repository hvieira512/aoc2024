package main

import (
	"fmt"
	"strconv"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func main() {
	utils.RenderDayHeader(9)
	lines, _ := utils.Strings("cmd/day9/input.txt")
	diskmap := lines[0]

	fmt.Printf("Part 1: %v\n", partOne(diskmap))
	fmt.Printf("Part 2: %v\n", partTwo(diskmap))
}

func partOne(diskmap string) int {
	disk := getDisk(diskmap)
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

func getDisk(diskmap string) []int {
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

// part two
func partTwo(diskmap string) int {
	files := map[int][2]int{}
	blanks := [][2]int{}

	fid := 0
	pos := 0

	for i, char := range diskmap {
		x, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			if x == 0 {
				panic("bro what, this shouldn't be there")
			}
			files[fid] = [2]int{pos, x}
			fid++
		} else {
			if x != 0 {
				blanks = append(blanks, [2]int{pos, x})
			}
		}
		pos += x
	}

	for fid > 0 {
		fid--
		pos, size := files[fid][0], files[fid][1]
		for i, blank := range blanks {
			start, length := blank[0], blank[1]
			if start >= pos {
				blanks = blanks[:i]
				break
			}
			if size <= length {
				files[fid] = [2]int{start, size}
				if size == length {
					blanks = append(blanks[:i], blanks[i+1:]...)
				} else {
					blanks[i] = [2]int{start + size, length - size}
				}
				break
			}
		}
	}

	total := 0
	for fid, file := range files {
		fmt.Println(fid, file)
		pos, size := file[0], file[1]
		for x := pos; x < pos+size; x++ {
			total += fid * x
		}
	}

	return total
}
