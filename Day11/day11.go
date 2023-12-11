package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

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

	rowExpansionMap := make(map[int]bool)
	colExpansionMap := make(map[int]bool)

	// Expand the galaxy rows
	for i := 0; i < gridLen; i++ {
		r := grid[i]
		if r == rowTemplate {
			rowExpansionMap[i] = true
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

		colExpansionMap[i] = true
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
	scale := 1000000

	for i, gp := range galaxyPositions {
		for _, gp2 := range galaxyPositions[i+1:] {
			dist := 0

			for r := math.Min(float64(gp.X), float64(gp2.X)); r < math.Max(float64(gp.X), float64(gp2.X)); r++ {
				if _, ok := colExpansionMap[int(r)]; ok {
					dist += scale
				} else {
					dist += 1
				}
			}

			for c := math.Min(float64(gp.Y), float64(gp2.Y)); c < math.Max(float64(gp.Y), float64(gp2.Y)); c++ {
				if _, ok := rowExpansionMap[int(c)]; ok {
					dist += scale
				} else {
					dist += 1
				}
			}
			sum += dist
		}
	}

	fmt.Println(sum)

}
