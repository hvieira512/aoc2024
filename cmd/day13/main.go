package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

type ClawMachine struct {
	Buttons [2][2]int64
	Prize   [2]int64
}

func main() {
	utils.RenderDayHeader(13)
	lines, _ := utils.Strings("cmd/day13/input.txt")
	machines := getMachines(lines)

	fmt.Printf("Part 1: %v\n", solve(machines, 1))
	fmt.Printf("Part 2: %v\n", solve(machines, 2))
}

func solve(machines []ClawMachine, part int) int64 {
	result := int64(0)
	for _, machine := range machines {
		if part == 2 {
			for i := range len(machine.Prize) {
				machine.Prize[i] += 10000000000000
			}
		}
		result += cramer(machine.Buttons, machine.Prize)
	}
	return result
}

func getMachines(lines []string) []ClawMachine {
	var machines []ClawMachine
	rows := len(lines)

	for i := 0; i < rows; i += 3 {
		if len(lines[i]) == 0 {
			i++
		}
		machines = append(machines, parseMachine(lines[i:i+3]))
	}

	return machines
}

func parseMachine(lines []string) ClawMachine {
	var machine ClawMachine
	for i := 0; i < 2; i++ {
		coords := parseCoordinates(lines[i][10:], "+")
		machine.Buttons[i] = [2]int64{coords[0], coords[1]}
	}
	coords := parseCoordinates(lines[2][7:], "=")
	machine.Prize = [2]int64{coords[0], coords[1]}
	return machine
}

func parseCoordinates(input, sep string) [2]int64 {
	parts := strings.Split(input, ", ")
	coords := [2]int64{}
	for i, part := range parts {
		splitPart := strings.Split(part, sep)
		coords[i], _ = strconv.ParseInt(splitPart[len(splitPart)-1], 10, 64)
	}
	return coords
}

func cramer(A [2][2]int64, B [2]int64) int64 {
	A[0][1], A[1][0] = A[1][0], A[0][1]

	detA := (A[0][0] * A[1][1]) - (A[0][1] * A[1][0])
	if detA == 0 {
		return 0
	}

	x := (B[0]*A[1][1] - B[1]*A[0][1]) / detA
	y := (A[0][0]*B[1] - B[0]*A[1][0]) / detA

	if !verifySolution(A, B, x, y) {
		return 0
	}

	return x*3 + y
}

func verifySolution(A [2][2]int64, B [2]int64, x, y int64) bool {
	return x*A[0][0]+y*A[0][1] == B[0] && x*A[1][0]+y*A[1][1] == B[1]
}
