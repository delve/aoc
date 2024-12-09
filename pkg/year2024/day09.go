package year2024

import (
	"aocgen/pkg/common"
)

type Day09 struct {
	diskMap []int
	files   map[int][2]int // [id][position,size]
}

func (p *Day09) parseInput(lines []string) {
	p.files = map[int][2]int{}
	if len(lines) > 1 {
		panic("Input too long")
	}
	input := lines[0]
	isFile := true
	fileIdx := 0
	diskIdx := 0
	for _, length := range input {
		intLen := common.Atoi(string(length))
		if isFile {
			if intLen < 1 {
				panic("filelength 0")
			}
			p.files[fileIdx] = [2]int{intLen, diskIdx}
			for i := 0; i < intLen; i++ {
				p.diskMap = append(p.diskMap, fileIdx)
			}
			fileIdx++
		} else {
			for i := 0; i < intLen; i++ {
				p.diskMap = append(p.diskMap, -1)
			}
		}
		diskIdx += intLen
		isFile = !isFile
	}
}

func (p Day09) checksum() int {
	checksum := 0
	for idx, val := range p.diskMap {
		if val >= 0 {
			checksum += idx * val
		}
	}
	return checksum
}

func (p *Day09) compact() {
	result := make([]int, len(p.diskMap))
	// copy(result, p.diskMap)
	movePointer := len(p.diskMap) - 1
	for idx, val := range p.diskMap {
		if idx > movePointer {
			for i := idx; i < len(result); i++ {
				result[i] = -1
			}
			break
		}
		if idx == movePointer {
			result[idx] = p.diskMap[idx]
			for i := idx + 1; i < len(result); i++ {
				result[i] = -1
			}
			break
		}
		if val < 0 {
			result[idx] = p.diskMap[movePointer]
			movePointer--
			for p.diskMap[movePointer] < 0 {
				movePointer--
			}
		} else {
			result[idx] = val
		}
	}
	p.diskMap = result
}

func (p *Day09) noFragCompact() {
	result := make([]int, len(p.diskMap))
	copy(result, p.diskMap)
	// fmt.Printf("%v\n", result)
	for fileId := len(p.files) - 1; fileId > -1; fileId-- {
		fileSize, fileLoc := p.files[fileId][0], p.files[fileId][1]
		scanIdx := 0
		found := false
		scanLimit := fileLoc - fileSize + 1
	noGap:
		for scanIdx = 0; scanIdx < scanLimit; scanIdx++ {
			for j := 0; j < fileSize; j++ {
				if result[scanIdx+j] > -1 {
					scanIdx += j
					continue noGap
				}
			}
			found = true
			break
		}
		if found {
			for i := 0; i < fileSize; i++ {
				result[scanIdx+i] = fileId
				result[fileLoc+i] = -1
			}
		}
		// fmt.Printf("%v\n", result)

	}
	p.diskMap = result
}

func (p Day09) PartA(lines []string) any {
	p.parseInput(lines[0:1])
	// fmt.Printf("%v\n", p.diskMap)
	// fmt.Printf("%v\n", p.files)
	p.compact()
	// fmt.Printf("%v", p.diskMap)
	return p.checksum()
}

func (p Day09) PartB(lines []string) any {
	p.parseInput(lines[0:1])
	// fmt.Printf("%v\n", p.diskMap)
	// fmt.Printf("%v\n", p.files)
	p.noFragCompact()
	// fmt.Printf("%v", p.diskMap)
	return p.checksum()
}
