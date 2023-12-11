package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

type QueueObject struct {
	Pos               Position
	DistanceFromStart int
}

// A structure to hold the necessary data for a goroutine worker
type Job struct {
	Id  int
	Gal Position
}

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

	grid := make([]string, 0)
	gridWidth := -1
	gridLen := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		grid = append(grid, row)

		gridLen++

		if gridWidth == -1 {
			gridWidth = len(row)
		}
	}

	rowTemplate := ""
	for i := 0; i < gridWidth; i++ {
		rowTemplate += "."
	}

	// Expand the galaxy rows
	for i := 0; i < gridLen; i++ {
		r := grid[i]
		if r == rowTemplate {
			grid = append(append(grid[:i], rowTemplate), grid[i:]...)
			i++
		}
	}

	gridLen = len(grid)

	// Expand the galaxy columns
	for i := 0; i < gridWidth; i++ {
		valid := true
		for j := 0; j < gridLen; j++ {
			if string(grid[j][i]) != "." {
				valid = false
			}
		}

		if !valid {
			continue
		}

		for j := 0; j < gridLen; j++ {
			grid[j] = grid[j][:i] + "." + grid[j][i:]
		}

		gridWidth++
		i++
	}

	galaxyPositions := make([]Position, 0)

	for i, r := range grid {
		for j, c := range r {
			if string(c) == "#" {
				galaxyPositions = append(galaxyPositions, Position{j, i})
			}
		}
	}

	fmt.Println(galaxyPositions)

	for _, r := range grid {
		fmt.Println(r)

	}

	sum := 0

	for i, gp := range galaxyPositions {
		for _, gp2 := range galaxyPositions[i+1:] {
			sum += int(math.Abs(float64(gp.X)-float64(gp2.X)) + math.Abs(float64(gp.Y)-float64(gp2.Y)))
		}
	}

	fmt.Println(sum)

}
