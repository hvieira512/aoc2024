package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

type ClawMachine struct {
	Buttons [2][2]int
	Prize   [2]int
}

func main() {
	utils.RenderDayHeader(13)
	lines, err := utils.ReadLines("cmd/day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	machines := getMachines(lines)

	fmt.Printf("Part 1: %v\n", partOne(machines))
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
			x, _ := strconv.Atoi(strings.Split(coordsAux[0], "+")[1])
			y, _ := strconv.Atoi(strings.Split(coordsAux[1], "+")[1])
			machine.Buttons[i] = [2]int{x, y}
		}

		// Get prize
		prizeAux := strings.Split(lines[r+2][7:], ": ")[0]
		coordsAux := strings.Split(prizeAux, ", ")
		x, _ := strconv.Atoi(strings.Split(coordsAux[0], "=")[1])
		y, _ := strconv.Atoi(strings.Split(coordsAux[1], "=")[1])
		machine.Prize = [2]int{x, y}

		machines = append(machines, machine)
	}

	return machines
}

func partOne(machines []ClawMachine) int {
	result := 0
	for _, machine := range machines {
		result += findCheapest(machine.Buttons, machine.Prize)
	}
	return result
}

func findCheapest(A [2][2]int, B [2]int) int {
	cheapest := [2]int{}
	maxPresses, minTokens := 100, -1

	for x := 1; x <= maxPresses; x++ {
		for y := 1; y <= maxPresses; y++ {
			if x*A[0][0]+y*A[1][0] == B[0] && x*A[0][1]+y*A[1][1] == B[1] {
				tokens := x*3 + y
				if minTokens == -1 || (tokens < minTokens) {
					minTokens = x*3 + y
					cheapest = [2]int{x, y}
				}
			}
		}
	}

	return cheapest[0]*3 + cheapest[1]
}
