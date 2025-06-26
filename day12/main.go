package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Garden [][]rune
type Tile struct{ i, j int }
type Dir struct{ i, j int }
type FenceNode struct {
	tile Tile
	dir  Dir
}

var DIRS []Dir = []Dir{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var VERT []Dir = []Dir{{1, 0}, {-1, 0}}
var HORIZ []Dir = []Dir{{0, 1}, {0, -1}}

func (g Garden) inGarden(t Tile) bool {
	return t.i >= 0 && t.j >= 0 && t.i < len(g) && t.j < len(g[0])
}

func (g Garden) get(t Tile) rune {
	return g[t.i][t.j]
}

func (t Tile) add(o Dir) Tile {
	return Tile{t.i + o.i, t.j + o.j}
}

func (fn FenceNode) move(dir Dir) FenceNode {
	return FenceNode{fn.tile.add(dir), fn.dir}
}

func calcAreaCostP1(area map[Tile]bool) int {
	areaSize, fenceLen := 0, 0
	for tile := range area {
		areaSize++
		for _, dir := range DIRS {
			if !area[tile.add(dir)] {
				fenceLen++
			}
		}
	}
	return areaSize * fenceLen
}

func removeLine(from FenceNode, fence map[FenceNode]bool) bool {
	if !fence[from] {
		return false
	}

	var dirs []Dir
	if slices.Contains(HORIZ, from.dir) {
		dirs = VERT
	} else {
		dirs = HORIZ
	}

	for _, dir := range dirs {
		for t := from.move(dir); fence[t]; t = t.move(dir) {
			fence[t] = false
		}
	}

	fence[from] = false
	return true
}

func calcAreaCostP2(area map[Tile]bool, width, height int) int {
	areaSize, fenceLen := 0, 0
	fence := make(map[FenceNode]bool)
	for tile := range area {
		areaSize++
		for _, dir := range DIRS {
			nextTile := tile.add(dir)
			if !area[nextTile] {
				fence[FenceNode{tile, dir}] = true
			}
		}
	}

	for i := -1; i < height+1; i++ {
		for j := -1; j < width+1; j++ {
			for _, dir := range DIRS {
				fn := FenceNode{Tile{i, j}, dir}
				if fence[fn] {
					removeLine(fn, fence)
					fenceLen++
				}
			}
		}
	}

	return areaSize * fenceLen
}

func getArea(garden Garden, cur Tile, visited map[Tile]bool, plantType rune, curTiles map[Tile]bool) {
	if !garden.inGarden(cur) || garden.get(cur) != plantType || visited[cur] {
		return
	}
	visited[cur] = true
	for _, dir := range DIRS {
		nextTile := cur.add(dir)
		curTiles[cur] = true
		getArea(garden, nextTile, visited, plantType, curTiles)
	}
}

func totalFencesPrice(garden Garden) (p1 int, p2 int) {
	costP1, costP2 := 0, 0
	visited := make(map[Tile]bool)
	for i, line := range garden {
		for j, r := range line {
			curTile := Tile{i, j}
			if visited[curTile] {
				continue
			}
			curArea := make(map[Tile]bool)
			getArea(garden, curTile, visited, r, curArea)
			costP1 += calcAreaCostP1(curArea)
			costP2 += calcAreaCostP2(curArea, len(garden[0]), len(garden))
		}
	}
	return costP1, costP2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	garden := Garden{}
	for scanner.Scan() {
		garden = append(garden, []rune(scanner.Text()))
	}

	priceP1, priceP2 := totalFencesPrice(garden)
	fmt.Println("part 1:", priceP1)
	fmt.Println("part 2:", priceP2)
}
