package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type FreeSpace struct {
	start, size int
}

func findFreeSpaces(disk []rune) []FreeSpace {
	spaces := []FreeSpace{}
	for i := 0; i < len(disk); i++ {
		if disk[i] == '.' {
			chunkSize := 0
			for j := i; j < len(disk) && disk[j] == '.'; j++ {
				chunkSize++
			}
			spaces = append(spaces, FreeSpace{
				start: i,
				size:  chunkSize,
			})
			i += chunkSize - 1
		}
	}
	return spaces
}

func findFreeSpace(spaces []FreeSpace, start, size int) int {
	for i := range len(spaces) {
		if spaces[i].start > start {
			break
		}
		if spaces[i].size >= size {
			return i
		}
	}
	return len(spaces)
}

func copyn(runes []rune, from, to, size int) {
	for i := range size {
		runes[to+i] = runes[from+i]
		runes[from+i] = '.'
	}
}

func fillFreeSpace(spaces []FreeSpace, idx, size int) {
	if size < spaces[idx].size {
		spaces[idx].size -= size
		spaces[idx].start += size
	} else {
		spaces = slices.Delete(spaces, idx, idx+1)
	}
}

func calcChecksum(disk []rune) int {
	checksum := 0
	for i, r := range disk {
		if r != '.' {
			checksum += i * (int(r - '0'))
		}
	}
	return checksum
}

func nextFreeSpace(disk []rune, start int) int {
	for i := start; i < len(disk); i++ {
		if disk[i] == '.' {
			return i
		}
	}
	return len(disk)
}

func P1Remap(disk []rune) {
	nextEmpty := 0
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == '.' {
			continue
		}
		nextEmpty = nextFreeSpace(disk, nextEmpty)
		if nextEmpty < len(disk) {
			if nextEmpty > i {
				break
			}
			disk[nextEmpty] = disk[i]
			disk[i] = '.'
		}
	}
}

func P2Remap(disk []rune) {
	freeSpaces := findFreeSpaces(disk)
	movedFileIDs := make(map[rune]bool)
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == '.' {
			continue
		}
		id := disk[i]
		if movedFileIDs[id] {
			continue
		}
		chunkSize := 0
		for i >= 0 && disk[i] == id {
			i--
			chunkSize++
		}
		if chunkSize == 0 {
			continue
		}
		i++
		chunkStart := i
		fsIdx := findFreeSpace(freeSpaces, chunkStart, chunkSize)
		if fsIdx < len(freeSpaces) {
			copyn(disk, chunkStart, freeSpaces[fsIdx].start, chunkSize)
			fillFreeSpace(freeSpaces, fsIdx, chunkSize)
			movedFileIDs[id] = true
		}
	}
}

func buildDisk(input string) []rune {
	disk := []rune{}
	for i, r := range input {
		dig := r - '0'
		for range dig {
			if i%2 == 0 {
				disk = append(disk, '0'+rune(i/2))
			} else {
				disk = append(disk, '.')
			}

		}
	}
	return disk
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := ""
	if scanner.Scan() {
		input = scanner.Text()
	}
	disk := buildDisk(input)

	diskP1 := make([]rune, len(disk))
	copy(diskP1, disk)
	P1Remap(diskP1)
	fmt.Println("part 1:", calcChecksum(diskP1))

	diskP2 := make([]rune, len(disk))
	copy(diskP2, disk)
	P2Remap(diskP2)
	fmt.Println("part 2:", calcChecksum(diskP2))
}
