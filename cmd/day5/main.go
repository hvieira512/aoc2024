package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

func parseInput(lines []string) ([][2]int, [][]int) {
	rules := [][2]int{}
	updates := [][]int{}

	idx := -1
	for i := range lines {
		if len(lines[i]) == 0 {
			idx = i
		}
	}
	for i := 0; i < idx; i++ {
		pagesRuleStr := strings.Split(lines[i], "|")
		fPageRule, _ := strconv.Atoi(pagesRuleStr[0])
		lPageRule, _ := strconv.Atoi(pagesRuleStr[1])
		rules = append(rules, [2]int{fPageRule, lPageRule})
	}
	for i := idx + 1; i < len(lines); i++ {
		update := []int{}
		pagesUpdateStr := strings.Split(lines[i], ",")
		for _, pageUpdateStr := range pagesUpdateStr {
			page, _ := strconv.Atoi(pageUpdateStr)
			update = append(update, page)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

func partOne(lines []string) int {
	total := 0

	rules, updates := parseInput(lines)

	return total
}

func main() {
	utils.RenderDayHeader(5)
	lines, _ := utils.ReadLines("cmd/day5/example.txt")

	fmt.Printf("Part 1: %v\n", partOne(lines))
}
