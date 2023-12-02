package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var configuration = map[string]int{"red": 12, "green": 13, "blue": 14}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open file!")
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ":")
		gameId, _ := strconv.Atoi(strings.Split(split[0], " ")[1])
		subsets := strings.Split(split[1], ";")
		// fmt.Println(subsets)

		maxVals := make(map[string]int)

		for _, hand := range subsets {
			diceByColor := strings.Split(hand, ",")

			for i := range diceByColor {
				diceByColor[i] = strings.TrimSpace(diceByColor[i])
			}

			for _, diceSet := range diceByColor {
				diceInfo := strings.Split(diceSet, " ")

				num, _ := strconv.Atoi(diceInfo[0])

				if _, ok := maxVals[diceInfo[1]]; !ok {
					maxVals[diceInfo[1]] = num
				}

				if maxVals[diceInfo[1]] < num {
					maxVals[diceInfo[1]] = num
				}
			}
		}

		valid := true
		for k, v := range maxVals {
			if v > configuration[k] {
				valid = false
			}
		}

		if valid {
			sum += gameId
		}
	}

	fmt.Println(sum)

}
