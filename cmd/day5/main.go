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
			break
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

func isOrdered(update []int, rules [][2]int) (bool, int) {
	pagePositions := map[int]int{}
	for i, page := range update {
		pagePositions[page] = i
	}

	for _, rule := range rules {
		a, b := rule[0], rule[1]
		posA, existsA := pagePositions[a]
		posB, existsB := pagePositions[b]
		if existsA && existsB && posA > posB {
			return false, 0
		}
	}

	return true, update[len(update)/2]
}

func partOne(rules [][2]int, updates [][]int) int {
	total := 0

	for _, update := range updates {
		good, mid := isOrdered(update, rules)
		if good {
			total += mid
		}
	}

	return total
}

func sortUpdate(update []int, rules [][2]int) []int {
	for {
		isSorted := true
		for i := 0; i < len(update)-1; i++ {
			if containsRule(rules, update[i+1], update[i]) {
				update[i], update[i+1] = update[i+1], update[i]
				isSorted = false
			}
		}

		if isSorted {
			return update
		}
	}
}

func containsRule(rules [][2]int, a, b int) bool {
	for _, rule := range rules {
		if rule[0] == a && rule[1] == b {
			return true
		}
	}
	return false
}

func partTwo(rules [][2]int, updates [][]int) int {
	total := 0

	for _, update := range updates {
		good, _ := isOrdered(update, rules)
		if !good {
			// order it
			update := sortUpdate(update, rules)
			mid := update[len(update)/2]
			total += mid
		}
	}

	return total
}

func main() {
	utils.RenderDayHeader(5)
	lines, _ := utils.Strings("cmd/day5/input.txt")
	rules, updates := parseInput(lines)

	fmt.Printf("Part 1: %v\n", partOne(rules, updates))
	fmt.Printf("Part 2: %v\n", partTwo(rules, updates))
}
