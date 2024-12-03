package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	u "github.com/hvieira512/aoc2024/cmd/utils"
)

func partOne(lines []string) int {
	result := 0

	for _, line := range lines {
		// sanitize string with regex
		line = regexPartOne(line)

		mults := strings.Split(line, "mul")
		for _, mult := range mults {
			if len(mult) == 0 {
				continue
			}
			firstValue := getMulValue(mult, '(', ',')
			secondValue := getMulValue(mult, ',', ')')

			result += firstValue * secondValue
		}
	}

	return result
}

func getMulValue(input string, charStart, charEnd rune) int {
	startIdx := strings.Index(input, string(charStart))
	endIdx := strings.Index(input, string(charEnd))

	if startIdx == -1 || endIdx == -1 || startIdx >= endIdx {
		return -9999
	}

	strValue := input[startIdx+1 : endIdx]

	value, err := strconv.Atoi(strValue)
	if err != nil {
		log.Fatal(err)
	}

	return value
}

func regexPartOne(input string) string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)
	return strings.Join(matches, "")
}

func partTwo(lines []string) int {
	result := 0

	for _, line := range lines {
		// line = "mul(3,4)mul(4,5)"
		// remove stuff between each don't() and do()
		line = justDoOrDontIt(line)

		// sanitize string with regex
		line = regexPartOne(line)
		fmt.Println(line)

		mults := strings.Split(line, "mul")
		for _, mult := range mults {
			if len(mult) == 0 {
				continue
			}
			firstValue := getMulValue(mult, '(', ',')
			secondValue := getMulValue(mult, ',', ')')

			result += firstValue * secondValue
		}
	}

	return result
}

func justDoOrDontIt(input string) string {
	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	return re.ReplaceAllString(input, "")
}

func main() {
	u.RenderDayHeader(3)
	lines, _ := u.ReadLines("cmd/day3/input.txt")

	fmt.Printf("Part 1: %v\n", partOne(lines))
	fmt.Printf("Part 2: %v\n", partTwo(lines))
}
