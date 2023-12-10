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
	"|": {Position{0, -1}, Position{0, 1}},
	"-": {Position{-1, 0}, Position{1, 0}},
	"L": {Position{0, -1}, Position{1, 0}},
	"J": {Position{0, -1}, Position{-1, 0}},
	"7": {Position{0, 1}, Position{-1, 0}},
	"F": {Position{0, 1}, Position{1, 0}},
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

func DetermineStartPipe(grid []string, gridLen, gridWidth int, start Position) (str string) {
	// Since we don't know the pipe under S, check all neighbors and find which are valid
	if string(grid[start.Y][start.X]) != "S" {
		return
	}

	surroundingPipes := GetSurroundingPipes(grid, gridLen, gridWidth, start)

	for k, v := range directionMap {
		candidateOne := Position{start.X + v[0].X, start.Y + v[0].Y}
		candidateTwo := Position{start.X + v[1].X, start.Y + v[1].Y}

		if (candidateOne == surroundingPipes[0] && candidateTwo == surroundingPipes[1]) || (candidateOne == surroundingPipes[1] && candidateTwo == surroundingPipes[0]) {
			return k
		}
	}
	return
}

func IsInLoop(char string, idx int, row string) bool {
	crossings := 0
	levelWithPipe := 0
	for _, c := range row[idx:] {
		switch string(c) {
		case "|":
			crossings++
			levelWithPipe = 0
		case "F":
			levelWithPipe = 1
		case "L":
			levelWithPipe = -1
		case "J":
			if levelWithPipe == 1 {
				crossings++
			}
			levelWithPipe = 0
		case "7":
			if levelWithPipe == -1 {
				crossings++

			}
			levelWithPipe = 0
		case ".":
			levelWithPipe = 0
		}
	}
	return crossings%2 == 1
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

	gridWithMainLoop := make([]string, gridLen)
	rowTemplate := ""
	for i := 0; i < gridWidth; i++ {
		rowTemplate += "."
	}
	for i := 0; i < gridLen; i++ {
		gridWithMainLoop[i] = rowTemplate
	}

	gridWithMainLoop[startCoords.Y] = gridWithMainLoop[startCoords.Y][0:startCoords.X] + DetermineStartPipe(grid, gridLen, gridWidth, startCoords) + gridWithMainLoop[startCoords.Y][startCoords.X+1:]

	// From starting position, recursively follow all pipes until all pipes are exhausted, track position with largest distance

	queue := make([]QueueObject, 1)
	queue[0] = QueueObject{startCoords, 0}

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

			if string(grid[n.Y][n.X]) != "S" {
				gridWithMainLoop[n.Y] = gridWithMainLoop[n.Y][0:n.X] + string(grid[n.Y][n.X]) + gridWithMainLoop[n.Y][n.X+1:]
			}

			if explored, ok := exploredMap[n]; !ok || !explored {
				queue = append(queue, QueueObject{n, cur.DistanceFromStart + 1})
			}
		}
	}

	for _, r := range gridWithMainLoop {
		fmt.Println(r)

	}

	fmt.Println(maxDistance)

	// Fidn grid with main loop first, then find which points are contained
	count := 0

	for rowIdx, row := range gridWithMainLoop {
		for i, c := range row {
			if string(c) == "." && IsInLoop(string(c), i, row) {
				fmt.Printf("character %s at row %d col %d is INSIDE of loop\n", string(c), rowIdx, i)
				count++
			}
		}
	}

	fmt.Println(count)

}
