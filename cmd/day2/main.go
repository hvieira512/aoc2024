package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	u "github.com/hvieira512/aoc2024/cmd/utils"
)

func getReports(lines []string) [][]int {
	reports := [][]int{}
	for _, line := range lines {
		lineStr := strings.Fields(line)
		levels := []int{}
		for _, level := range lineStr {
			level, _ := strconv.Atoi(level)
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}
	return reports
}

func isReportSafe(report []int) bool {
	op := isIncreasingOrDecreasing(report, 0, 1)

	for i := 0; i < len(report); i++ {
		if i+1 == len(report) {
			break
		}

		diff := math.Abs(float64(report[i] - report[i+1]))
		if diff > 3 || diff == 0 {
			return false
		}

		switch op {
		case "increment":
			if report[i] > report[i+1] {
				return false
			}
		case "decrement":
			if report[i] < report[i+1] {
				return false
			}
		}

	}
	return true
}

func isIncreasingOrDecreasing(report []int, a int, b int) string {
	op := ""

	if report[a] > report[b] {
		op = "decrement"
	} else {
		op = "increment"
	}

	return op
}

func partOne(lines []string) int {
	reports := getReports(lines)

	count := 0
	for _, report := range reports {
		if isReportSafe(report) {
			count++
		}
	}

	return count
}

func partTwo(lines []string) int {
	reports := getReports(lines)

	count := 0
	for _, report := range reports {
		if isReportSafeV2(report) {
			count++
		}
	}
	return count
}

func isReportSafeV2(report []int) bool {
	if !isReportSafe(report) {
		for i := range report {
			newAttempt := append([]int{}, report[:i]...)
			newAttempt = append(newAttempt, report[i+1:]...)

			if isReportSafe(newAttempt) {
				return true
			}
		}
		return false
	}
	return true
}

func main() {
	u.RenderDayHeader(2)
	lines, err := u.Strings("cmd/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %v\n", partOne(lines))
	fmt.Printf("Part 2: %v\n", partTwo(lines))
}
