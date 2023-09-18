package year2022

import (
	"aocgen/pkg/common"
	"container/ring"
)

type Day20 struct{}

func parseInput20(lines []string, encKey int) *ring.Ring {

	var inputs []int
	for _, line := range lines {
		inputs = append(inputs, common.Atoi(line))
	}

	r := ring.New(len(inputs))
	for i, v := range inputs {
		r.Value = [2]int{i, v * encKey}
		r = r.Next()
	}
	return r
}

func findIdx(r *ring.Ring, i int) *ring.Ring {
	for r.Value.([2]int)[0] != i {
		r = r.Next()
	}
	return r
}

func getCoords(r *ring.Ring) (int, int, int) {
	var a, b, c int
	for i := 1; i <= 3000; i++ {
		r = r.Next()
		// a @ 1000
		if i == 1000 {
			a = r.Value.([2]int)[1]
		}
		// b @ 2000
		if i == 2000 {
			b = r.Value.([2]int)[1]
		}
		// c @ 3000
		if i == 3000 {
			c = r.Value.([2]int)[1]
		}
	}

	return a, b, c
}

func shuffleRing(r *ring.Ring) *ring.Ring {
	for i := 0; i < r.Len(); i++ {
		r = findIdx(r, i)
		// back it up one because unlink takes r.next
		r = r.Prev()
		q := r.Unlink(1)
		// move back to the original index
		r = r.Next()
		distance := q.Value.([2]int)[1] % r.Len()
		// Link() sets r.next
		r = r.Move(distance - 1)
		r = r.Link(q)
	}
	return r
}

func (p Day20) PartA(lines []string) any {
	r := parseInput20(lines, 1)
	r = shuffleRing(r)
	// move to -value- 0
	for r.Value.([2]int)[1] != 0 {
		r = r.Next()
	}

	// get the 3 digits
	a, b, c := getCoords(r)
	return a + b + c
}

func (p Day20) PartB(lines []string) any {
	r := parseInput20(lines, 811589153)
	// 10 rounds of mixing
	for i := 0; i < 10; i++ {
		r = shuffleRing(r)
	}

	// move to -value- 0
	for r.Value.([2]int)[1] != 0 {
		r = r.Next()
	}

	// get the 3 digits
	a, b, c := getCoords(r)
	return a + b + c
}
