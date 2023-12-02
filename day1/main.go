package main

import (
	"fmt"
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

	regexMatchPattern := `[0-9]`
	regexMatch := regexp.MustCompile(regexMatchPattern)

	for _, line := range inputLineArray {
		numbers := regexMatch.FindAllString(line, -1)
		value, _ := strconv.Atoi(fmt.Sprintf("%s%s", numbers[0], numbers[len(numbers)-1]))
		sum += value
	}

	fmt.Printf("Part 1 - Total Sum: %v\n", sum)
}

func part2() {
	inputLineArray := read.ReadStrArrayByLine("./exinput.txt")

	sum := 0

	//leaving for reference, doesn't work because golang regex does not support overlapping matches like 'eightwo'
	//digitMatchPattern := `([1-9]|one|two|three|four|five|six|seven|eight|nine)`

	digitMatchPattern := `[1-9]`
	digitMatch := regexp.MustCompile(digitMatchPattern)
	oneMP := `one`
	oneM := regexp.MustCompile(oneMP)
	twoMP := `two`
	twoM := regexp.MustCompile(twoMP)
	threeMP := `three`
	threeM := regexp.MustCompile(threeMP)
	fourMP := `four`
	fourM := regexp.MustCompile(fourMP)
	fiveMP := `five`
	fiveM := regexp.MustCompile(fiveMP)
	sixMP := `six`
	sixM := regexp.MustCompile(sixMP)
	sevenMP := `seven`
	sevenM := regexp.MustCompile(sevenMP)
	eightMP := `eight`
	eightM := regexp.MustCompile(eightMP)
	nineMP := `nine`
	nineM := regexp.MustCompile(nineMP)

	for _, line := range inputLineArray {
		numbers := digitMatch.FindAllStringIndex(line, -1)
		ones := oneM.FindAllStringIndex(line, -1)
		twos := twoM.FindAllStringIndex(line, -1)
		threes := threeM.FindAllStringIndex(line, -1)
		fours := fourM.FindAllStringIndex(line, -1)
		fives := fiveM.FindAllStringIndex(line, -1)
		sixs := sixM.FindAllStringIndex(line, -1)
		sevens := sevenM.FindAllStringIndex(line, -1)
		eights := eightM.FindAllStringIndex(line, -1)
		nines := nineM.FindAllStringIndex(line, -1)

		first := getNumberString(numbers[0])
		last := getNumberString(numbers[len(numbers)-1])
		fmt.Printf("First: %s Last: %s\n", first, last)
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
