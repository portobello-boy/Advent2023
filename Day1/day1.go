package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

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

		for _, r := range str {
			if unicode.IsNumber(r) {
				calibrationValue[0], _ = strconv.Atoi(string(r))
				break
			}
		}

		for _, r := range reverse(str) {
			if unicode.IsNumber(r) {
				calibrationValue[1], _ = strconv.Atoi(string(r))
				break
			}
		}

		fmt.Println(calibrationValue)

		sum += 10*calibrationValue[0] + calibrationValue[1]
	}

	fmt.Println(sum)

}
