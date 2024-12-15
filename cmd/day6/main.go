package main

import (
	"fmt"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

type Coordinate struct {
	x, y int
}

type Guard struct {
	pos       Coordinate
	direction string
}

func getGuard(guardMap []string) Guard {
	allGuardChars := "^>v<"

	for i := range guardMap {
		guardFound := strings.IndexAny(guardMap[i], allGuardChars)
		direction := ""
		if guardFound != -1 {
			switch string(guardMap[i][guardFound]) {
			case "^":
				direction = "up"
			case ">":
				direction = "right"
			case "v":
				direction = "down"
			case "<":
				direction = "left"
			default:
				panic("Invalid Character for the Guard")
			}
			return Guard{
				pos:       Coordinate{x: i, y: guardFound},
				direction: direction,
			}
		}
	}

	return Guard{
		pos:       Coordinate{x: -1, y: -1},
		direction: "",
	}
}
func containsPos(allDistinctPos []Coordinate, guardPos Coordinate) bool {
	for _, dP := range allDistinctPos {
		if dP == guardPos {
			return true
		}
	}
	return false

}

func partOne(guardMap []string) int {
	allVisitedPos := []Coordinate{}

	rows := len(guardMap)
	cols := len(guardMap[0])

	guard := getGuard(guardMap)
	if guard.pos.x == -1 && guard.pos.y == -1 {
		panic("Guard wasn't found!")
	}

	directionOffsets := map[string]Coordinate{
		"up":    {-1, 0},
		"down":  {1, 0},
		"right": {0, 1},
		"left":  {0, -1},
	}

	rotateRight := map[string]string{
		"up":    "right",
		"right": "down",
		"down":  "left",
		"left":  "up",
	}

	for {
		offset := directionOffsets[guard.direction]
		nextDirection := Coordinate{x: guard.pos.x + offset.x, y: guard.pos.y + offset.y}

		inBounds := nextDirection.x >= 0 && nextDirection.y >= 0 &&
			nextDirection.x < rows && nextDirection.y < cols

		if !inBounds {
			break
		}

		nx := nextDirection.x
		ny := nextDirection.y

		if guardMap[nx][ny] == '#' {
			guard.direction = rotateRight[guard.direction]
			continue
		}

		guard.pos = nextDirection

		if !containsPos(allVisitedPos, guard.pos) {
			allVisitedPos = append(allVisitedPos, guard.pos)
		}
	}

	return len(allVisitedPos)
}

func partTwo(guardMap []string) int {
	rows := len(guardMap)
	cols := len(guardMap[0])
	loop := 0

	possibilities := []Coordinate{}
	for i := range guardMap {
		for j := range guardMap[i] {
			if string(guardMap[i][j]) != "#" {
				newPossibility := Coordinate{x: i, y: j}
				possibilities = append(possibilities, newPossibility)
			}
		}
	}

	directionOffsets := map[string]Coordinate{
		"up":    {-1, 0},
		"down":  {1, 0},
		"right": {0, 1},
		"left":  {0, -1},
	}

	rotateRight := map[string]string{
		"up":    "right",
		"right": "down",
		"down":  "left",
		"left":  "up",
	}

	guard := getGuard(guardMap)
	if guard.pos.x == -1 && guard.pos.y == -1 {
		panic("Guard wasn't found!")
	}

	origMap := make([]string, len(guardMap))
	copy(origMap, guardMap)
	initialGuard := guard

	for _, p := range possibilities {
		guard := initialGuard
		row := []rune(guardMap[p.x])
		row[p.y] = '#'
		guardMap[p.x] = string(row)

		allVisitedPos := []Coordinate{}
		visited := make(map[Coordinate]string)
		for {
			offset := directionOffsets[guard.direction]
			nextDirection := Coordinate{x: guard.pos.x + offset.x, y: guard.pos.y + offset.y}

			inBounds := nextDirection.x >= 0 && nextDirection.y >= 0 &&
				nextDirection.x < rows && nextDirection.y < cols

			if !inBounds {
				break
			}

			nx := nextDirection.x
			ny := nextDirection.y

			if guardMap[nx][ny] == '#' {
				guard.direction = rotateRight[guard.direction]
				continue
			}

			guard.pos = nextDirection

			allVisitedPos = append(allVisitedPos, guard.pos)
			if visited[guard.pos] == guard.direction {
				loop++
				break
			}

			visited[guard.pos] = guard.direction
		}
		copy(guardMap, origMap)
	}

	return loop
}

func main() {
	guardMap, _ := utils.Strings("cmd/day6/input.txt")

	fmt.Printf("Part 1: %v\n", partOne(guardMap))
	fmt.Printf("Part 2: %v\n", partTwo(guardMap))
}
