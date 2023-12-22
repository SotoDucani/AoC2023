package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/SotoDucani/AoC2022/internal/read"
	"github.com/fatih/color"
)

type foundNumber struct {
	number       int
	indexes      []int
	rowFound     int
	xLeft        int
	xRight       int
	yLow         int
	yHigh        int
	isPartNumber bool
}

var grid [][]string

func parseInput(file string) []string {
	lineArray := read.ReadStrArrayByLine(file)

	for _, line := range lineArray {
		charArray := read.StrToCharArray(line)
		grid = append(grid, charArray)
	}
	return lineArray
}

func findNumberInfo(lineArray []string) map[string]foundNumber {
	numbersMap := make(map[string]foundNumber)
	digitRegex := regexp.MustCompile(`\d+`)
	for row := 0; row < len(lineArray); row++ {
		foundNumberIndexes := digitRegex.FindAllStringIndex(lineArray[row], -1)
		foundNumbers := digitRegex.FindAllString(lineArray[row], -1)
		// fmt.Printf("Found number indexes: %v\n", foundNumberIndexes)
		// fmt.Printf("Found numbers: %v\n", foundNumbers)
		for i := 0; i < len(foundNumbers); i++ {
			intNumber, _ := strconv.Atoi(foundNumbers[i])
			foundInfo := foundNumber{
				number:       intNumber,
				indexes:      foundNumberIndexes[i],
				rowFound:     row,
				isPartNumber: false,
			}
			foundInfo.xLeft, foundInfo.xRight, foundInfo.yLow, foundInfo.yHigh = setSearchLimits(foundInfo.indexes, len(lineArray[row]), len(lineArray), row)
			val, ok := numbersMap[fmt.Sprintf("%v:%v,%v", foundInfo.number, foundInfo.indexes[0], foundInfo.rowFound)]
			if ok {
				fmt.Printf("OVERWRITING EXISTING MAP ENTRY YOU IDIOT: %v\n", val)
			}
			numbersMap[fmt.Sprintf("%v:%v,%v", foundInfo.number, foundInfo.indexes[0], foundInfo.rowFound)] = foundInfo
		}
	}
	return numbersMap
}

func setSearchLimits(numberIndex []int, xMax int, yMax int, rowFound int) (int, int, int, int) {
	xLeft := numberIndex[0] - 1
	if xLeft < 0 {
		xLeft = 0
	}
	xRight := numberIndex[1] + 2 //+2 so we can keep a plain < on loops
	if xRight > xMax {
		xRight = xMax
	}
	yLow := rowFound - 1
	if yLow < 0 {
		yLow = 0
	}
	yHigh := rowFound + 2 //+2 so we can keep a plain < on loops
	if yHigh > yMax {
		yHigh = yMax
	}
	return xLeft, xRight, yLow, yHigh
}

func checkIfPartNumber(numberInfo foundNumber) bool {
	symbolRegex := regexp.MustCompile(`[^\d.]`)
	for curY := numberInfo.yLow; curY < numberInfo.yHigh; curY++ {
		for curX := numberInfo.xLeft; curX < numberInfo.xRight; curX++ {
			if symbolRegex.MatchString(grid[curY][curX]) {
				grid[curY][curX] = fmt.Sprintf("%v", color.HiGreenString(grid[curY][curX]))
				for i := numberInfo.indexes[0]; i < numberInfo.indexes[1]; i++ {
					grid[numberInfo.rowFound][i] = fmt.Sprintf("%v", color.HiGreenString(grid[numberInfo.rowFound][i]))
				}
				// renderMap()
				// time.Sleep(5 * time.Second)
				// fmt.Printf("For number %v at %v,%v, searching grid %v,%v: %v\n", numberInfo.number, numberInfo.indexes[0], numberInfo.rowFound, curX, curY, grid[curY][curX])
				return true
			}
		}
	}
	return false
}

func renderMap() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%v", grid[y][x])
		}
		fmt.Printf("\n")
	}
}

func part1() {
	lineArray := parseInput("./input.txt")

	sum := 0

	numbersMap := findNumberInfo(lineArray)
	//fmt.Printf("%v\n", numbersMap)

	for _, numberInfo := range numbersMap {
		numberInfo.isPartNumber = checkIfPartNumber(numberInfo)
		if numberInfo.isPartNumber {
			//fmt.Printf("%v is a part number\n", numberInfo.number)
			sum += numberInfo.number
		}
	}

	renderMap()

	fmt.Printf("Part 1 Sum: %v\n", sum)
}

func part2() {}

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
