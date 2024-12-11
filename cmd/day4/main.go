package main

import (
	"fmt"

	u "github.com/hvieira512/aoc2024/cmd/utils"
)

func inBounds(x, y, n int) bool {
	return x >= 0 && y >= 0 && x < n && y < n
}

func checkDirection(x, y, dx, dy int, word string, grid [][]rune) bool {
	n := len(grid)
	wordRunes := []rune(word)

	for i := 0; i < len(word); i++ {
		nx, ny := x+dx*i, y+dy*i

		if !inBounds(nx, ny, n) || grid[nx][ny] != wordRunes[i] {
			return false
		}
	}
	return true
}

func partOne(grid [][]rune, word string) int {
	n := len(grid)
	count := 0
	directions := [][2]int{
		{-1, 0}, {1, 0},
		{0, -1}, {0, 1},
		{-1, -1}, {1, 1},
		{-1, 1}, {1, -1},
	}

	for i := range n {
		for j := range n {
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

func hasXmas(i, j, n int, grid [][]rune) bool {
	if !(1 <= i && i < n-1 && 1 <= j && j < n-1) {
		return false
	}

	diag1 := string(grid[i-1][j-1]) + string(grid[i][j]) + string(grid[i+1][j+1])
	diag2 := string(grid[i-1][j+1]) + string(grid[i][j]) + string(grid[i+1][j-1])

	if (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM") {
		return true
	}

	return false
}

func partTwo(grid [][]rune) int {
	count := 0
	n := len(grid)

	for i := range n {
		for j := range n {
			if grid[i][j] == 'A' {
				if hasXmas(i, j, n, grid) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	grid, _ := u.ReadGridRune("cmd/day4/input.txt")

	fmt.Printf("Part 1: %v\n", partOne(grid, "XMAS"))
	fmt.Printf("Part 2: %v\n", partTwo(grid))
}
