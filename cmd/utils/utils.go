package utils

import (
	"bufio"
	"fmt"
	"os"
)

func RenderDayHeader(day int) {
	fmt.Println("-----------------------------")
	fmt.Println("---- Advent of Code 2024 ----")
	fmt.Println("-----------------------------")
	fmt.Printf("----        Day %v        ----\n", day)
	fmt.Println("-----------------------------")
}

func ReadLines(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
