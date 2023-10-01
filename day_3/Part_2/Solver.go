package part2

import (
	"bufio"
	"log"
	"os"
)

func findBadges(group []string) string {

	charCount := make(map[rune]int)
	for _, v := range group[0] {
		charCount[v]++
	}
	for _, line := range group[1:] {
		lineCharCount := make(map[rune]int)
		for _, v := range line {
			lineCharCount[v]++
		}
		for char := range charCount {
			if lineCharCount[char] == 0 {
				delete(charCount, char)
			}
		}
	}

	for char := range charCount {
		if len(charCount) > 1 {
			log.Panic("more than one badge founded in group:", group)
		}
		return string(char)
	}

	return "a"
}

func getPriority(itemType rune) int {
	if itemType >= 'a' && itemType <= 'z' {
		return int(itemType - 'a' + 1)
	} else if itemType >= 'A' && itemType <= 'Z' {
		return int(itemType - 'A' + 27)
	} else {
		return 0
	}
}

func SolvePart2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Panic(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// badges := make(map[string]int)
	var counter = 0 // used for separating the groups
	var group []string
	badgesArr := []int{}
	// the code is missing on Scan
	for {
		if counter <= 2 {
			flag := scanner.Scan()
			if !flag {
				break
			}
			group = append(group, scanner.Text())
			counter++
		} else {
			badge := findBadges(group)
			// badges[badge] = getPriority(rune(badge[0]))
			badgesArr = append(badgesArr, getPriority(rune(badge[0])))
			counter = 0
			group = []string{}
			continue
		}
	}
	sum := 0
	for i, v := range badgesArr {
		println(i)
		sum += v
	}
	println(sum)

}
