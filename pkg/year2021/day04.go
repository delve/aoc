package year2021

import (
	"aocgen/pkg/common"
	"fmt"
	"strings"
)

type Day04 struct {
	calls  []int
	boards map[string]bingoBoard
}

type bingoBoard struct {
	bingo  bool
	spaces [5][5]struct {
		number int
		called bool
	}
}

func (b *bingoBoard) testBoard() bool {
	if b.bingo {
		return true
	}

	for i := 0; i < 5; i++ {
		if b.spaces[i][0].called && b.spaces[i][1].called && b.spaces[i][2].called && b.spaces[i][3].called && b.spaces[i][4].called {
			b.bingo = true
		}
		if b.spaces[0][i].called && b.spaces[1][i].called && b.spaces[2][i].called && b.spaces[3][i].called && b.spaces[4][i].called {
			b.bingo = true
		}
	}

	return b.bingo
}

func (b bingoBoard) call(search int) bingoBoard {
	hit := false
	for i := 0; i < len(b.spaces); i++ {
		for j := 0; j < len(b.spaces[i]); j++ {
			if b.spaces[i][j].number == search {
				b.spaces[i][j].called = true
				hit = true
			}
		}
	}

	if hit {
		b.testBoard()
	}
	return b
}

func (b bingoBoard) String() string {
	isCalled := func(b bool) string {
		if b {
			return "X"
		} else {
			return " "
		}
	}
	output := "Has bingoed: "
	if b.bingo {
		output += "Yes\n"
	} else {
		output += "No\n"
	}

	for row := 0; row < len(b.spaces); row++ {
		for column := 0; column < len(b.spaces[row]); column++ {
			output += fmt.Sprintf("%2d:%s|", b.spaces[row][column].number, isCalled(b.spaces[row][column].called))
		}
		output += "\n"
	}
	return output
}

func (b bingoBoard) calcScore(winningNumber int) int {
	boardScore := 0
	for row := 0; row < len(b.spaces); row++ {
		for column := 0; column < len(b.spaces[row]); column++ {
			if !b.spaces[row][column].called {
				boardScore += b.spaces[row][column].number
			}
		}
	}
	return boardScore * winningNumber
}

func (p *Day04) parseBingo(lines []string) {
	p.boards = map[string]bingoBoard{}
	nums := strings.Split(lines[0], ",")
	for _, number := range nums {
		p.calls = append(p.calls, common.Atoi(number))
	}
	boardData := lines[2 : len(lines)-1]
	space := func(c rune) bool {
		return c == ' '
	}
	boardId := 0
	for i := 0; i < len(boardData); { //explicitly incrementing inside loop
		board := bingoBoard{}
		for row := 0; row < 5; row++ {
			nums := strings.FieldsFunc(boardData[i], space)
			for column, number := range nums {
				board.spaces[row][column].number = common.Atoi(number)
			}
			i++ // next board line
		}
		p.boards[fmt.Sprintf("%d", boardId)] = board
		boardId++
		i++ // skip the blank line
	}
}

func (p Day04) PartA(lines []string) any {
	p.parseBingo(lines)

	winningNumber := -1
	winningBoard := ""
	for _, callNumber := range p.calls {
		for boardId := range p.boards {
			p.boards[boardId] = p.boards[boardId].call(callNumber)
			if p.boards[boardId].bingo {
				winningNumber = callNumber
				winningBoard = boardId
			}
		}
		if winningNumber > -1 {
			break
		}
	}
	for boardId := range p.boards {
		fmt.Printf("%v\n", p.boards[boardId])	
	}
	 
	return fmt.Sprintf("%v", p.boards[winningBoard].calcScore(winningNumber))
}

func (p Day04) PartB(lines []string) any {
	p.parseBingo(lines)

	var winningNumber int
	var winningBoard bingoBoard
	bingoedBoards := []string{}
	for _, callNumber := range p.calls {
		for boardId := range p.boards {
			p.boards[boardId] = p.boards[boardId].call(callNumber)
			if p.boards[boardId].bingo {
				// fmt.Printf("Board ID %d has won on callnumber %d\n", p.boards[i].id, callNumber)
				bingoedBoards = append(bingoedBoards, boardId)
				winningNumber = callNumber
				winningBoard = p.boards[boardId]
			}
		}
		for _, boardId := range bingoedBoards {
			if _, ok := p.boards[boardId]; !ok {
				delete(p.boards, boardId)
			}
		}

	}
	fmt.Printf("%v\n%v", winningNumber, winningBoard)

	return fmt.Sprintf("%v", winningBoard.calcScore(winningNumber))
}
