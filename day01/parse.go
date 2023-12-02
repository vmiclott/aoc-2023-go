package day01

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse(inputFileName string, isPart2 bool) ([]int, error) {
	readFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var calibrationValues = []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, err
		}
		if isPart2 {
			if debug {
				fmt.Printf("line before converting words %s\n", line)
			}
			line = convertDigitWordsToDigits(line)
			if debug {
				fmt.Printf("line after converting words %s\n", line)
			}
		}
		calibrationValue, err := parseCalibrationValue(line)
		if err != nil {
			return nil, err
		}
		calibrationValues = append(calibrationValues, calibrationValue)
	}

	return calibrationValues, nil
}

var digits = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

var wordsToDigits = [][]string{
	{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}, {
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	},
}

func convertDigitWordsToDigits(s string) string {
	words := wordsToDigits[0]
	digits := wordsToDigits[1]
	i := 0
	for i < len(s) {
		suffix := s[i:]
		for j := range words {
			if strings.HasPrefix(suffix, words[j]) {
				s = s[:i] + digits[j] + s[i+1:]
			}
		}
		i++
	}
	return s
}

func parseCalibrationValue(calibrationValueLine string) (int, error) {
	firstDigit, err := findFirstDigit(calibrationValueLine)
	if err != nil {
		return -1, err
	}
	lastDigit, err := findLastDigit(calibrationValueLine)
	if err != nil {
		return -1, err
	}
	return 10*firstDigit + lastDigit, nil
}

func findFirstDigit(s string) (int, error) {
	for i := range s {
		if digit, ok := digits[s[i]]; ok {
			return digit, nil
		}
	}
	return -1, fmt.Errorf("string %s does not contain a digit", s)
}

func findLastDigit(s string) (int, error) {
	for i := len(s) - 1; i >= 0; i-- {
		if digit, ok := digits[s[i]]; ok {
			return digit, nil
		}
	}
	return -1, fmt.Errorf("string %s does not contain a digit", s)
}
