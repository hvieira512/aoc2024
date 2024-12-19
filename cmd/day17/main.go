package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2024/cmd/utils"
)

const (
	Filename = "input"
)

func partTwo(registers map[rune]int, program []int) int {
	if len(program) < 2 || program[len(program)-2] != 3 || program[len(program)-1] != 0 {
		panic("program does not end with JNZ 0")
	}

	var find func(target []int, ans int) int
	find = func(target []int, ans int) int {
		if len(target) == 0 {
			return ans
		}
		for t := 0; t < 8; t++ {
			registers['A'] = (ans << 3) | t
			registers['B'], registers['C'] = 0, 0
			var output *int
			adv3 := false

			combo := func(operand int) int {
				switch {
				case operand >= 0 && operand <= 3:
					return operand
				case operand == 4:
					return registers['A']
				case operand == 5:
					return registers['B']
				case operand == 6:
					return registers['C']
				default:
					panic(fmt.Sprintf("unrecognized combo operand %d", operand))
				}
			}

			for ip := 0; ip < len(program)-2; ip += 2 {
				ins := program[ip]
				operand := program[ip+1]

				switch ins {
				case 0:
					if adv3 {
						panic("program has multiple ADVs")
					}
					if operand != 3 {
						panic("program has ADV with operand other than 3")
					}
					adv3 = true
				case 1:
					registers['B'] = registers['B'] ^ operand
				case 2:
					registers['B'] = combo(operand) % 8
				case 3:
					panic("program has JNZ inside expected loop body")
				case 4:
					registers['B'] = registers['B'] ^ registers['C']
				case 5:
					if output != nil {
						panic("program has multiple OUT")
					}
					out := combo(operand) % 8
					output = &out
				case 6:
					registers['B'] = registers['A'] >> combo(operand)
				case 7:
					registers['C'] = registers['A'] >> combo(operand)
				default:
					panic(fmt.Sprintf("unrecognized instruction %d", ins))
				}

				if output != nil && *output == target[len(target)-1] {
					sub := find(target[:len(target)-1], registers['A'])
					if sub != -1 {
						return sub
					}
				}
			}
		}
		return -1 // If no solution is found
	}

	return find(program, 0)
}

func partOne(registers map[rune]int, program []int) string {
	ip := 0
	output := []int{}

	combo := func(operand int) int {
		if operand >= 0 && operand <= 3 {
			return operand
		} else if operand == 4 {
			return registers['A']
		} else if operand == 5 {
			return registers['B']
		} else if operand == 6 {
			return registers['C']
		} else {
			panic(fmt.Sprintf("unrecognized combo operand %d", operand))
		}
	}

	for ip < len(program) {
		ins := program[ip]
		operand := program[ip+1]
		ip += 2

		if ins == 0 {
			registers['A'] = registers['A'] >> combo(operand)
		} else if ins == 1 {
			registers['B'] = registers['B'] ^ operand
		} else if ins == 2 {
			registers['B'] = combo(operand) % 8
		} else if ins == 3 {
			if registers['A'] != 0 {
				ip = operand
				continue
			}
		} else if ins == 4 {
			registers['B'] = registers['B'] ^ registers['C']
		} else if ins == 5 {
			output = append(output, combo(operand)%8)
		} else if ins == 6 {
			registers['B'] = registers['A'] >> combo(operand)
		} else if ins == 7 {
			registers['C'] = registers['A'] >> combo(operand)
		} else {
			panic(fmt.Sprintf("unrecognized instruction %d", ins))
		}
	}

	var out string
	for i, val := range output {
		aux := ""
		aux += fmt.Sprintf("%d", val)
		if i != len(output)-1 {
			aux += ","
		}
		out += aux
	}
	return out
}

func main() {
	utils.RenderDayHeader(17)
	lines, _ := utils.Strings("cmd/day17/" + Filename + ".txt")

	registers, program := parseInput(lines)

	fmt.Printf("Part 1: %v\n", partOne(registers, program))
	fmt.Printf("Part 2: %v\n", partTwo(registers, program))
}

func parseInput(lines []string) (map[rune]int, []int) {
	registers := map[rune]int{
		'A': 0,
		'B': 0,
		'C': 0,
	}
	registers['A'], _ = strconv.Atoi(strings.Split(lines[0], ": ")[1])

	program := []int{}
	for _, aux := range strings.Split(strings.Split(lines[4], ": ")[1], ",") {
		num, _ := strconv.Atoi(aux)
		program = append(program, num)
	}

	return registers, program
}
