package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/SotoDucani/AoC2022/internal/read"
)

func getNumberString(number string) string {
	var first string
	switch number {
	case "one":
		first = "1"
	case "two":
		first = "2"
	case "three":
		first = "3"
	case "four":
		first = "4"
	case "five":
		first = "5"
	case "six":
		first = "6"
	case "seven":
		first = "7"
	case "eight":
		first = "8"
	case "nine":
		first = "9"
	default:
		first = number
	}
	return first
}

func part1() {
	inputLineArray := read.ReadStrArrayByLine("./input.txt")

	sum := 0

	regexMatch := regexp.MustCompile(`\d`)

	for _, line := range inputLineArray {
		numbers := regexMatch.FindAllString(line, -1)
		value, _ := strconv.Atoi(fmt.Sprintf("%s%s", numbers[0], numbers[len(numbers)-1]))
		sum += value
	}

	fmt.Printf("Part 1 - Total Sum: %v\n", sum)
}

func part2() {
	inputLineArray := read.ReadStrArrayByLine("./input.txt")
	sum := 0
	//leaving for reference, doesn't work because golang regex does not support overlapping matches like 'eightwo'
	//digitMatch := regexp.MustCompile(`([1-9]|one|two|three|four|five|six|seven|eight|nine)`)
	digitMatch := regexp.MustCompile(`\d`)
	oneM := regexp.MustCompile(`one`)
	twoM := regexp.MustCompile(`two`)
	threeM := regexp.MustCompile(`three`)
	fourM := regexp.MustCompile(`four`)
	fiveM := regexp.MustCompile(`five`)
	sixM := regexp.MustCompile(`six`)
	sevenM := regexp.MustCompile(`seven`)
	eightM := regexp.MustCompile(`eight`)
	nineM := regexp.MustCompile(`nine`)

	for _, line := range inputLineArray {
		limitMatches := map[string][]int{
			"first": {math.MaxInt, math.MaxInt},
			"last":  {0, 0},
		}
		numbers := digitMatch.FindAllStringIndex(line, -1)
		ones := oneM.FindAllStringIndex(line, -1)
		twos := twoM.FindAllStringIndex(line, -1)
		threes := threeM.FindAllStringIndex(line, -1)
		fours := fourM.FindAllStringIndex(line, -1)
		fives := fiveM.FindAllStringIndex(line, -1)
		sixes := sixM.FindAllStringIndex(line, -1)
		sevens := sevenM.FindAllStringIndex(line, -1)
		eights := eightM.FindAllStringIndex(line, -1)
		nines := nineM.FindAllStringIndex(line, -1)
		matchesLists := [][][]int{
			numbers, ones, twos, threes, fours, fives, sixes, sevens, eights, nines,
		}
		for _, matches := range matchesLists {
			//fmt.Printf("%v\n", matches)
			for _, indexPair := range matches {
				//Find First by checking the starting index of the match
				// If less, it's earlier in the string
				if indexPair[0] < limitMatches["first"][0] {
					// if less, replace the limit matches with this match
					limitMatches["first"][0] = indexPair[0]
					limitMatches["first"][1] = indexPair[1]
				}
				//Find Last by checking the starting index of the match
				// If greater, it's later in the string
				if indexPair[0] > limitMatches["last"][0] {
					limitMatches["last"][0] = indexPair[0]
					limitMatches["last"][1] = indexPair[1]
				}
			}
		}
		firstMatch := line[limitMatches["first"][0]:limitMatches["first"][1]]
		lastMatch := line[limitMatches["last"][0]:limitMatches["last"][1]]
		//fmt.Printf("Matched: %s, %s\n", firstMatch, lastMatch)

		first := getNumberString(firstMatch)
		last := getNumberString(lastMatch)
		if last == "" {
			last = first
		}
		fmt.Printf("%s - %s%s\n", line, first, last)
		value, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		sum += value
	}

	fmt.Printf("Part 2 - Total Sum: %v\n", sum)
}

func main() {
	p1b := time.Now()
	part1()
	mid := time.Now()
	part2()
	p2a := time.Now()
	part1Time := mid.Sub(p1b)
	part2Time := p2a.Sub(mid)
	fmt.Printf("Part 1 Time: %dμs\nPart 2 Time: %dμs\n", part1Time.Microseconds(), part2Time.Microseconds())
}
