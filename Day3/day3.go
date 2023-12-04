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

func GetSurroundingNumbers(schematic []string, rowIdx, colIdx int) (nums []int, numsCount int) {
	firstRowIdx := int(math.Max(float64(rowIdx)-1, 0))
	lastRowIdx := int(math.Min(float64(rowIdx)+1, float64(len(schematic))-1))

	firstColIdx := int(math.Max(float64(colIdx)-1, 0))
	lastColIdx := int(math.Min(float64(colIdx)+1, float64(len(schematic[rowIdx]))))

	// look at surrounding rows
	for i := firstRowIdx; i <= lastRowIdx; i++ {

		// look at surrounding cols
		for j := firstColIdx; j <= lastColIdx; j++ {
			char := rune(schematic[i][j])

			if !unicode.IsNumber(char) {
				continue
			}

			startCol, endCol := j, j

			// look backwards to find start of number
			for {
				if startCol <= 0 {
					break
				}
				char = rune(schematic[i][startCol-1])
				if !unicode.IsNumber(char) {
					break
				}
				startCol--
			}

			// look forwards to find start of number
			for {
				j++
				endCol = j
				if j >= len(schematic[i]) {
					break
				}
				char = rune(schematic[i][endCol])
				if !unicode.IsNumber(char) {
					break
				}
			}

			val, _ := strconv.Atoi(schematic[i][startCol:endCol])

			nums = append(nums, val)
			numsCount = len(nums)

		}
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

			// if the current rune isn't a *, skip it
			if string(char) != "*" {
				continue
			}

			nums, numsCount := GetSurroundingNumbers(schematic, rowIndex, colIndex)

			if numsCount != 2 {
				continue
			}

			sum += nums[0] * nums[1]
		}
	}

	fmt.Println(sum)

}
