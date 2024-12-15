package main

import (
	"fmt"
	"slices"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func main() {
	utils.RenderDayHeader(8)
	grid, _ := utils.Runes("cmd/day8/input.txt")

	fmt.Printf("Part 1: %v\n", partOne(grid))
	fmt.Printf("Part 2: %v\n", partTwo(grid))
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

	return count
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

func partTwo(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])

	antennas := getAntennas(grid, rows, cols)
	antinodes := getAntinodesV2(antennas, rows, cols)

	return len(antinodes)
}

func getAntinodesV2(antennas map[rune][][2]int, rows, cols int) map[[2]int]struct{} {
	antinodes := make(map[[2]int]struct{})

	for _, array := range antennas {
		for i := 0; i < len(array); i++ {
			for j := 0; j < len(array); j++ {
				if i == j {
					continue
				}
				r1, c1 := array[i][0], array[i][1]
				r2, c2 := array[j][0], array[j][1]
				dr := r2 - r1
				dc := c2 - c1
				r := r1
				c := c1
				for r >= 0 && r < rows && c >= 0 && c < cols {
					antinodes[[2]int{r, c}] = struct{}{}
					r += dr
					c += dc
				}
			}
		}
	}
	return antinodes
}
