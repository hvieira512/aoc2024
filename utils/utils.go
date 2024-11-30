package utils

import "fmt"

func RenderHeader(day int) {
	fmt.Println("---- Advent of Code 2024 ----")
	fmt.Println("-----------------------------")
	fmt.Printf("----------- Day %v -----------", day)
	fmt.Println("-----------------------------")
}
