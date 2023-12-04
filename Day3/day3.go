package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

func LoadSchematic(filename string) (schematic []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("Couldn't open file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}

	return
}

func GetSurroundingSlice(schematic []string, rowIdx, colIdxStart, colIdxEnd int) (frame []string) {
	firstRowIdx := int(math.Max(float64(rowIdx)-1, 0))
	lastRowIdx := int(math.Min(float64(rowIdx)+1, float64(len(schematic))-1))

	firstColIdx := int(math.Max(float64(colIdxStart)-1, 0))
	lastColIdx := int(math.Min(float64(colIdxEnd)+1, float64(len(schematic[rowIdx]))))

	for i := firstRowIdx; i <= lastRowIdx; i++ {

		frame = append(frame, schematic[i][firstColIdx:lastColIdx])
	}

	return
}

func IsFrameValid(frame []string) bool {
	for _, fRow := range frame {
		for _, fChar := range fRow {
			if !unicode.IsNumber(fChar) && string(fChar) != "." {
				return true
			}
		}
	}
	return false
}

func main() {
	schematic := LoadSchematic("input.txt")

	sum := 0

	// iterate over rows
	for rowIndex, row := range schematic {
		// iterate over cols
		for colIndex := 0; colIndex < len(row); colIndex++ {
			char := rune(row[colIndex])
			// if current rune isn't a number, skip it
			if !unicode.IsNumber(char) {
				continue
			}

			startCol, endCol := colIndex, colIndex

			for unicode.IsNumber(char) {
				endCol++
				colIndex++
				if colIndex >= len(row) {
					break
				}
				char = rune(row[colIndex])
			}

			frame := GetSurroundingSlice(schematic, rowIndex, startCol, endCol)

			if IsFrameValid(frame) {
				val, _ := strconv.Atoi(schematic[rowIndex][startCol:endCol])

				sum += val
			}
		}
	}

	fmt.Println(sum)

}
