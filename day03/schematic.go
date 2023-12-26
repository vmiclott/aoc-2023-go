package day03

import (
	"fmt"
	"strconv"
)

type schematic [][]string

func (schematic schematic) sumOfParts() (int, error) {
	height := len(schematic)
	if height < 1 {
		return 0, fmt.Errorf("schematic has no rows")
	}
	width := len(schematic[0])
	if width < 1 {
		return 0, fmt.Errorf("schematic has no columns")
	}

	hasNumber := false
	currNumber := 0
	isPart := false
	sum := 0
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			c := schematic[h][w]
			if debug {
				fmt.Printf("[%d, %d]: %s\n", h, w, c)
			}
			if c == "." {
				if debug {
					fmt.Println("handling .")
				}
				if hasNumber && isPart {
					if debug {
						fmt.Printf("has number and is part increasing sum from %d to %d\n", sum, sum+currNumber)
					}
					sum += currNumber
				}
				hasNumber = false
				isPart = false
				currNumber = 0
				continue
			}

			if digit, isDigit := schematic.getDigit(h, w); isDigit {
				if debug {
					fmt.Println("handling digit")
				}
				if hasNumber {
					if debug {
						fmt.Printf("has number increasing curr number from %d to %d\n", currNumber, currNumber*10+digit)
					}
					currNumber = currNumber*10 + digit
				} else {
					if debug {
						fmt.Printf("has no number setting curr number to %d\n", digit)
					}
					hasNumber = true
					currNumber = digit
				}
				if !isPart {
					isPart = schematic.posIsPart(h, w)
				}
			} else {
				if debug {
					fmt.Println("handling special character")
				}
				if hasNumber && isPart {
					if debug {
						fmt.Printf("has number and is part increasing sum from %d to %d\n", sum, sum+currNumber)
					}
					sum += currNumber
				}
				hasNumber = false
				isPart = false
				currNumber = 0
			}
		}
		if hasNumber && isPart {
			if debug {
				fmt.Printf("has number and is part increasing sum from %d to %d\n", sum, sum+currNumber)
			}
			sum += currNumber
			hasNumber = false
			isPart = false
			currNumber = 0
		}
	}
	return sum, nil
}

func (schematic schematic) sumOfGearRatios() (int, error) {
	height := len(schematic)
	if height < 1 {
		return 0, fmt.Errorf("schematic has no rows")
	}
	width := len(schematic[0])
	if width < 1 {
		return 0, fmt.Errorf("schematic has no columns")
	}

	sum := 0
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if schematic[h][w] == "*" {
				sum += schematic.gearRatio(h, w)
			}
		}
	}

	return sum, nil
}

func (schematic schematic) getDigit(h int, w int) (int, bool) {
	if !schematic.posIsValid(h, w) {
		return 0, false
	}
	c := schematic[h][w]
	digit, err := strconv.Atoi(c)
	if err != nil {
		return 0, false
	}
	return digit, true
}

func (schematic schematic) gearRatio(h int, w int) int {
	ratio := 1
	partsCount := 0
	// row above
	if _, isDigit := schematic.getDigit(h+1, w); isDigit {
		partsCount++
		ratio *= schematic.getPartAtPos(h+1, w)
	} else {
		if _, isDigit := schematic.getDigit(h+1, w-1); isDigit {
			partsCount++
			ratio *= schematic.getPartAtPos(h+1, w-1)
		}
		if _, isDigit := schematic.getDigit(h+1, w+1); isDigit {
			partsCount++
			ratio *= schematic.getPartAtPos(h+1, w+1)
		}
	}

	// same row
	if _, isDigit := schematic.getDigit(h, w-1); isDigit {
		partsCount++
		ratio *= schematic.getPartAtPos(h, w-1)
	}
	if _, isDigit := schematic.getDigit(h, w+1); isDigit {
		partsCount++
		ratio *= schematic.getPartAtPos(h, w+1)
	}

	// row below
	if _, isDigit := schematic.getDigit(h-1, w); isDigit {
		partsCount++
		ratio *= schematic.getPartAtPos(h-1, w)
	} else {
		if _, isDigit := schematic.getDigit(h-1, w-1); isDigit {
			partsCount++
			ratio *= schematic.getPartAtPos(h-1, w-1)
		}
		if _, isDigit := schematic.getDigit(h-1, w+1); isDigit {
			partsCount++
			ratio *= schematic.getPartAtPos(h-1, w+1)
		}
	}

	if partsCount == 2 {
		if debug {
			fmt.Printf("gear at pos [%d, %d] with ratio %d\n", h, w, ratio)
		}
		return ratio
	}

	if debug {
		fmt.Printf("not gear at pos [%d, %d] because %d parts\n", h, w, partsCount)
	}
	return 0
}

func (schematic schematic) getPartAtPos(h int, w int) int {
	width := len(schematic[h])
	digit, _ := schematic.getDigit(h, w)
	digits := []int{digit}
	for j := w - 1; j >= 0; j-- {
		if digit, isDigit := schematic.getDigit(h, j); isDigit {
			digits = append([]int{digit}, digits...)
		} else {
			break
		}
	}
	for j := w + 1; j < width; j++ {
		if digit, isDigit := schematic.getDigit(h, j); isDigit {
			digits = append(digits, digit)
		} else {
			break
		}
	}
	factor := 1
	total := 0
	for i := len(digits) - 1; i >= 0; i-- {
		total += factor * digits[i]
		factor *= 10
	}
	return total
}

func (schematic schematic) posIsPart(h int, w int) bool {
	for _, neighbor := range schematic.getNeighbors(h, w) {
		c := schematic[neighbor[0]][neighbor[1]]
		_, isDigit := schematic.getDigit(neighbor[0], neighbor[1])
		if c != "." && !isDigit {
			return true
		}
	}
	return false
}

func (schematic schematic) getNeighbors(h int, w int) [][2]int {
	neighbors := [][2]int{
		{h - 1, w - 1},
		{h - 1, w},
		{h - 1, w + 1},
		{h, w - 1},
		{h, w + 1},
		{h + 1, w - 1},
		{h + 1, w},
		{h + 1, w + 1},
	}
	validNeighbors := [][2]int{}
	for _, pos := range neighbors {
		if schematic.posIsValid(pos[0], pos[1]) {
			validNeighbors = append(validNeighbors, pos)
		}
	}
	return validNeighbors
}

func (schematic schematic) posIsValid(h int, w int) bool {
	height := len(schematic)
	width := len(schematic[0])
	return h >= 0 && h < height && w >= 0 && w < width
}
