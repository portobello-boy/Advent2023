package main

import (
	"bufio"
	"fmt"
	"os"
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

func SplitAndTrim(str, delim string) []string {
	return TrimSpaces(strings.Split(str, delim))
}

func AtEnd(positions []string) bool {
	for _, pos := range positions {
		if string(pos[2]) != "Z" {
			return false
		}
	}
	return true
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open file!")
	}
	defer file.Close()

	directions := make(map[string][]string)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	routePattern := scanner.Text()
	scanner.Scan()

	startingPatterns := make([]string, 0)

	for scanner.Scan() {
		mapping := SplitAndTrim(scanner.Text(), "=")

		key := mapping[0]
		valuesStr := strings.TrimSuffix(strings.TrimPrefix(mapping[1], "("), ")")
		values := SplitAndTrim(valuesStr, ",")

		if string(key[2]) == "A" {
			startingPatterns = append(startingPatterns, key)
		}

		directions[key] = values

	}

	routePatternLen := len(routePattern)
	startingPatternStepsCounts := make([]int, 0)

	for _, pos := range startingPatterns {
		countSteps := 0
		routePatternIdx := 0
		for !AtEnd([]string{pos}) {
			dir := -1
			switch string(routePattern[routePatternIdx]) {
			case "R":
				dir = 1
			case "L":
				dir = 0
			}

			pos = directions[pos][dir]

			countSteps++
			routePatternIdx += 1
			routePatternIdx %= routePatternLen
		}
		startingPatternStepsCounts = append(startingPatternStepsCounts, countSteps)
	}

	fmt.Println(LCM(startingPatternStepsCounts[0], startingPatternStepsCounts[1], startingPatternStepsCounts...))
}
