package main

import (
	"fmt"
	"log"
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
	for range report {
		op := getOperator(report, 0, 1)
		fmt.Println(op)

		for i := 1; i < len(report); i++ {
			fmt.Println(op, report[i], report[i+1])
			switch op {
			case "<":
				if report[i] > report[i+1] {
					return false
				}
			case ">":
				if report[i] < report[i+1] {
					return false
				}
			}
		}
	}

	return true
}

func getOperator(report []int, firstIndex int, secondIndex int) string {
	op := ""
	if report[firstIndex] > report[secondIndex] {
		op = "<"
	} else {
		op = ">"
	}
	return op
}

func partOne(lines []string) int {
	// store the reports with their levels
	reports := getReports(lines)

	count := 0
	for _, report := range reports {
		if isReportSafe(report) {
			fmt.Printf("%v is safe\n", report)
			count++
		} else {
			fmt.Printf("%v is unsafe\n", report)
		}
	}

	return count
}

func main() {
	u.RenderDayHeader(2)
	lines, err := u.ReadLines("cmd/day2/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %v\n", partOne(lines))
	// fmt.Printf("Part 2: %v\n", partTwo(lines))
}

func partTwo(lines []string) int {
	panic("unimplemented")
}
