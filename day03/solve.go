package day03

import "fmt"

var debug bool

func solvePart1(schematic schematic) int {
	sum, err := schematic.sumOfParts()
	if err != nil {
		panic(err)
	}
	return sum
}

func solvePart2(schematic schematic) int {
	sum, err := schematic.sumOfGearRatios()
	if err != nil {
		panic(err)
	}
	return sum
}

func Solve(inputFileName string, d bool) error {
	debug = d
	schematic, err := parse(inputFileName)
	if err != nil {
		return err
	}
	fmt.Printf("Part one: %v\n", solvePart1(schematic))
	fmt.Printf("Part one: %v\n", solvePart2(schematic))
	return nil
}
