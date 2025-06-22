package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	i, j int
}

func (n Node) add(other Node) Node {
	return Node{n.i + other.i, n.j + other.j}
}

func (n Node) sub(other Node) Node {
	return Node{n.i - other.i, n.j - other.j}
}

func diff(a, b Node) (int, int) {
	return a.i - b.i, a.j - b.j
}

func onBoard(a Node, w, h int) bool {
	return a.i >= 0 && a.j >= 0 && a.i < w && a.j < h
}

func antiNodesP1(locs []Node, w, h int) []Node {
	nodes := []Node{}
	for i, n1 := range locs {
		for _, n2 := range locs[i:] {
			if n1 == n2 {
				continue
			}
			di, dj := diff(n1, n2)
			p1, p2 := n1.add(Node{di, dj}), n2.add(Node{-di, -dj})
			if onBoard(p1, w, h) {
				nodes = append(nodes, p1)
			}
			if onBoard(p2, w, h) {
				nodes = append(nodes, p2)
			}
		}
	}
	return nodes
}

func antiNodesP2(locs []Node, w, h int) []Node {
	nodes := []Node{}
	for i, n1 := range locs {
		for _, n2 := range locs[i:] {
			if n1 == n2 {
				continue
			}
			di, dj := diff(n1, n2)
			diffNode := Node{di, dj}
			p1 := n1
			for onBoard(p1, w, h) {
				nodes = append(nodes, p1)
				p1 = p1.add(diffNode)
			}
			p2 := n2
			for onBoard(p2, w, h) {
				nodes = append(nodes, p2)
				p2 = p2.sub(diffNode)
			}
		}
	}
	return nodes
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	antenMap := make(map[rune][]Node)

	i, line := 0, ""
	for scanner.Scan() {
		line = scanner.Text()
		for j, r := range line {
			if r != '.' {
				antenMap[r] = append(antenMap[r], Node{i, j})
			}
		}
		i++
	}

	width := len(line)
	height := i
	nodeSetP1 := make(map[Node]bool)
	nodeSetP2 := make(map[Node]bool)

	for _, locs := range antenMap {
		for _, n := range antiNodesP1(locs, height, width) {
			nodeSetP1[n] = true
		}
		for _, n := range antiNodesP2(locs, height, width) {
			nodeSetP2[n] = true
		}
	}

	fmt.Println("part 1:", len(nodeSetP1))
	fmt.Println("part 2:", len(nodeSetP2))
}
