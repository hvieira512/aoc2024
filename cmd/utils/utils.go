package utils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func RenderDayHeader(day int) {
	if day < 10 {
		fmt.Println("-----------------------------")
		fmt.Println("---- Advent of Code 2024 ----")
		fmt.Println("-----------------------------")
		fmt.Printf("----        Day %v        ----\n", day)
		fmt.Println("-----------------------------")
	} else {
		fmt.Println("-----------------------------")
		fmt.Println("---- Advent of Code 2024 ----")
		fmt.Println("-----------------------------")
		fmt.Printf("---        Day %v        ----\n", day)
		fmt.Println("-----------------------------")
	}
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

func ReadGridRune(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	grid := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func DeleteAtIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func SliceExists(target []int, list [][]int) bool {
	for _, slice := range list {
		if reflect.DeepEqual(slice, target) {
			return true
		}
	}
	return false
}

func ReadGridInt(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	grid := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		row := []int{}
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func InBoundsGrid(x, y, n int) bool {
	return x >= 0 && y >= 0 && x < n && y < n
}
