package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/SotoDucani/AoC2022/internal/read"
)

func part1() {
	games := read.ReadStrArrayByLine("./input.txt")

	gameRegex := regexp.MustCompile(`^Game (\d+)`)
	diceRegex := regexp.MustCompile(`(\d+) (red|green|blue)`)

	allowedMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sumOfIds := 0

	for _, game := range games {
		gameId := gameRegex.FindStringSubmatch(game)[1]
		//fmt.Printf("Game ID is: %s\n", gameId)
		gameIdNum, _ := strconv.Atoi(gameId)
		diceArray := diceRegex.FindAllStringSubmatch(game, -1)
		gameAllowed := true
		for _, matchedDice := range diceArray {
			//fmt.Printf("Dice Info: %s\n", matchedDice[0])
			diceNum, _ := strconv.Atoi(matchedDice[1])
			if allowedMap[matchedDice[2]] < diceNum {
				gameAllowed = false
				// We know it's bad, no need to continue search
				break
			}
		}
		if gameAllowed {
			// If we get here, game was all allowed
			//fmt.Printf("Game %s was allowed\n", gameId)
			sumOfIds += gameIdNum
		}
	}

	fmt.Printf("Part 1 sum: %v\n", sumOfIds)
}

func part2() {
	games := read.ReadStrArrayByLine("./input.txt")

	gameRegex := regexp.MustCompile(`^Game (\d+)`)
	diceRegex := regexp.MustCompile(`(\d+) (red|green|blue)`)

	sumOfPowers := 0

	for _, game := range games {
		maxMap := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		gameId := gameRegex.FindStringSubmatch(game)[1]
		//fmt.Printf("Game ID is: %s\n", gameId)
		diceArray := diceRegex.FindAllStringSubmatch(game, -1)
		for _, matchedDice := range diceArray {
			//fmt.Printf("Dice Info: %s\n", matchedDice[0])
			diceNum, _ := strconv.Atoi(matchedDice[1])
			if maxMap[matchedDice[2]] < diceNum {
				maxMap[matchedDice[2]] = diceNum
			}
		}
		power := 1
		for _, num := range maxMap {
			power = power * num
		}
		fmt.Printf("Power of game %s is %v\n", gameId, power)
		sumOfPowers += power
	}

	fmt.Printf("Part 2 sum: %v\n", sumOfPowers)
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
