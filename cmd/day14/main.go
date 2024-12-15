package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

type Robot struct {
	pos [2]int
	vel [2]int
}

const (
	Width    = 101
	Height   = 103
	Seconds  = 100
	Filename = "input"
)

var midX int = (Width - 1) / 2
var midY int = (Height - 1) / 2

func parseRobots(lines []string) []Robot {
	var robots []Robot
	for _, line := range lines {
		robots = append(robots, parseRobot(line))
	}
	return robots
}

func parseRobot(line string) Robot {
	var robot Robot
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(line, -1)

	getNum := func(idx int) int {
		num, _ := strconv.Atoi(matches[idx])
		return num
	}

	robot.pos = [2]int{getNum(0), getNum(1)}
	robot.vel = [2]int{getNum(2), getNum(3)}

	return robot
}

func main() {
	utils.RenderDayHeader(14)
	lines, _ := utils.ReadLines("cmd/day14/" + Filename + ".txt")
	robots := parseRobots(lines)

	fmt.Printf("Part 1: %v\n", partOne(robots))
	fmt.Printf("Part 2: %v\n", partTwo(robots))
}

func moveRobot(robot Robot, seconds int) [2]int {
	x := (robot.pos[0] + robot.vel[0]*seconds) % Width
	if x < 0 {
		x += Width
	}
	y := (robot.pos[1] + robot.vel[1]*seconds) % Height
	if y < 0 {
		y += Height
	}

	return [2]int{x, y}
}

func countRobotsInQuads(robots []Robot, seconds int) (int, int, int, int) {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, robot := range robots {
		robot.pos = moveRobot(robot, seconds)
		switch {
		case robot.pos[0] > midX && robot.pos[1] > midY:
			q1++
		case robot.pos[0] > midX && robot.pos[1] < midY:
			q2++
		case robot.pos[0] < midX && robot.pos[1] < midY:
			q3++
		case robot.pos[0] < midX && robot.pos[1] > midY:
			q4++
		}
	}

	return q1, q2, q3, q4
}

func partOne(robots []Robot) int {
	q1, q2, q3, q4 := countRobotsInQuads(robots, Seconds)

	return q1 * q2 * q3 * q4
}

func partTwo(robots []Robot) int {
	minSF := math.Inf(1)
	bestIteration := -1

	for second := range Width * Height {
		q1, q2, q3, q4 := countRobotsInQuads(robots, second)

		sf := float64(q1 * q2 * q3 * q4)
		if sf < minSF {
			minSF = sf
			bestIteration = second
		}
	}

	displayGrid(robots, bestIteration)

	return bestIteration
}

func displayGrid(robots []Robot, seconds int) {
	grid := make([][]string, Height)
	for i := range grid {
		grid[i] = make([]string, Width)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for _, robot := range robots {
		robot.pos = moveRobot(robot, seconds)
		x, y := robot.pos[0], robot.pos[1]
		grid[y][x] = "#"
	}

	for _, row := range grid {
		fmt.Println(row)
	}
}
