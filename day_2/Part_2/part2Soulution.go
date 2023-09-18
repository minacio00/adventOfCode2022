package part2

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
func returnLabels(op string) string {
	switch op {
	case "A":
		op = "Rock"
	case "B":
		op = "Papper"
	case "C":
		op = "Scissors"
	}
	return op
}

// returns the amount of points the player gets for losing the round
func loseRound(oppenent string) int {
	switch oppenent {
	case "Rock":
		return 3 // 3 because if i need to lose to rock that means I played scissors
	case "Papper":
		return 1
	case "Scissors":
		return 2
	default:
		log.Panicf("Invalid oppennt string: %s", oppenent)
		return -1
	}
}

func drawRound(oppenent string) int {
	switch oppenent {
	case "Rock":
		return 1 + 3
	case "Papper":
		return 2 + 3
	case "Scissors":
		return 3 + 3
	default:
		log.Panicf("Invalid oppennt string: %s", oppenent)
		return -1
	}
}
func winRound(oppenent string) int {
	switch oppenent {
	case "Rock":
		return 2 + 6
	case "Papper":
		return 3 + 6
	case "Scissors":
		return 1 + 6
	default:
		log.Panicf("Invalid oppennt string: %s", oppenent)
		return -1
	}
}

func SolvePart2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Panic(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		player := string(line[1])
		oppenent := returnLabels(string(line[0]))
		switch player {
		case "X":
			totalPoints += loseRound(oppenent)
		case "Y":
			totalPoints += drawRound(oppenent)
		case "Z":
			totalPoints += winRound(oppenent)
		}

	}
	println(totalPoints)
}
