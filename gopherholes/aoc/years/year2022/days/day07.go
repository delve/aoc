package year2022

import (
	"aocgen/pkg/common"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Day07 struct{}

type node struct {
	name     string
	size     int
	children []*node
	parent   *node
}

func (n node) getChild(c string) *node {
	for _, child := range n.children {
		if child.name == c {
			return child
		}
	}
	return nil
}

// unused
// func (n node) getChildren() []string {
// 	var ret []string = nil

// 	if len(n.children) > 0 {
// 		for _, child := range n.children {
// 			ret = append(ret, child.name)
// 		}
// 	}
// 	return ret
// }

func (n *node) addChild(c string, s int) {
	childNode := node{name: c, size: s, parent: n}
	n.children = append(n.children, &childNode)
	// add the size all the way up
	for d := n; d != nil; d = d.parent {
		d.size += s
	}
}

func (n node) walk(cb func(node)) {
	cb(n)
	for _, c := range n.children {
		c.walk(cb)
	}
}

type cli struct {
	pwd *node
}

func (c *cli) cd(d string) error {
	switch d {
	case "..":
		if c.pwd.parent != nil {
			c.pwd = c.pwd.parent
		} else {
			return errors.New("attempt to cd .. past root")
		}
	case "/":
		// could instead walk up the tree from here and assign pwd once if i CBA to redo it
		for d := c.pwd.parent; d != nil; {
			c.cd("..")
		}
	default:
		child := c.pwd.getChild(d)
		if child != nil {
			c.pwd = child
		} else {
			return fmt.Errorf("no such dir in %s", c.pwd.name)
		}
	}

	return nil
}

func parseInput07(lines []string) node {
	fsTree := node{name: "/", parent: nil}
	user := cli{pwd: &fsTree}

	for _, line := range lines {
		tokens := strings.Split(line, " ")

		switch tokens[0] { // what have we here...
		case "$":
			if tokens[1] == "cd" {
				user.cd(tokens[2])
			} // we onlt care about cd, ls is effectively a noop
		case "dir": // the name of a directory, add it to the tree
			if user.pwd.getChild(tokens[1]) == nil {
				user.pwd.addChild(tokens[1], 0)
			}
		default: // probably a file entry, add it to the tree if we can parse the size
			size, err := strconv.Atoi(tokens[0])
			common.Check(err)
			if user.pwd.getChild(tokens[1]) == nil {
				user.pwd.addChild(tokens[1], size)
			}
		}
	}

	return fsTree
}

func (p Day07) PartA(lines []string) any {
	limit := 100000
	sumOfDirsOverLimit := 0

	fsTree := parseInput07(lines)
	fsTree.walk(func(n node) {
		if len(n.children) > 0 && n.size <= limit {
			sumOfDirsOverLimit += n.size
		}
	})

	return sumOfDirsOverLimit
}

func (p Day07) PartB(lines []string) any {
	fsTree := parseInput07(lines)

	// total 70000000, reqd 30000000
	spaceToFree := 30000000 - (70000000 - fsTree.size)
	sizeOfSmallestTargetDir := 0

	var candidateDirs []node

	fsTree.walk(func(n node) {
		if len(n.children) > 0 && n.size >= spaceToFree {
			candidateDirs = append(candidateDirs, n)
		}
	})

	targetSize := 70000000
	for _, n := range candidateDirs {
		if n.size >= spaceToFree && n.size < targetSize {
			targetSize = n.size
		}
	}
	sizeOfSmallestTargetDir = targetSize

	return sizeOfSmallestTargetDir
}
