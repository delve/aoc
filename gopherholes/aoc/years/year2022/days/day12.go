package year2022

import (
	"fmt"
	"math"
	"strings"

	as "github.com/beefsack/go-astar"
	"github.com/sirupsen/logrus"
)

type Day12 struct{}

const (
	kindFrom    = -1
	kindTo      = 27
	kindPath    = 28
	kindBlocked = 29
)

const (
	kinda = iota
	kindb
	kindc
	kindd
	kinde
	kindf
	kindg
	kindh
	kindi
	kindj
	kindk
	kindl
	kindm
	kindn
	kindo
	kindp
	kindq
	kindr
	kinds
	kindt
	kindu
	kindv
	kindw
	kindx
	kindy
	kindz
)

// kindRunes map tile kinds to output runes
var kindRunes = map[int]rune{
	kindFrom:    'S',
	kindTo:      'E',
	kindPath:    '•',
	kindBlocked: '█',
	kinda:       'a',
	kindb:       'b',
	kindc:       'c',
	kindd:       'd',
	kinde:       'e',
	kindf:       'f',
	kindg:       'g',
	kindh:       'h',
	kindi:       'i',
	kindj:       'j',
	kindk:       'k',
	kindl:       'l',
	kindm:       'm',
	kindn:       'n',
	kindo:       'o',
	kindp:       'p',
	kindq:       'q',
	kindr:       'r',
	kinds:       's',
	kindt:       't',
	kindu:       'u',
	kindv:       'v',
	kindw:       'w',
	kindx:       'x',
	kindy:       'y',
	kindz:       'z',
}

// runeKindss map input runes to tile kinds
var runeKinds = map[rune]int{
	'S': kindFrom,
	'E': kindTo,
	'a': kinda,
	'b': kindb,
	'c': kindc,
	'd': kindd,
	'e': kinde,
	'f': kindf,
	'g': kindg,
	'h': kindh,
	'i': kindi,
	'j': kindj,
	'k': kindk,
	'l': kindl,
	'm': kindm,
	'n': kindn,
	'o': kindo,
	'p': kindp,
	'q': kindq,
	'r': kindr,
	's': kinds,
	't': kindt,
	'u': kindu,
	'v': kindv,
	'w': kindw,
	'x': kindx,
	'y': kindy,
	'z': kindz,
}

type tile struct {
	kind int
	// height is how tall a tile is potentially affecting movement
	height int
	// XY coords
	x, y int
	// W is a reference to the world that the tile is a part of
	w world
}

// PathNeighbors returns the neighbors of the tile, excluding blockers,
//
//	tiles off the edge of the map, and tiles too tall to reach
func (t *tile) PathNeighbors() []as.Pather {
	neighbors := []as.Pather{}
	for _, offset := range [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	} {
		n := t.w.tile(t.x+offset[0], t.y+offset[1])
		if n != nil {
			// is this neighbor blocked? and
			if n.kind != kindBlocked &&
				// is this neighbor short enough to reach?
				t.height+1 >= n.height {
				neighbors = append(neighbors, n)
			}
		}
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the neighboring tile (always 1!)
func (t *tile) PathNeighborCost(to as.Pather) float64 {
	return 1.0
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
//
//	between non-adjacent tiles
func (t *tile) PathEstimatedCost(to as.Pather) float64 {
	toT := to.(*tile)
	absX := toT.x - t.x
	if absX < 0 {
		absX = -absX
	}
	absY := toT.y - t.y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

// world is a 2d map of tiles
type world map[int]map[int]*tile

// tile gets the tile at the given coords in the world
func (w world) tile(x, y int) *tile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

// setTile sets a tile at the given coords in the world
func (w world) setTile(t *tile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*tile{}
	}
	w[x][y] = t
	t.x = x
	t.y = y
	t.w = w
}

// firstOfKind gets the first tile in the world of the given kind
// used to get the start & end tiles, as there should only be 1 of each
func (w world) firstOfKind(kind int) *tile {
	for _, row := range w {
		for _, t := range row {
			if t.kind == kind {
				return t
			}
		}
	}
	return nil
}

// from gets the start tile from the world
func (w world) from() *tile {
	return w.firstOfKind(kindFrom)
}

// to gets the end tile from the world
func (w world) to() *tile {
	return w.firstOfKind(kindTo)
}

// renderPath renders a path on top of a world
//
//	assumes a simple rectangle shape
func (w world) renderPath(path []as.Pather) string {
	width := len(w)
	if width == 0 {
		return ""
	}
	height := len(w[0])
	pathLocs := map[string]bool{}
	for _, p := range path {
		pT := p.(*tile)
		pathLocs[fmt.Sprintf("%d,%d", pT.x, pT.y)] = true
	}
	rows := make([]string, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			t := w.tile(x, y)
			r := ' '
			if pathLocs[fmt.Sprintf("%d,%d", x, y)] {
				r = kindRunes[kindPath]
			} else if t != nil {
				r = kindRunes[t.kind]
			}
			rows[y] += string(r)
		}
	}
	return strings.Join(rows, "\n")
}

// parseInput parses the textual representation of a world in the input into a map
func parseInput12(lines []string) world {
	w := world{}
	for y, row := range lines {
		for x, raw := range row {
			kind, ok := runeKinds[raw]
			if !ok {
				kind = kindBlocked
			}
			newTile := tile{kind: kind}
			switch kind {
			case kindFrom:
				newTile.height = kinda
			case kindTo:
				newTile.height = kindz
			default:
				newTile.height = kind
			}
			w.setTile(&newTile, x, y)
		}
	}

	return w
}

func (p Day12) PartA(lines []string) any {
	w := parseInput12(lines)
	// fmt.Printf("Input world:\n%s", w.renderPath([]as.Pather{}))
	_, dist, found := as.Path(w.from(), w.to())
	if !found {
		logrus.Fatal("Could not find a path")
	}

	return dist
}

func (p Day12) PartB(lines []string) any {
	w := parseInput12(lines)

	var starters []*tile
	for _, row := range w {
		for _, cell := range row {
			if cell.height == kinda {
				starters = append(starters, cell)
			}
		}
	}

	length := math.MaxFloat64
	// var startCell *tile
	var solvePath []as.Pather

	for _, cell := range starters {
		w.from().kind = kinda
		cell.kind = kindFrom
		path, dist, found := as.Path(w.from(), w.to())
		if found && dist < length {
			length = dist
			// startCell = cell
			solvePath = path
		}
	}
	fmt.Printf("solved part 2:\n%s", w.renderPath(solvePath))

	return length
}
