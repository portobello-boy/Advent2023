package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open file!")
	}
	defer file.Close()

	// maps := make(map[string]map[int]int)
	maps := make(map[string][]func(int) int)
	mapOrder := make([]string, 0)

	scanner := bufio.NewScanner(file)
	mapRegex, _ := regexp.Compile(`.* map:`)

	scanner.Scan()
	seedStr := strings.Split(scanner.Text(), ":")
	seeds := strings.TrimSpace(seedStr[1])
	seedSlice := StringSliceToIntSlice(strings.Split(seeds, " "))

	fmt.Println(seedSlice)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if matched := mapRegex.MatchString(line); matched {
			mapInfoStr := strings.Split(line, " ")
			mapName := mapInfoStr[0]
			mapOrder = append(mapOrder, mapName)
			maps[mapName] = make([]func(int) int, 0)

			scanner.Scan()
			for mapDetails := scanner.Text(); mapDetails != ""; mapDetails = scanner.Text() {
				scanner.Scan()

				mapDetailsSlice := StringSliceToIntSlice(strings.Split(mapDetails, " "))
				// for i := 0; i < mapDetailsSlice[2]; i++ {
				// 	maps[mapName][mapDetailsSlice[1]+i] = mapDetailsSlice[0] + i
				// }
				maps[mapName] = append(maps[mapName], func(a int) int {
					destStart := mapDetailsSlice[0]
					sourceStart := mapDetailsSlice[1]
					rangeLength := mapDetailsSlice[2]

					if !(a-sourceStart > 0 && a-(sourceStart+rangeLength) < 0) {
						return a
					}
					return destStart + (a - sourceStart)
				})
				// fmt.Println(mapDetailsSlice)

			}
		}
	}
	// fmt.Printf("%v\n", maps)

	minValue := -1
	for _, seed := range seedSlice {
		finalValue := seed
		// for _, seedMapping := range maps {
		for _, seedMapping := range mapOrder {
			for _, function := range maps[seedMapping] {
				if tmpVal := function(finalValue); tmpVal != finalValue {
					finalValue = tmpVal
					break
				}
			}
			// if mapping, ok := maps[seedMapping][finalValue]; ok {
			// 	// fmt.Printf("%s: %d -> %d\n", seedMapping, finalValue, mapping)
			// 	finalValue = mapping
			// }
		}
		fmt.Println(seed, finalValue)

		if minValue == -1 {
			minValue = finalValue
		}
		minValue = int(math.Min(float64(minValue), float64(finalValue)))
	}

	fmt.Println("MIN VALUE", minValue)

}
