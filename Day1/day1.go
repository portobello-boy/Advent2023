package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numericStrings = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var reversedNumericStrings = [9]string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open file!")
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		calibrationValue := [2]int{}

		// First find any numeric character
		// Second find if any numeric strings exist before that character
		minNumDigitIndex := 0

		for _, r := range str {
			if unicode.IsNumber(r) {
				calibrationValue[0], _ = strconv.Atoi(string(r))
				break
			}
			minNumDigitIndex++
		}

		for numValue, num := range numericStrings {
			idx := strings.Index(str, num)
			if idx != -1 && idx < minNumDigitIndex {
				calibrationValue[0] = numValue + 1
				minNumDigitIndex = idx
			}
		}

		// Reset for reversed string
		minNumDigitIndex = 0
		reversedStr := reverse(str)

		for _, r := range reversedStr {
			if unicode.IsNumber(r) {
				calibrationValue[1], _ = strconv.Atoi(string(r))
				break
			}
			minNumDigitIndex++
		}

		for numValue, num := range reversedNumericStrings {
			idx := strings.Index(reversedStr, num)
			if idx != -1 && idx < minNumDigitIndex {
				calibrationValue[1] = numValue + 1
				minNumDigitIndex = idx
			}
		}

		// fmt.Println(calibrationValue)

		sum += 10*calibrationValue[0] + calibrationValue[1]
	}

	fmt.Println(sum)

}
