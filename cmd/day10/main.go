package main

import (
	"container/list"
	"fmt"

	u "github.com/hvieira512/aoc2024/cmd/utils"
)

func getTrailHeads(grid [][]int, rows, cols int) [][2]int {
	trailheads := [][2]int{}

	for r := range rows {
		for c := range cols {
			if grid[r][c] == 0 {
				trailheads = append(trailheads, [2]int{r, c})
			}
		}
	}
	return trailheads
}

func score(grid [][]int, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])

	q := [][2]int{{r, c}}
	seen := map[[2]int]bool{{r, c}: true}
	summits := 0

	directions := [][2]int{
		{-1, 0}, {1, 0},
		{0, -1}, {0, 1},
	}

	// BFS, btw
	for len(q) > 0 {
		cr, cc := q[0][0], q[0][1]
		q = q[1:]

		for _, dir := range directions {
			nr, nc := cr+dir[0], cc+dir[1]
			if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
				continue
			}

			if grid[nr][nc] != grid[cr][cc]+1 {
				continue
			}

			if seen[[2]int{nr, nc}] {
				continue
			}
			seen[[2]int{nr, nc}] = true

			if grid[nr][nc] == 9 {
				summits++
			} else {
				q = append(q, [2]int{nr, nc})
			}
		}
	}

	return summits
}

func partOne(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	trailheads := getTrailHeads(grid, rows, cols)

	result := 0
	for _, start := range trailheads {
		result += score(grid, start[0], start[1])
	}

	return result
}

func main() {
	u.RenderDayHeader(10)
	grid, _ := u.ReadGridInt("cmd/day10/input.txt")

	fmt.Printf("Part 1: %v\n", partOne(grid))
	fmt.Printf("Part 2: %v\n", partTwo(grid))
}

func partTwo(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	trailheads := getTrailHeads(grid, rows, cols)

	result := 0
	for _, start := range trailheads {
		result += score2(grid, start[0], start[1])
	}
	return result
}

func score2(grid [][]int, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])

	q := list.New()
	q.PushBack([2]int{r, c})
	seen := make(map[[2]int]int)
	seen[[2]int{r, c}] = 1
	trails := 0

	for q.Len() > 0 {
		elem := q.Front()
		q.Remove(elem)
		current := elem.Value.([2]int)
		cr, cc := current[0], current[1]

		if grid[cr][cc] == 9 {
			trails += seen[[2]int{cr, cc}]
		}

		for _, dir := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			nr, nc := cr+dir[0], cc+dir[1]
			if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
				continue
			}
			if grid[nr][nc] != grid[cr][cc]+1 {
				continue
			}
			if _, exists := seen[[2]int{nr, nc}]; exists {
				seen[[2]int{nr, nc}] += seen[[2]int{cr, cc}]
				continue
			}
			seen[[2]int{nr, nc}] = seen[[2]int{cr, cc}]
			q.PushBack([2]int{nr, nc})
		}
	}

	return trails
}
