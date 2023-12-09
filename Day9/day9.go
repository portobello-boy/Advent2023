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

func SliceAllZeroes(intSlice []int) bool {
	for _, i := range intSlice {
		if i != 0 {
			return false
		}
	}

	return true
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
		differenceHistory := make([][]int, 1)
		ints := StringSliceToIntSlice(strings.Split(scanner.Text(), " "))
		intsLen := len(ints)
		// fmt.Println(ints)

		differenceHistory[0] = ints

		curSlice := differenceHistory[0]
		for !SliceAllZeroes(curSlice) {
			curSliceLen := len(curSlice)
			differences := make([]int, curSliceLen-1)

			for i := 0; i < curSliceLen-1; i++ {
				differences[i] = curSlice[i+1] - curSlice[i]
			}

			differenceHistory = append(differenceHistory, differences)
			curSlice = differences
		}

		fmt.Println(differenceHistory)

		nextVal := 0
		for i, differences := range differenceHistory {
			nextVal += differences[intsLen-i-1]
		}

		fmt.Println(nextVal)
		sum += nextVal

	}
	fmt.Println(sum)

}
