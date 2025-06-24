package main

import (
	"bufio"
	"fmt"
	"os"
)

type TopoMap []string
type Node struct{ i, j int }
type Dir struct{ i, j int }

var DIRS []Dir = []Dir{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func (tm TopoMap) onMap(n Node) bool {
	return n.i >= 0 && n.j >= 0 && n.i < len(tm) && n.j < len(tm[0])
}

func (n Node) add(d Dir) Node {
	return Node{n.i + d.i, n.j + d.j}
}

func countTopsFrom(tMap TopoMap, prev byte, n Node) map[Node]bool {
	if !(tMap).onMap(n) || tMap[n.i][n.j] != prev+1 {
		return nil
	}
	if tMap[n.i][n.j] == '9' {
		ret := make(map[Node]bool)
		ret[n] = true
		return ret
	}
	trails := make(map[Node]bool)
	for _, dir := range DIRS {
		nodes := countTopsFrom(tMap, tMap[n.i][n.j], n.add(dir))
		for n := range nodes {
			trails[n] = true
		}
	}
	return trails
}

func countTrailsFrom(tMap TopoMap, prev byte, n Node) []Node {
	if !(tMap).onMap(n) || tMap[n.i][n.j] != prev+1 {
		return nil
	}
	if tMap[n.i][n.j] == '9' {
		return []Node{n}
	}
	trails := []Node{}
	for _, dir := range DIRS {
		nodes := countTrailsFrom(tMap, tMap[n.i][n.j], n.add(dir))
		trails = append(trails, nodes...)
	}
	return trails
}

func countTrails(tMap TopoMap) int {
	count := 0
	for i, line := range tMap {
		for j, r := range line {
			if r == '0' {
				count += len(countTrailsFrom(tMap, '0'-1, Node{i, j}))
			}
		}
	}
	return count
}

func countTops(tMap TopoMap) int {
	count := 0
	for i, line := range tMap {
		for j, r := range line {
			if r == '0' {
				count += len(countTopsFrom(tMap, '0'-1, Node{i, j}))
			}
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	topoMap := []string{}
	for scanner.Scan() {
		topoMap = append(topoMap, scanner.Text())
	}

	fmt.Println("part 1:", countTops(topoMap))
	fmt.Println("part 2:", countTrails(topoMap))
}
