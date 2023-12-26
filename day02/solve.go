package day02

import "fmt"

var debug bool

func solvePart1(games []game) int {
	sum := 0
	for _, game := range games {
		if game.isPossible(12, 13, 14) {
			sum += game.id
		}
	}
	return sum
}

func solvePart2(games []game) int {
	sum := 0
	for _, game := range games {
		sum += game.power()
	}
	return sum
}

func Solve(inputFileName string, d bool) error {
	debug = d
	games, err := parse(inputFileName)
	if err != nil {
		return err
	}
	fmt.Printf("Part one: %v\n", solvePart1(games))
	fmt.Printf("Part one: %v\n", solvePart2(games))
	return nil
}
