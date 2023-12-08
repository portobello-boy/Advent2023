package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type Node struct {
// 	Left  *Node
// 	Right *Node
// }

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

	for scanner.Scan() {
		mapping := SplitAndTrim(scanner.Text(), "=")
		// fmt.Println(mapping)

		key := mapping[0]
		valuesStr := strings.TrimSuffix(strings.TrimPrefix(mapping[1], "("), ")")
		values := SplitAndTrim(valuesStr, ",")

		// fmt.Println(key, values)
		directions[key] = values

	}

	// fmt.Println(directions)

	cur := "AAA"
	countSteps := 0
	routePatternIdx := 0
	routePatternLen := len(routePattern)

	for cur != "ZZZ" {
		switch string(routePattern[routePatternIdx]) {
		case "R":
			cur = directions[cur][1]
		case "L":
			cur = directions[cur][0]
		}
		countSteps++
		routePatternIdx += 1
		routePatternIdx %= routePatternLen
	}

	fmt.Println(countSteps)

}
