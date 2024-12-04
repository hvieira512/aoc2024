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
		xRow := xCoord.x
		xCol := xCoord.y

		adjLetterCoords := []Coordinate{
			{x: xRow - 1, y: xCol - 1},
			{x: xRow - 1, y: xCol},
			{x: xRow - 1, y: xCol + 1},
			{x: xRow + 1, y: xCol - 1},
			{x: xRow + 1, y: xCol},
			{x: xRow + 1, y: xCol + 1},
			{x: xRow, y: xCol - 1},
			{x: xRow, y: xCol + 1},
		}

		nextLetters := []string{"M", "A", "S"}
		for _, adjLetter := range adjLetterCoords {
			if isCoordOutOfBounds(adjLetter, lines) {
				continue
			}

			x, y := adjLetter.x, adjLetter.y
			for _, nextLetter := range nextLetters {
				if string(lines[x][y]) == nextLetter {
					// save the direction of the next letter
					nextLetterCoord := Coordinate{
						x: x - xRow,
						y: y - xCol,
					}

					fmt.Println(nextLetterCoord)
				}
			}
		}
	}

	return result
}

func isCoordOutOfBounds(coord Coordinate, lines []string) bool {
	x, y := coord.x, coord.y
	xLen, yLen := len(lines)-1, len(lines[0])-1

	return x < 0 || x > xLen || y < 0 || y > yLen
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
