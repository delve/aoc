package year2024

type Day04 struct{}

func row(c complex128) int {
	return int(real(c))
}

func column(c complex128) int {
	return int(imag(c))
}

// clock shaped directions
func wordSearchUp(input []string, start complex128, keyword string) bool {
	// bounds checking
	if row(start) < len(keyword)-1 {
		// not enough room for vertical xmas
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
	if column(start) > len(input[row(start)])-len(keyword) {
		// not enough room for horizontal xmas
		return false
	}

	for idx, letter := range keyword {
		if input[row(start)][column(start)-idx] != byte(letter) {
			return false
		}
	}
	return true
}

func wordSearchUpRight(input []string, start complex128, keyword string) bool {
	return true
}

func wordSearchDownright(input []string, start complex128, keyword string) bool {
	return false
}

func wordSearchDown(input []string, start complex128, keyword string) bool {
	return false
}

func wordSearchDownLeft(input []string, start complex128, keyword string) bool {
	return false
}

func wordSearchLeft(input []string, start complex128, keyword string) bool {
	return false
}

func wordSearchUpLeft(input []string, start complex128, keyword string) bool {
	return false
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

				if wordSearchDownright(input, position, keyword) {
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

	return "implement_me"
}

func (p Day04) PartB(lines []string) any {
	return "implement_me"
}
