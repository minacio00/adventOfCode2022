package part1

import (
	"bufio"
	"log"
	"os"
	"strings"
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

func returnPlayerScore(played string) int {
	switch played {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return -1
	}
}
func returnLabels(op string, pl string) (string, string) {
	switch pl {
	case "X":
		pl = "Rock"
	case "Y":
		pl = "Papper"
	case "Z":
		pl = "Scissors"
	default:

	}
	switch op {
	case "A":
		op = "Rock"
	case "B":
		op = "Papper"
	case "C":
		op = "Scissors"
	}
	return op, pl
}
func SolvePart1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Panic(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()
		println(totalPoints)
		line = strings.ReplaceAll(line, " ", "")
		opponent, player := returnLabels(string(line[0]), string(line[1]))
		if opponent == player {
			totalPoints += 3 + returnPlayerScore(string(line[1]))
		} else {
			r := returnPlayerScore(string(line[1]))
			switch char := string(line[0]); char {
			case "A":
				if r == 2 {
					// if 2 equals papper
					totalPoints += r + 6
				} else if r == 3 {
					// if 3 scissor
					totalPoints += r
				}
			case "B":
				if r == 1 {
					// if 1 equal rock
					totalPoints += r
				} else if r == 3 {
					totalPoints += r + 6
				}
			case "C":
				if r == 1 {
					totalPoints += r + 6
				} else if r == 2 {
					totalPoints += r
				}
			}

		}

	}

	println(totalPoints)
	var t Play = 1
	println(t.String())
}
