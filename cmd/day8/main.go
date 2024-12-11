package main

import (
	"fmt"
	"slices"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func main() {
	grid, _ := utils.ReadGridRune("cmd/day8/input.txt")

	fmt.Printf("Part 1: %v\n", partOne(grid))
}

func partOne(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])

	antennas := getAntennas(grid, rows, cols)
	antinodes := getAntinodes(antennas)

	count := 0
	for _, antinode := range antinodes {
		r, c := antinode[0], antinode[1]
		if r >= 0 && r < rows && c >= 0 && c < cols {
			count++
		}
	}

	displayMap(grid, rows, cols, antinodes)

	return count
}

func displayMap(grid [][]rune, rows, cols int, antinodes [][2]int) {
	for r := range rows {
		for c := range cols {
			for _, a := range antinodes {
				if r == a[0] && c == a[1] {
					fmt.Printf("#")
				}
			}
			fmt.Printf("%v", string(grid[r][c]))
		}
		fmt.Println()
	}
}

func getAntennas(grid [][]rune, rows, cols int) map[rune][][2]int {
	antennas := map[rune][][2]int{}

	for r := range rows {
		for c := range cols {
			if grid[r][c] != '.' {
				antennas[grid[r][c]] = append(
					antennas[grid[r][c]],
					[2]int{r, c})
			}
		}
	}

	return antennas
}

func getAntinodes(antennas map[rune][][2]int) [][2]int {
	antinodes := [][2]int{}

	for _, array := range antennas {
		for i := 0; i < len(array); i++ {
			for j := i + 1; j < len(array); j++ {
				r1, c1 := array[i][0], array[i][1]
				r2, c2 := array[j][0], array[j][1]

				x := 2*r1 - r2
				y := 2*c1 - c2
				if !slices.Contains(antinodes, [2]int{x, y}) {
					antinodes = append(antinodes, [2]int{x, y})
				}
				x = 2*r2 - r1
				y = 2*c2 - c1
				if !slices.Contains(antinodes, [2]int{x, y}) {
					antinodes = append(antinodes, [2]int{x, y})
				}
			}
		}
	}
	return antinodes
}
