package main

import (
	part2 "day2/Part_2"
)

type Play int

const (
	X Play = iota
	Y
	Z
)

func (p *Play) String() string {
	return [...]string{"X", "Y", "Z"}[*p]
}

func main() {
	// part1.SolvePart1()
	part2.SolvePart2()
}
