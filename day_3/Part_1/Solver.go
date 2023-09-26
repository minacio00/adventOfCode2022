package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getPriority(itemType rune) int {
	if itemType >= 'a' && itemType <= 'z' {
		return int(itemType - 'a' + 1)
	} else if itemType >= 'A' && itemType <= 'Z' {
		return int(itemType - 'A' + 27)
	} else {
		return 0
	}
}
func getRepeatedItems(s string) int {
	length := len(s)
	half := length / 2
	for i, v := range s[:half] {
		for j, x := range s[half:] {
			if v == x {
				fmt.Printf("i: %d, j: %d\n", i, j)
				sum := getPriority(x) + getPriority(v)
				return sum
			}
		}
	}
	return -1
}

func SolvePart1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Panic(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var total int // sum of the priority number of the items that appear on both compartments
	for scanner.Scan() {
		line := scanner.Text()
		total += getRepeatedItems(line)
	}
	println(total / 2) // gambiarra do krl o problema é que os repetidos estão sendo somados duas vezes
}
