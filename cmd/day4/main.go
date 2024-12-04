package main

import (
	"fmt"

	u "github.com/hvieira512/aoc2024/cmd/utils"
)

type Coordinate struct {
	x int
	y int
}

func partOne(lines []string) (result int) {
	allXCoords := getAllCharPos(lines, "X")

	for _, xCoord := range allXCoords {
		x := xCoord.x
		y := xCoord.y

		directions := []Coordinate{
			{x: x - 1, y: y - 1},
			{x: x - 1, y: y},
			{x: x - 1, y: y + 1},
			{x: x + 1, y: y - 1},
			{x: x + 1, y: y},
			{x: x + 1, y: y + 1},
			{x: x, y: y - 1},
			{x: x, y: y + 1},
		}

		if foundXmas(directions, lines, xCoord) {
			result++
		}
	}

	return result
}

func foundXmas(directions []Coordinate, lines []string, xCoords Coordinate) bool {
	for _, direction := range directions {
		mCoords := Coordinate{x: direction.x, y: direction.y}
		// fmt.Printf("Checking M at: %v\n", mCoords)
		if isCoordOutOfBounds(mCoords, lines) || string(lines[mCoords.x][mCoords.y]) != "M" {
			continue
		}

		aCoords := getNextCoords(mCoords, xCoords, lines, "A")
		if aCoords.x == -1 || aCoords.y == -1 {
			continue
		}

		sCoords := getNextCoords(aCoords, mCoords, lines, "S")
		if sCoords.x == -1 || sCoords.y == -1 {
			continue
		}

		fmt.Printf("%v[%v,%v], ", string(lines[xCoords.x][xCoords.y]), xCoords.x, xCoords.y)
		fmt.Printf("%v[%v,%v], ", string(lines[mCoords.x][mCoords.y]), mCoords.x, mCoords.y)
		fmt.Printf("%v[%v,%v], ", string(lines[aCoords.x][aCoords.y]), aCoords.x, aCoords.y)
		fmt.Printf("%v[%v,%v]\n", string(lines[sCoords.x][sCoords.y]), sCoords.x, sCoords.y)
		return true
	}
	return false
}

func getNextCoords(current, previous Coordinate, lines []string, target string) Coordinate {
	dx := current.x - previous.x
	dy := current.y - previous.y
	next := Coordinate{x: current.x + dx, y: current.y + dy}

	if isCoordOutOfBounds(next, lines) || string(lines[next.x][next.y]) != target {
		return Coordinate{x: -1, y: -1}
	}
	return next
}

func isCoordOutOfBounds(coord Coordinate, lines []string) bool {
	return coord.x < 0 || coord.x >= len(lines) || coord.y < 0 || coord.y >= len(lines[coord.x])
}

func getAllCharPos(lines []string, char string) []Coordinate {
	coordinates := []Coordinate{}

	for i := range lines {
		for j := range lines[i] {
			if string(lines[i][j]) == char {
				newCoordinate := Coordinate{x: i, y: j}
				coordinates = append(coordinates, newCoordinate)
			}
		}
	}

	return coordinates
}

func partTwo(lines []string) (result int) {
	for _, line := range lines {
		fmt.Println(line)
	}

	return result
}

func main() {
	u.RenderDayHeader(4)
	lines, _ := u.ReadLines("cmd/day4/example.txt")

	fmt.Printf("Part 1: %v\n", partOne(lines))
	// fmt.Printf("Part 2: %v\n", partTwo(lines))
}
