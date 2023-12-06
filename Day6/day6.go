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

func f(x, t int) int {
	return (t - x) * x
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	timeStr := scanner.Text()
	scanner.Scan()
	distStr := scanner.Text()

	timeSplit := TrimSpaces(strings.Split(timeStr, ":"))
	distSplit := TrimSpaces(strings.Split(distStr, ":"))

	time, _ := strconv.Atoi(strings.Join(TrimSpaces(strings.Split(timeSplit[1], " ")), ""))
	dist, _ := strconv.Atoi(strings.Join(TrimSpaces(strings.Split(distSplit[1], " ")), ""))

	fmt.Println(time, dist)

	// times := StringSliceToIntSlice(TrimSpaces(strings.Split(timeSplit[1], " ")))
	// dists := StringSliceToIntSlice(TrimSpaces(strings.Split(distSplit[1], " ")))

	// fmt.Println(times, dists)

	// prod := 1
	// for i, t := range times {

	count := 0

	for x := 0; x <= time; x++ {
		if f(x, time) > dist {
			count++
		}
	}

	// 	prod *= count
	// }
	fmt.Println(count)

}
