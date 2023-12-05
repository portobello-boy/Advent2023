package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func StringSliceToIntSlice(stringSlice []string) (intSlice []int) {
	for _, str := range stringSlice {
		val, _ := strconv.Atoi(str)
		intSlice = append(intSlice, val)
	}
	return
}

func TrimSpaces(strSlice []string) (trimmed []string) {
	for _, str := range strSlice {
		if str == "" {
			continue
		}
		// fmt.Printf("OG: '%v', TRIMMED: '%v'\n", str, strings.TrimSpace(str))
		trimmed = append(trimmed, strings.TrimSpace(str))
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open file!")
	}
	defer file.Close()

	scratchcards := make(map[int]int)
	current := 0
	scratchcards[current] = 0 // Initialize first scratchcard

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scratchcards[current]++ // Got new scratchcard
		split := strings.Split(scanner.Text(), ":")

		// split goddamn everything, trim everything
		split = TrimSpaces(split)
		numberSets := TrimSpaces(strings.Split(split[1], "|"))
		winningSplit := TrimSpaces(strings.Split(numberSets[0], " "))
		obtainedSplit := TrimSpaces(strings.Split(numberSets[1], " "))

		winningSet := StringSliceToIntSlice(TrimSpaces(winningSplit))
		winningMap := make(map[int]bool)

		// generate map of winning numbers
		for _, num := range winningSet {
			winningMap[num] = true
		}

		obtainedSet := StringSliceToIntSlice(TrimSpaces(obtainedSplit))
		sort.Slice(obtainedSet, func(a, b int) bool {
			return obtainedSet[a] < obtainedSet[b]
		})

		// get count of winning numbers in obtainedSet
		count := 0
		for _, num := range obtainedSet {
			if _, ok := winningMap[num]; ok {
				count++
			}
		}

		// iterate over count of winning numbers
		for i := 1; i <= count; i++ {
			if _, ok := scratchcards[current+i]; !ok {
				scratchcards[current+i] = 0
			}
			// increment the count of future scratchcards with the current scratchcard's count
			scratchcards[current+i] += scratchcards[current]
		}

		current++
	}

	sum := 0
	for _, count := range scratchcards {
		sum += count
	}

	fmt.Println(sum)

}
