package main

import (
	"fmt"

	u "github.com/hvieira512/aoc2024/cmd/utils"
)

func checkDirection(x, y, dx, dy int, word string, grid [][]rune) bool {
	rows := len(grid)
	cols := len(grid[0])
	wordRunes := []rune(word)

	for i := 0; i < len(word); i++ {
		nx, ny := x+dx*i, y+dy*i
		outOfBounds := nx < 0 || ny < 0 || nx >= rows || ny >= cols

		if outOfBounds || grid[nx][ny] != wordRunes[i] {
			return false
		}
	}
	return true
}

func partOne(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	directions := [][2]int{
		{-1, 0}, {1, 0},
		{0, -1}, {0, 1},
		{-1, -1}, {1, 1},
		{-1, 1}, {1, -1},
	}

	for i := range rows {
		for j := range cols {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if checkDirection(i, j, dx, dy, word, grid) {
					count++
				}
			}
		}
	}

	return count
}

func partTwo(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	directions := [][2]int{
		{-1, -1}, {1, 1},
		{-1, 1}, {1, -1},
	}

	for i := range rows {
		for j := range cols {
			if grid[i][j] == 'A' {
				runesCount := map[rune]int{}
				for _, dir := range directions {
					nx, ny := i+dir[0], j+dir[1]
					outOfBounds := nx < 0 || ny < 0 || nx >= rows || ny >= cols
					if !outOfBounds && grid[nx][ny] != 'X' {
						runesCount[grid[nx][ny]]++
					}
				}
				// delete non-valid x-mas
				if len(runesCount) != 2 || runesCount['M'] != 2 || runesCount['S'] != 2 {
					delete(runesCount, 'M')
					delete(runesCount, 'S')
				}

				if runesCount['M'] == 2 && runesCount['S'] == 2 {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	grid, err := u.ReadGrid("cmd/day4/example.txt")
	if err != nil {
		fmt.Printf("Error loading grid: %v\n", err)
		return
	}
	// fmt.Printf("Part 1: %v\n", partOne(grid, "XMAS"))
	fmt.Printf("Part 2: %v\n", partTwo(grid))
}
