package year2022

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Day08 struct{}

type tree struct {
	height                            int
	left, top, right, bottom, visible bool
	scenicScore                       int
}

func parseInput08(lines []string) [][]tree {
	grid := [][]tree{}

	for _, line := range lines {
		inLine := strings.Split(line, "")
		var row []tree
		for _, sHeight := range inLine {
			if iHeight, err := strconv.Atoi(sHeight); err == nil {
				row = append(row, tree{height: iHeight, visible: false})
			} else {
				logrus.Fatalf("can't convert %s to int", sHeight)
			}
		}
		grid = append(grid, row)
	}

	// set visibility
	for y, row := range grid {
		for x, element := range row {
			sceneAccum := 1
			// left visible?
			if x == 0 {
				sceneAccum = sceneAccum * 0
				grid[y][x].left = true
				grid[y][x].visible = true
			} else {
				for q := x - 1; q > -1; q-- {
					if grid[y][q].height >= element.height {
						sceneAccum = sceneAccum * (x - q)
						break
					}
					if q == 0 {
						grid[y][x].left = true
						grid[y][x].visible = true
						sceneAccum = sceneAccum * (x - q)
					}
				}
			}

			// top visible?
			if y == 0 {
				sceneAccum = sceneAccum * 0
				grid[y][x].top = true
				grid[y][x].visible = true
			} else {
				for q := y - 1; q > -1; q-- {
					if grid[q][x].height >= element.height {
						sceneAccum = sceneAccum * (y - q)
						break
					}
					if q == 0 {
						grid[y][x].top = true
						grid[y][x].visible = true
						sceneAccum = sceneAccum * (y - q)
					}
				}
			}

			// right visible?
			if x == len(grid[0])-1 {
				sceneAccum = sceneAccum * 0
				grid[y][x].right = true
				grid[y][x].visible = true
			} else {
				for q := x + 1; q < len(grid[0]); q++ {
					if grid[y][q].height >= element.height {
						sceneAccum = sceneAccum * (q - x)
						break
					}
					if q == len(grid[0])-1 {
						grid[y][x].right = true
						grid[y][x].visible = true
						sceneAccum = sceneAccum * (q - x)
					}
				}
			}

			// bottom visible?
			if y == len(grid)-1 {
				sceneAccum = sceneAccum * 0
				grid[y][x].bottom = true
				grid[y][x].visible = true
			} else {
				for q := y + 1; q < len(grid); q++ {
					if grid[q][x].height >= element.height {
						sceneAccum = sceneAccum * (q - y)
						break
					}
					if q == len(grid)-1 {
						grid[y][x].bottom = true
						grid[y][x].visible = true
						sceneAccum = sceneAccum * (q - y)
					}
				}
			}
			grid[y][x].scenicScore = sceneAccum
		}
	}
	return grid
}

func (p Day08) PartA(lines []string) any {
	grid := parseInput08(lines)
	visibleTrees := 0

	for _, row := range grid {
		for _, e := range row {
			if e.visible {
				visibleTrees++
			}
		}
	}

	return visibleTrees
}

func (p Day08) PartB(lines []string) any {
	grid := parseInput08(lines)
	topScore := 0

	for _, row := range grid {
		for _, e := range row {
			if e.scenicScore > topScore {
				topScore = e.scenicScore
			}
		}
	}

	return topScore
}
