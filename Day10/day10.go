package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	X int
	Y int
}

type QueueObject struct {
	Pos               Position
	DistanceFromStart int
}

var directionMap = map[string][]Position{
	"|": []Position{Position{0, -1}, Position{0, 1}},
	"-": []Position{Position{-1, 0}, Position{1, 0}},
	"L": []Position{Position{0, -1}, Position{1, 0}},
	"J": []Position{Position{0, -1}, Position{-1, 0}},
	"7": []Position{Position{0, 1}, Position{-1, 0}},
	"F": []Position{Position{0, 1}, Position{1, 0}},
}

func StringSliceToIntSlice(stringSlice []string) (intSlice []int) {
	for _, str := range stringSlice {
		val, _ := strconv.Atoi(str)
		intSlice = append(intSlice, val)
	}
	return
}

func IsInGrid(grid []string, gridLen, gridWidth int, pos Position) bool {
	return pos.X >= 0 && pos.X < gridWidth && pos.Y >= 0 && pos.Y < gridLen
}

func GetSurroundingPipes(grid []string, gridLen, gridWidth int, cur Position) (neighbors []Position) {
	neighboringPositions := []Position{
		Position{-1, 0},
		Position{1, 0},
		Position{0, -1},
		Position{0, 1},
	}

	// Since we don't know the pipe under S, check all neighbors and find which are valid
	if string(grid[cur.Y][cur.X]) == "S" {
		for _, pos := range neighboringPositions {
			candidateNeighbor := Position{cur.X + pos.X, cur.Y + pos.Y}
			if IsInGrid(grid, gridLen, gridWidth, candidateNeighbor) {
				for _, v := range directionMap[string(grid[candidateNeighbor.Y][candidateNeighbor.X])] {
					tmp := Position{candidateNeighbor.X + v.X, candidateNeighbor.Y + v.Y}
					if string(grid[tmp.Y][tmp.X]) == "S" {
						neighbors = append(neighbors, candidateNeighbor)
					}
				}
			}
		}

		return
	}

	for _, pos := range directionMap[string(grid[cur.Y][cur.X])] {
		candidateNeighbor := Position{cur.X + pos.X, cur.Y + pos.Y}
		if IsInGrid(grid, gridLen, gridWidth, candidateNeighbor) {
			if string(grid[candidateNeighbor.Y][candidateNeighbor.X]) != "." {
				neighbors = append(neighbors, candidateNeighbor)
			}
		}
	}

	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open file!")
	}
	defer file.Close()

	grid := make([]string, 0)

	startCoords := Position{-1, -1}
	gridLen := 0
	gridWidth := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		grid = append(grid, row)

		for i, c := range row {
			if string(c) == "S" {
				startCoords.X = i
				startCoords.Y = gridLen
			}
		}

		if gridWidth == -1 {
			gridWidth = len(row)
		}

		gridLen++
	}

	distanceMap := make(map[Position]int)
	exploredMap := make(map[Position]bool)
	// From starting position, recursively follow all pipes until all pipes are exhausted, track position with largest distance

	queue := make([]QueueObject, 1)
	queue[0] = QueueObject{startCoords, 0}

	// fmt.Println(grid, startCoords, distanceMap, exploredMap, queue)
	// fmt.Println(startCoords)

	// fmt.Println(GetSurroundingPipes(grid, gridLen, gridWidth, startCoords))

	maxDistance := 0

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		exploredMap[cur.Pos] = true
		distanceMap[cur.Pos] = cur.DistanceFromStart

		// enqueue surrounding pipes which haven't been explored
		neighbors := GetSurroundingPipes(grid, gridLen, gridWidth, cur.Pos)

		if cur.DistanceFromStart > maxDistance {
			maxDistance = cur.DistanceFromStart
		}

		for _, n := range neighbors {
			if explored, ok := exploredMap[n]; !ok || !explored {
				queue = append(queue, QueueObject{n, cur.DistanceFromStart + 1})
			}
		}
	}

	fmt.Println(maxDistance)

}
