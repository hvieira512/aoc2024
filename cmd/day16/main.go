package main

import (
	"container/heap"
	"fmt"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

const (
	Filename = "input"
)

type State struct{ Cost, Row, Col, DR, DC int }
type PriorityQueue []*State

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Cost < pq[j].Cost }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*State)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func getChar(grid [][]rune) (int, int) {
	rows, cols := len(grid), len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'S' {
				return r, c
			}
		}
	}

	return -1, -1
}

func partOne(grid [][]rune) int {
	sr, sc := getChar(grid)
	rows, cols := len(grid), len(grid[0])

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &State{Cost: 0, Row: sr, Col: sc, DR: 0, DC: 1})
	seen := map[[4]int]bool{{sr, sc, 0, 1}: true}

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*State)
		cost, r, c, dr, dc := curr.Cost, curr.Row, curr.Col, curr.DR, curr.DC

		if grid[r][c] == 'E' {
			return cost
		}

		moves := []struct {
			NewCost, NR, NC, NDR, NDC int
		}{
			{cost + 1, r + dr, c + dc, dr, dc},
			{cost + 1000, r, c, dc, -dr},
			{cost + 1000, r, c, -dc, dr},
		}

		for _, move := range moves {
			nr, nc := move.NR, move.NC

			if nr < 0 || nc < 0 || nr >= rows || nc >= cols || grid[nr][nc] == '#' {
				continue
			}

			state := [4]int{nr, nc, move.NDR, move.NDC}
			if seen[state] {
				continue
			}
			seen[state] = true
			heap.Push(pq, &State{
				Cost: move.NewCost,
				Row:  nr,
				Col:  nc,
				DR:   move.NDR,
				DC:   move.NDC,
			})
		}
	}

	return -1
}

func partTwo(grid [][]rune) int {
	sr, sc := getChar(grid)
	rows, cols := len(grid), len(grid[0])

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &State{Cost: 0, Row: sr, Col: sc, DR: 0, DC: 1})

	lowestCost := make(map[[4]int]int)
	lowestCost[[4]int{sr, sc, 0, 1}] = 0

	backtrack := make(map[[4]int][][4]int)
	bestCost := int(^uint(0) >> 1) // max int value
	endStates := make(map[[4]int]bool)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*State)
		cost, r, c, dr, dc := curr.Cost, curr.Row, curr.Col, curr.DR, curr.DC

		if cost > lowestCost[[4]int{r, c, dr, dc}] {
			continue
		}
		if grid[r][c] == 'E' {
			if cost > bestCost {
				break
			}
			bestCost = cost
			endStates[[4]int{r, c, dr, dc}] = true
		}

		moves := [][5]int{
			{cost + 1, r + dr, c + dc, dr, dc},
			{cost + 1000, r, c, dc, -dr},
			{cost + 1000, r, c, -dc, dr},
		}

		for _, move := range moves {
			newCost, nr, nc, ndr, ndc := move[0], move[1], move[2], move[3], move[4]

			if nr < 0 || nc < 0 || nr >= rows || nc >= cols || grid[nr][nc] == '#' {
				continue
			}

			state := [4]int{nr, nc, ndr, ndc}
			lowest, exists := lowestCost[state]
			if !exists {
				lowest = int(^uint(0) >> 1) // max int value
			}

			if newCost > lowest {
				continue
			}
			if newCost < lowest {
				backtrack[state] = [][4]int{}
				lowestCost[state] = newCost
			}
			backtrack[state] = append(backtrack[state], [4]int{r, c, dr, dc})
			heap.Push(pq, &State{Cost: newCost, Row: nr, Col: nc, DR: ndr, DC: ndc})
		}
	}

	states := []([4]int){}
	seen := make(map[[4]int]bool)

	for state := range endStates {
		states = append(states, state)
		seen[state] = true
	}

	for len(states) > 0 {
		key := states[0]
		states = states[1:]

		for _, last := range backtrack[key] {
			if seen[last] {
				continue
			}
			seen[last] = true
			states = append(states, last)
		}
	}

	// Count unique grid cells in seen states
	uniqueCells := make(map[[2]int]bool)
	for state := range seen {
		uniqueCells[[2]int{state[0], state[1]}] = true
	}

	return len(uniqueCells)
}

func main() {
	grid, err := utils.Runes("cmd/day16/" + Filename + ".txt")
	if err != nil {
		fmt.Println("Error reading grid:", err)
		return
	}

	fmt.Printf("Part 1: %v\n", partOne(grid))
	fmt.Printf("Part 2: %v\n", partTwo(grid))
}
