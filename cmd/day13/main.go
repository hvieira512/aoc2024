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
	lines, _ := utils.ReadLines("cmd/day13/input.txt")
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

	for r := 0; r < rows; r += 3 {
		if len(lines[r]) == 0 {
			r++
		}

		var machine ClawMachine
		for i := 0; i <= 1; i++ {
			buttonAux := strings.Split(lines[r+i][10:], ": ")[0]
			coordsAux := strings.Split(buttonAux, ", ")
			x, _ := strconv.ParseInt(strings.Split(coordsAux[0], "+")[1], 10, 64)
			y, _ := strconv.ParseInt(strings.Split(coordsAux[1], "+")[1], 10, 64)
			machine.Buttons[i] = [2]int64{x, y}
		}

		prizeAux := strings.Split(lines[r+2][7:], ": ")[0]
		coordsAux := strings.Split(prizeAux, ", ")
		x, _ := strconv.ParseInt(strings.Split(coordsAux[0], "=")[1], 10, 64)
		y, _ := strconv.ParseInt(strings.Split(coordsAux[1], "=")[1], 10, 64)
		machine.Prize = [2]int64{x, y}

		machines = append(machines, machine)
	}

	return machines
}

func cramer(A [2][2]int64, B [2]int64) int64 {
	aux := A[1][0]
	A[1][0] = A[0][1]
	A[0][1] = aux

	detA := (A[0][0] * A[1][1]) - (A[0][1] * A[1][0])
	if detA == 0 {
		return 0
	}

	x := (B[0]*A[1][1] - B[1]*A[0][1]) / detA
	y := (A[0][0]*B[1] - B[0]*A[1][0]) / detA

	if x*A[0][0]+y*A[0][1] != B[0] || x*A[1][0]+y*A[1][1] != B[1] {
		return 0
	}

	return x*3 + y
}
