package main

import (
	"fmt"
	"slices"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

const (
	Filename = "example"
)

func main() {
	utils.RenderDayHeader(15)
	lines, _ := utils.Strings("cmd/day15/" + Filename + ".txt")

	warehouse, moves, robot := parseInput(lines)
	fmt.Printf("Part 1: %v\n", partOne(warehouse, moves, robot))

	lines = scaleMap(lines)
	warehouse, _, robot = parseInput(lines)
	fmt.Printf("Part 2: %v\n", partTwo(warehouse, moves, robot))
}

func scaleMap(lines []string) []string {
	var scaled []string
	tileExpansion := map[rune]string{
		'#': "##",
		'O': "[]",
		'.': "..",
		'@': "@.",
	}

	for _, row := range lines {
		scaledRow := ""
		for _, tile := range row {
			if expanded, ok := tileExpansion[tile]; ok {
				scaledRow += expanded
			} else {
				scaledRow += string(tile)
			}
		}
		scaled = append(scaled, scaledRow)
	}

	return scaled
}

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
		start := slices.Index(lines, "") + 1

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

func partOne(warehouse map[rune][][2]int, moves string, robot [2]int) int {
	for _, move := range moves {
		next := getNextDirection(move, robot)

		if slices.Contains(warehouse['#'], next) {
			continue
		}

		if !slices.Contains(warehouse['O'], next) {
			robot = next
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

		for char, positions := range warehouse {
			for i, pos := range positions {
				if pos == next {
					warehouse[char][i] = free
					break
				}
			}
		}
		robot = next
	}

	sum := 0
	for _, box := range warehouse['O'] {
		sum += (100 * box[0]) + box[1]
	}

	return sum
}

func renderMap(warehouse map[rune][][2]int, robot [2]int) {
	length := 0
	for _, positions := range warehouse {
		for _, pos := range positions {
			if pos[0] == 0 {
				length++
			}
		}
	}
	for r := 0; r < length/2; r++ {
		for c := 0; c < length; c++ {
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

func GetInBetweenPositions(start, end [2]int) (positions [][2]int) {
	if start[0] == end[0] {
		x := start[0]
		y1, y2 := start[1], end[1]
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1 + 1; y < y2; y++ {
			positions = append(positions, [2]int{x, y})
		}
	} else if start[1] == end[1] {
		y := start[1]
		x1, x2 := start[0], end[0]

		if x1 > x2 {
			x1, x2 = x2, x1
		}

		for x := x1 + 1; x < x2; x++ {
			positions = append(positions, [2]int{x, y})
		}
	}
	return positions
}

func partTwo(warehouse map[rune][][2]int, moves string, robot [2]int) int {
	renderMap(warehouse, robot)
	for _, move := range moves {
		next := getNextDirection(move, robot)

		if slices.Contains(warehouse['#'], next) {
			fmt.Printf("Move %v:\n", string(move))
			renderMap(warehouse, robot)
			continue
		}

		if !slices.Contains(warehouse['['], next) && !slices.Contains(warehouse[']'], next) {
			robot = next
			fmt.Printf("Move %v:\n", string(move))
			renderMap(warehouse, robot)
			continue
		}

		canMove := func(next [2]int) [2]int {
			for {
				after := getNextDirection(move, next)
				if !slices.Contains(warehouse['#'], after) &&
					!slices.Contains(warehouse['['], after) &&
					!slices.Contains(warehouse[']'], after) {
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

		for _, curr := range GetInBetweenPositions(robot, free) {
			next := getNextDirection(move, curr)
			after := getNextDirection(move, next)
			for i, pos := range warehouse[']'] {
				if curr == pos {
					warehouse[']'][i] = next
				}
			}
			for i, poss := range warehouse['['] {
				if next == poss {
					warehouse['['][i] = after
				}
			}
		}
		robot = next
		fmt.Printf("Move %v:\n", string(move))
		renderMap(warehouse, robot)
	}

	renderMap(warehouse, robot)

	sum := 0
	for _, box := range warehouse['O'] {
		sum += (100 * box[0]) + box[1]
	}

	return sum
}
