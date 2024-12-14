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

var rows int = 11
var cols int = 7
var seconds int = 100

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
	lines, _ := utils.ReadLines("cmd/day14/example.txt")
	robots := parseRobots(lines)

	fmt.Printf("Part 1: %v\n", partOne(robots))
}

func partOne(robots []Robot) int {
	midX, midY := rows/2, cols/2
	quadrants := map[int][][2]int{
		1: {}, 2: {}, 3: {}, 4: {},
	}

	for _, robot := range robots {
		robot.pos = moveRobot(robot)
		robot.pos = [2]int{robot.pos[1], robot.pos[0]}

		var quad int
		switch {
		case robot.pos[0] > midX && robot.pos[1] > midY:
			quad = 1
		case robot.pos[0] > midX && robot.pos[1] < midY:
			quad = 2
		case robot.pos[0] < midX && robot.pos[1] < midY:
			quad = 3
		case robot.pos[0] < midX && robot.pos[1] > midY:
			quad = 4
		default:
			continue
		}
		quadrants[quad] = append(quadrants[quad], robot.pos)
	}

	return len(quadrants[1]) * len(quadrants[2]) * len(quadrants[3]) * len(quadrants[4])
}

func moveRobot(robot Robot) [2]int {
	px, py, vx, vy := robot.pos[0], robot.pos[1], robot.vel[0], robot.vel[1]
	return [2]int{
		((px+vx*seconds)%rows + rows) % rows,
		((py+vy*seconds)%cols + cols) % cols,
	}
}
