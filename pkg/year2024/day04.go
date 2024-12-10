package year2024

import (
	"math"
)

type Day04 struct{}

func row(c complex128) int {
	return int(real(c))
}

func column(c complex128) int {
	return int(imag(c))
}

// length and bounds are one-based, position is xero-based
func checkBounds(direction rune, position complex128, length int, bounds complex128) bool {
	// 0-base the inputs
	ln := length - 1
	height := row(bounds) - 1
	width := column(bounds) - 1
	if direction == 'U' {
		if row(position) < ln {
			// not enough room up
			return false
		}
	}
	if direction == 'D' {
		if row(position) > height-ln {
			// not enough room down
			return false
		}
	}
	if direction == 'L' {
		if column(position) < ln {
			// not enough room left
			return false
		}
	}
	if direction == 'R' {
		if column(position) > width-ln {
			// not enough room right
			return false
		}
	}

	return true
}

// clock shaped directions
func wordSearchUp(input []string, start complex128, keyword string) bool {
	// bounds checking
	if !checkBounds('U', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}

	for idx, letter := range keyword {
		if input[row(start)-idx][column(start)] != byte(letter) {
			return false
		}
	}
	return true
}

func wordSearchRight(input []string, start complex128, keyword string) bool {
	// bounds checking
	if !checkBounds('R', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}

	for idx, letter := range keyword {
		if input[row(start)][column(start)+idx] != byte(letter) {
			return false
		}
	}
	return true
}

func wordSearchUpRight(input []string, start complex128, keyword string) bool {
	// bounds checking
	if !checkBounds('U', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}
	if !checkBounds('R', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}

	for idx, letter := range keyword {
		if input[row(start)-idx][column(start)+idx] != byte(letter) {
			return false
		}
	}

	return true
}

func wordSearchDownRight(input []string, start complex128, keyword string) bool {
	// bounds checking
	if !checkBounds('D', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}
	if !checkBounds('R', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}

	for idx, letter := range keyword {
		if input[row(start)+idx][column(start)+idx] != byte(letter) {
			return false
		}
	}

	return true
}

func wordSearchDown(input []string, start complex128, keyword string) bool {
	// bounds checking
	if !checkBounds('D', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}

	for idx, letter := range keyword {
		if input[row(start)+idx][column(start)] != byte(letter) {
			return false
		}
	}

	return true
}

func wordSearchDownLeft(input []string, start complex128, keyword string) bool {
	// bounds checking
	if !checkBounds('D', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}
	if !checkBounds('L', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}

	for idx, letter := range keyword {
		if input[row(start)+idx][column(start)-idx] != byte(letter) {
			return false
		}
	}

	return true
}

func wordSearchLeft(input []string, start complex128, keyword string) bool {
	// bounds checking
	if !checkBounds('L', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}

	for idx, letter := range keyword {
		if input[row(start)][column(start)-idx] != byte(letter) {
			return false
		}
	}

	return true
}

func wordSearchUpLeft(input []string, start complex128, keyword string) bool {
	// bounds checking
	if !checkBounds('U', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}
	if !checkBounds('L', start, len(keyword), complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}

	for idx, letter := range keyword {
		if input[row(start)-idx][column(start)-idx] != byte(letter) {
			return false
		}
	}

	return true
}

func wordCrossCheck(input []string, position complex128, keyword string) bool {
	// add one to leglenth or bounds checking fails
	legLength := int(math.Floor(float64(len(keyword))/2)) + 1
	topMas := false
	botMas := false

	// bounds checking
	if !checkBounds('U', position, legLength, complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}
	if !checkBounds('D', position, legLength, complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}
	if !checkBounds('L', position, legLength, complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}
	if !checkBounds('R', position, legLength, complex(float64(len(input)), float64(len(input[0])))) {
		return false
	}

	if wordSearchDownRight(input, position-1-1i, keyword) || wordSearchUpLeft(input, position+1+1i, keyword) {
		topMas = true
	}

	if wordSearchUpRight(input, position+1-1i, keyword) || wordSearchDownLeft(input, position-1+1i, keyword) {
		botMas = true
	}

	if topMas && botMas {
		return true
	} else {
		return false
	}
}

func (p Day04) PartA(lines []string) any {
	input := lines[:len(lines)-1]
	hits := 0
	keyword := "XMAS"
	for row, data := range input {
		for column, letter := range data {
			if letter == rune(keyword[0]) {
				position := complex(float64(row), float64(column))
				if wordSearchUp(input, position, keyword) {
					hits++
				}

				if wordSearchUpRight(input, position, keyword) {
					hits++
				}

				if wordSearchRight(input, position, keyword) {
					hits++
				}

				if wordSearchDownRight(input, position, keyword) {
					hits++
				}

				if wordSearchDown(input, position, keyword) {
					hits++
				}

				if wordSearchDownLeft(input, position, keyword) {
					hits++
				}

				if wordSearchLeft(input, position, keyword) {
					hits++
				}

				if wordSearchUpLeft(input, position, keyword) {
					hits++
				}
			}
		}
	}

	return hits
}

func (p Day04) PartB(lines []string) any {
	input := lines[:len(lines)-1]
	hits := 0
	keyword := "MAS"
	for row, data := range input {
		for column, letter := range data {
			if letter == 'A' {
				position := complex(float64(row), float64(column))
				if wordCrossCheck(input, position, keyword){
					hits++
				}
			}
		}
	}

	return hits

}
