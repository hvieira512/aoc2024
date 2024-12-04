package main

import (
	"fmt"

	u "github.com/hvieira512/aoc2024/cmd/utils"
)

var directions = [][2]int{
	{-1, 0}, {1, 0},
	{0, -1}, {0, 1},
	{-1, -1}, {1, 1},
	{-1, 1}, {1, -1},
}

func partOne(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	wordRunes := []rune(word)
	count := 0

	checkDirection := func(x, y, dx, dy int) bool {
		for i := 0; i < len(word); i++ {
			nx, ny := x+dx*i, y+dy*i
			fmt.Println(nx, ny)
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != wordRunes[i] {
				return false
			}
		}
		return true
	}

	for i := range rows {
		for j := range cols {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if checkDirection(i, j, dx, dy) {
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
	fmt.Printf("Part 1: %v\n", partOne(grid, "XMAS"))
}
