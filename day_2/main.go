package main

import (
	"bufio"
	"os"
	"strings"
)

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
func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()
	scanner := bufio.NewScanner(file)
	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		if string(line[0]) == string(line[1]) {
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
}
