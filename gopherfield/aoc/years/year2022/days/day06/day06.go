package year2022

import (
	"aocgen/pkg/common"
	"bufio"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/golang-collections/collections/stack"
)

type Day06 struct{}

// returns the position of the beginning of a packet of given length in data
func findPacketStart(data string, length int) int {
	beginningOfPacket := 0

	byteCount := len(data)

	reader := bufio.NewReader(strings.NewReader(data))
	// noise data with matches
	var noise stack.Stack
	// size of signal
	// for part 2

	for i := 0; i < byteCount; i++ {
		window := mapset.NewSet[byte]()
		dupe := false
		char, err := reader.ReadByte()
		common.Check(err)
		window.Add(char)
		// look ahead the marker length without advancing, opportunistic match
		lookahead, err := reader.Peek(length - 1)
		common.Check(err)
		for i := 0; i < len(lookahead); i++ {
			if !window.Add(lookahead[i]) {
				dupe = true
			}

		}
		if dupe {
			noise.Push(char)
		} else {
			break
		}
	}

	beginningOfPacket = noise.Len()
	return beginningOfPacket
}

func (p Day06) PartA(lines []string) any {
	markerLen := 4
	endOfFirstPacket := findPacketStart(lines[0], markerLen) + markerLen
	return endOfFirstPacket
}

func (p Day06) PartB(lines []string) any {
	markerLen := 14
	endOfFirstPacket := findPacketStart(lines[0], markerLen) + markerLen
	return endOfFirstPacket
}
