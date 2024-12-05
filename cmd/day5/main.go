package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func getContents(lines []string) ([][]int, [][]int) {
	getIdx := func() int {
		idx := -1
		for i := range lines {
			if lines[i] == "" {
				idx = i
			}
		}
		return idx
	}

	getValuesSep := func(a, b int, sep string) [][]int {
		values := [][]int{}
		for i := a; i < b; i++ {
			strAux := strings.Split(lines[i], sep)

			lineValue := []int{}
			for j := range strAux {
				nvalue, _ := strconv.Atoi(strAux[j])
				lineValue = append(lineValue, nvalue)
			}
			values = append(values, lineValue)
		}
		return values
	}

	idx := getIdx()
	rules := getValuesSep(0, idx, "|")
	updates := getValuesSep(idx, len(lines), ",")

	return rules, updates
}

func partOne(rules [][]int, updates [][]int) int {
	result := 0

	fmt.Println(rules)
	for _, update := range updates {
		fmt.Println(update)
	}

	return result
}

func main() {
	lines, _ := utils.ReadLines("cmd/day5/example.txt")
	rules, updates := getContents(lines)

	fmt.Printf("Part 1: %v\n", partOne(rules, updates))
}
