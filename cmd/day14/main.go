package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

type Robot struct {
	pos [2]int
	vel [2]int
}

const (
	Width   = 101
	Height  = 103
	Seconds = 100
)

func parseRobots(lines []string) []Robot {
	var robots []Robot
	n := len(lines)

	for i := range n {
		robots = append(robots, parseRobot(lines[i]))
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
	lines, _ := utils.ReadLines("cmd/day14/input.txt")
	robots := parseRobots(lines)

	fmt.Printf("Part 1: %v\n", partOne(robots))
}

func partOne(robots []Robot) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	midX, midY := (Width-1)/2, (Height-1)/2

	for _, robot := range robots {
		robot.pos = moveRobot(robot)

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

	return q1 * q2 * q3 * q4
}

func moveRobot(robot Robot) [2]int {
	px, py := robot.pos[0], robot.pos[1]
	vx, vy := robot.vel[0], robot.vel[1]

	x := (px + vx*Seconds) % Width
	if x < 0 {
		x += Width
	}
	y := (py + vy*Seconds) % Height
	if y < 0 {
		y += Height
	}

	return [2]int{x, y}
}
