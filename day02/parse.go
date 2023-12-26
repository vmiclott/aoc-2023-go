package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse(inputFileName string) ([]game, error) {
	readFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var games = []game{}
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, err
		}
		game, err := parseGame(line)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}

func parseGame(line string) (game, error) {
	id, err := parseGameId(line)
	if err != nil {
		return game{}, err
	}
	rounds, err := parseRounds(line)
	if err != nil {
		return game{}, err
	}
	return game{id: id, rounds: rounds}, nil
}

func parseGameId(line string) (int, error) {
	id := strings.Split(strings.Split(line, ":")[0], " ")[1]
	return strconv.Atoi(id)
}

func parseRounds(line string) ([]round, error) {
	roundStrings := strings.Split(strings.Split(line, ":")[1], ";")
	rounds := []round{}
	for _, roundString := range roundStrings {
		round, err := parseRound(roundString)
		if err != nil {
			return nil, err
		}
		rounds = append(rounds, round)
	}
	return rounds, nil
}

func parseRound(roundString string) (round, error) {
	cubes := strings.Split(roundString, ",")
	round := round{}
	for _, cube := range cubes {
		cube = strings.Trim(cube, " ")
		parts := strings.Split(cube, " ")
		count, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		round[colorMap[parts[1]]] = count
	}
	return round, nil
}
