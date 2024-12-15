package main

import (
	"fmt"
	"slices"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

const (
	Filename = "example"
)

func parseInput(lines []string) (map[rune][][2]int, string, [2]int) {
	getWarehouse := func() (map[rune][][2]int, [2]int) {
		warehouse := make(map[rune][][2]int)
		var robot [2]int
		start := 0
		end := slices.Index(lines, "")

		for i := start; i < end; i++ {
			for j := 0; j < len(lines[i]); j++ {
				char := rune(lines[i][j])
				if char == '@' {
					robot = [2]int{i, j}
				} else if char != '.' {
					warehouse[char] = append(warehouse[char], [2]int{i, j})
				}
			}
		}
		return warehouse, robot
	}

	getMoves := func() string {
		var moves string
		start := slices.Index(lines, "")

		for i := start; i < len(lines); i++ {
			moves += lines[i]
		}
		return moves
	}

	warehouse, robot := getWarehouse()
	moves := getMoves()

	return warehouse, moves, robot
}

func getNextDirection(move rune, robot [2]int) [2]int {
	nx, ny := -1, -1
	switch move {
	case '^':
		nx, ny = robot[0]-1, robot[1]
	case '>':
		nx, ny = robot[0], robot[1]+1
	case 'v':
		nx, ny = robot[0]+1, robot[1]
	case '<':
		nx, ny = robot[0], robot[1]-1
	}
	return [2]int{nx, ny}
}

func main() {
	utils.RenderDayHeader(15)
	lines, _ := utils.Strings("cmd/day15/" + Filename + ".txt")
	warehouse, moves, robot := parseInput(lines)

	fmt.Printf("Part 1: %v\n", partOne(warehouse, moves, robot))
	// fmt.Printf("Part 2: %v\n", partTwo(lines))
}

func partOne(warehouse map[rune][][2]int, moves string, robot [2]int) int {
	for _, move := range moves {
		next := getNextDirection(move, robot)

		if slices.Contains(warehouse['#'], next) {
			renderMap(warehouse, move, robot)
			continue
		}

		if !slices.Contains(warehouse['O'], next) {
			robot = next
			renderMap(warehouse, move, robot)
			continue
		}

		canMove := func(next [2]int) [2]int {
			for {
				after := getNextDirection(move, next)
				if !slices.Contains(warehouse['#'], after) && !slices.Contains(warehouse['O'], after) {
					return after
				} else if slices.Contains(warehouse['#'], after) {
					return [2]int{-1, -1}
				}
				next = after
			}
		}

		free := canMove(next)
		if free == [2]int{-1, -1} {
			continue
		}

		// free space -> freePos
		// robot -> robot

		renderMap(warehouse, move, robot)
	}

	return 0
}

func renderMap(warehouse map[rune][][2]int, move rune, robot [2]int) {
	rows, cols := 8, 8

	fmt.Printf("Move %v:\n", string(move))
	for r := range rows {
		for c := range cols {
			curr := [2]int{r, c}

			if robot == curr {
				fmt.Printf("@")
				continue
			}

			found := false
			for char, positions := range warehouse {
				for _, pos := range positions {
					if pos == curr {
						fmt.Printf("%v", string(char))
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func partTwo(lines []string) int {
	panic("unimplemented")
}
