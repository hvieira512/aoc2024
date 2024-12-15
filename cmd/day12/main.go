package main

import (
	"fmt"
	"slices"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func main() {
	utils.RenderDayHeader(12)
	grid, _ := utils.Runes("cmd/day12/input.txt")
	n := len(grid)
	regions := getRegions(grid, n)

	fmt.Printf("Part 1: %v\n", partOne(regions))
	fmt.Printf("Part 2: %v\n", partTwo(regions))
}

func partOne(regions [][][2]int) int {
	result := 0

	for _, region := range regions {
		result += len(region) * getRegionPerimeter(region)
	}

	return result
}

func getRegionPerimeter(region [][2]int) int {
	perimeter := 0
	directions := [][2]int{
		{0, -1}, {0, 1},
		{1, 0}, {-1, 0},
	}

	for _, cell := range region {
		perimeter += 4
		for _, dir := range directions {
			nr, nc := cell[0]+dir[0], cell[1]+dir[1]
			if slices.Contains(region, [2]int{nr, nc}) {
				perimeter--
			}
		}
	}

	return perimeter
}

func getRegions(grid [][]rune, n int) [][][2]int {
	regions := [][][2]int{}
	seen := map[[2]int]bool{}
	directions := [][2]int{
		{-1, 0}, {1, 0},
		{0, -1}, {0, 1},
	}

	for r := range n {
		for c := range n {
			if seen[[2]int{r, c}] {
				continue
			}

			seen[[2]int{r, c}] = true
			region := [][2]int{{r, c}}
			queue := [][2]int{{r, c}}
			letter := grid[r][c]

			for len(queue) > 0 {
				cr, cc := queue[0][0], queue[0][1]
				queue = queue[1:]

				for _, dir := range directions {
					nr, nc := cr+dir[0], cc+dir[1]

					if !utils.InBoundsGrid(nr, nc, n) {
						continue
					}

					if grid[nr][nc] != letter {
						continue
					}

					if seen[[2]int{nr, nc}] {
						continue
					}

					seen[[2]int{nr, nc}] = true
					region = append(region, [2]int{nr, nc})
					queue = append(queue, [2]int{nr, nc})
				}
			}
			regions = append(regions, region)
		}
	}

	return regions
}

func partTwo(regions [][][2]int) int {
	result := 0

	for _, region := range regions {
		result += len(region) * sides(region)
	}

	return result
}

func sides(region [][2]int) int {
	cornerCandidates := map[[2]float64]struct{}{}

	// convert region to float64 for later calcs
	reg := [][2]float64{}
	for _, cell := range region {
		cr, cc := float64(cell[0]), float64(cell[1])
		reg = append(reg, [2]float64{cr, cc})
	}

	for _, r := range reg {
		for _, offset := range [][2]float64{
			{-0.5, -0.5}, {0.5, -0.5}, {0.5, 0.5}, {-0.5, 0.5},
		} {
			cornerCandidates[[2]float64{r[0] + offset[0], r[1] + offset[1]}] = struct{}{}
		}
	}

	corners := 0
	for corner := range cornerCandidates {
		config := [4]bool{
			slices.Contains(reg, [2]float64{corner[0] - 0.5, corner[1] - 0.5}),
			slices.Contains(reg, [2]float64{corner[0] + 0.5, corner[1] - 0.5}),
			slices.Contains(reg, [2]float64{corner[0] + 0.5, corner[1] + 0.5}),
			slices.Contains(reg, [2]float64{corner[0] - 0.5, corner[1] + 0.5}),
		}

		number := 0
		for _, c := range config {
			if c {
				number++
			}
		}

		switch number {
		case 1:
			corners++
		case 2:
			if (config[0] && !config[1] && config[2] && !config[3]) || (!config[0] && config[1] && !config[2] && config[3]) {
				corners += 2
			}
		case 3:
			corners++
		}
	}

	return corners
}
