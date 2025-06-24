package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func countStones(stones map[string]int) int {
	count := 0
	for _, amt := range stones {
		count += amt
	}
	return count
}

func trimZeros(s string) string {
	var i int
	for i = range len(s) {
		if s[i] != '0' {
			break
		}
	}
	runes := []rune{}
	for j := i; j < len(s); j++ {
		runes = append(runes, rune(s[j]))
	}
	return string(runes)
}

func blink(stones map[string]int) map[string]int {
	blinkMap := make(map[string]int, len(stones))

	for s, amt := range stones {
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatal("error parsing:", s)
		}

		if val == 0 {
			blinkMap["1"] += amt
		} else if len(s)%2 == 0 {
			left, right := s[:len(s)/2], s[len(s)/2:]
			right = trimZeros(right)
			blinkMap[left] += amt
			blinkMap[right] += amt
		} else {
			val = 2024 * val
			newStone := strconv.Itoa(int(val))
			blinkMap[newStone] += amt
		}
	}
	return blinkMap
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("file read f")
	}

	stones := make(map[string]int)
	stonesSlice := strings.Fields(input)
	for _, stone := range stonesSlice {
		stones[stone]++
	}

	for range 25 {
		stones = blink(stones)
	}
	fmt.Println("part 1:", countStones(stones))

	for range 50 {
		stones = blink(stones)
	}
	fmt.Println("part 2:", countStones(stones))
}
