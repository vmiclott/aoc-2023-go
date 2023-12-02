package day01

import "fmt"

var debug bool

func solve(calibrationValues []int) int {
	sum := 0
	for _, calibrationValue := range calibrationValues {
		sum += calibrationValue
	}

	return sum
}

func Solve(inputFileName string, d bool) error {
	debug = d
	calibrationValuesPart1, err := parse(inputFileName, false)
	if err != nil {
		return err
	}
	fmt.Printf("Part one: %v\n", solve(calibrationValuesPart1))

	calibrationValuesPart2, err := parse(inputFileName, true)
	if err != nil {
		return err
	}
	fmt.Printf("Part two: %v\n", solve(calibrationValuesPart2))
	return nil
}
