package year2022

import (
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sirupsen/logrus"
)

type Day03 struct{}

type rucksack struct {
	contents         string
	pocket1, pocket2 string
	mistakeItem      rune
	mistakePriority  int
}

type badge struct {
	item     rune
	priority int
}

func parseLine(s string) (string, string) {
	l := len(s)
	if (l % 2) != 0 {
		logrus.Fatalf("length not even: %d", l)
	}
	l = l / 2
	return s[:l], s[l:]

}

func (r *rucksack) examinePockets() {
	p1 := mapset.NewSet([]rune(r.pocket1)...)
	p2 := mapset.NewSet([]rune(r.pocket2)...)
	intersection := p1.Intersect(p2)
	if intersection.Cardinality() > 1 {
		logrus.Fatalf("More than 1 intersecting item in %s : %s", r.pocket1, r.pocket2)
	}
	intersect, _ := intersection.Pop()
	r.mistakeItem = intersect
}

func determinePriority(i rune) int {
	if i < 97 {
		return int(i - 38)
	}
	return int(i - 96)
}

func findBadge(r1, r2, r3 rucksack) rune {
	s1 := mapset.NewSet(strings.Split(r1.contents, "")...)
	s2 := mapset.NewSet(strings.Split(r2.contents, "")...)
	s3 := mapset.NewSet(strings.Split(r3.contents, "")...)

	intersection := s1.Intersect(s2).Intersect(s3)
	if intersection.Cardinality() > 1 {
		logrus.Fatalf("More than 1 intersecting item in %s : %s : %s", r1.contents, r2.contents, r3.contents)
	}

	inter, _ := intersection.Pop()
	return []rune(inter)[0]
}

func parseInput03(lines []string) []rucksack {
	rucksacks := []rucksack{}
	for _, line := range lines {
		sack := rucksack{
			contents:        line,
			pocket1:         "",
			pocket2:         "",
			mistakeItem:     0,
			mistakePriority: 0,
		}
		sack.pocket1, sack.pocket2 = parseLine(line)
		sack.examinePockets()
		sack.mistakePriority = determinePriority(sack.mistakeItem)
		rucksacks = append(rucksacks, sack)
	}
	return rucksacks
}

func (p Day03) PartA(lines []string) any {
	rucksacks := parseInput03(lines)
	priority := 0
	for _, r := range rucksacks {
		priority += r.mistakePriority
	}
	return fmt.Sprintf("Total priority: %d", priority)
}

func (p Day03) PartB(lines []string) any {
	rucksacks := parseInput03(lines)
	badges := []badge{}
	for i := 0; i < len(rucksacks); i += 3 {
		// findBadge(rucksacks[i],rucksacks[i+1],rucksacks[i+2])
		thisBadge := badge{
			item:     findBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2]),
			priority: 0,
		}
		thisBadge.priority = determinePriority(thisBadge.item)
		badges = append(badges, thisBadge)
	}

	priority := 0
	for _, b := range badges {
		priority += b.priority
	}
	return fmt.Sprintf("Total badge priority: %d", priority)
}
