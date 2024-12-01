package utils

import (
	"fmt"
	"os"
	"strings"
)

func RenderDayHeader(day int) {
	fmt.Println("-----------------------------")
	fmt.Println("---- Advent of Code 2024 ----")
	fmt.Println("-----------------------------")
	fmt.Printf("----        Day %v        ----\n", day)
	fmt.Println("-----------------------------")
}

func ReadLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	dataStr := strings.TrimSpace(string(data))
	lines := strings.Split(dataStr, "\n")

	return lines, nil
}

func DeleteAtIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
