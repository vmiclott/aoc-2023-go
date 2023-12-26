package day03

import (
	"bufio"
	"os"
	"strings"
)

func parse(inputFileName string) (schematic, error) {
	readFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var schematic = schematic{}
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, err
		}
		schematic = append(schematic, strings.Split(line, ""))
	}

	return schematic, nil
}
