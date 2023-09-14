package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()
	scanner := bufio.NewScanner(file)
	// var mostCalories [3]int
	var elfsTotal []int
	var currElfTotal int = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfsTotal = append(elfsTotal, currElfTotal)
			currElfTotal = 0
			continue
		}
		number, _ := strconv.Atoi(line)
		currElfTotal += number
		// for i := range mostCalories {
		// 	if currElfTotal > mostCalories[i] {
		// 		mostCalories[i] = currElfTotal
		// 		break
		// 	}

		// }
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfsTotal)))
	top3 := 0
	for i := 0; i < 3; i++ {
		top3 += elfsTotal[i]
	}
	println(top3)
}
