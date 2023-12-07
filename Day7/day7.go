package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cards = []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

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

func SortHand(a, b string) bool {
	idxA := 0
	idxB := 0

	idx := 0
	for a[idx] == b[idx] {
		idx++
	}

	for i, c := range cards {
		if string(a[idx]) == c {
			idxA = i
			break
		}
	}

	for i, c := range cards {
		if string(b[idx]) == c {
			idxB = i
			break
		}
	}

	return idxA > idxB
}

func DetermineType(hand string) int {
	handMap := make(map[string]int)

	for _, c := range hand {
		card := string(c)
		if _, ok := handMap[card]; !ok {
			handMap[card] = 0
		}
		handMap[card]++
	}

	highestType := ""

	if _, ok := handMap["J"]; ok {
		highestOtherCount := 0
		highestOtherLetter := "J"
		for k, v := range handMap {
			if v > highestOtherCount && k != "J" {
				highestOtherCount = v
				highestOtherLetter = k
			}
		}

		highestType = highestOtherLetter
	}

	newHand := ""
	for _, c := range hand {
		if string(c) == "J" {
			newHand += highestType
		} else {
			newHand += string(c)
		}
	}

	for k := range handMap {
		delete(handMap, k)
	}

	for _, c := range newHand {
		card := string(c)
		if _, ok := handMap[card]; !ok {
			handMap[card] = 0
		}
		handMap[card]++
	}

	if len(handMap) == 1 {
		return 7 // all cards are same, 5 of a kind
	}

	if len(handMap) == 2 {
		for _, v := range handMap {
			if v == 4 {
				return 6 // 4 of a kind
			} else if v == 3 {
				return 5 // full house
			}
		}
	}

	if len(handMap) == 3 {
		countPairs := 0
		for _, v := range handMap {
			if v == 3 {
				return 4 // 3 of a kind
			} else if v == 2 {
				countPairs++
			}

			if countPairs == 2 {
				return 3 // 2 pair
			}
		}
	}

	if len(handMap) == 4 {
		return 2 // one pair
	}

	return 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open file!")
	}
	defer file.Close()

	// bids := make([]int, 0)
	// handInfoMap := make(map[string][]int)
	handBidMap := make(map[string]int)
	handTypeMap := make(map[int][]string)
	// idx := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := SplitAndTrim(scanner.Text(), " ")
		// hand := SplitAndTrim(split[0], "")
		hand := split[0]
		bid, _ := strconv.Atoi(split[1])

		handType := DetermineType(hand)

		handBidMap[hand] = bid

		if _, ok := handTypeMap[handType]; !ok {
			handTypeMap[handType] = make([]string, 0)
		}
		handTypeMap[handType] = append(handTypeMap[handType], hand)
	}

	// Sort hands within map
	for _, hands := range handTypeMap {
		sort.Slice(hands, func(i, j int) bool {
			return SortHand(hands[i], hands[j])
		})
	}

	score := 0
	rank := 1
	for i := 1; i <= 7; i++ {
		if hands, ok := handTypeMap[i]; ok {
			for _, h := range hands {
				score += rank * handBidMap[h]
				rank++
			}
		}
	}

	fmt.Println(score)

}
