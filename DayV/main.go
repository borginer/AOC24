package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInts(strs []string) ([]int, error) {
	var nums []int
	for _, str := range strs {
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nums, err
		}
		nums = append(nums, int(n))
	}
	return nums, nil
}

func medianRightUpdate(upt string, order map[int][]int) int {
	numsStr := strings.Split(upt, ",")
	nums, err := parseInts(numsStr)
	if err != nil {
		log.Fatal("parsing f")
	}
	var prev []int
	for _, num := range nums {
		for _, p := range prev {
			if slices.Contains(order[p], num) {
				return 0
			}
		}
		prev = append(prev, num)
	}
	return nums[len(nums)/2]
}

func medianWrongUpdate(upt string, order map[int][]int) int {
	numsStr := strings.Split(upt, ",")
	nums, err := parseInts(numsStr)
	if err != nil {
		log.Fatal("parsing f")
	}

	alreadyOrdered := true
	isOrdered := false
	for !isOrdered {
		isOrdered = true
		prevs := []int{}
		for i, num := range nums {
			for _, p := range prevs {
				if slices.Contains(order[p], num) {
					alreadyOrdered = false
					isOrdered = false
					j := slices.Index(nums, p)
					nums[i], nums[j] = nums[j], nums[i]
				}
			}
			prevs = append(prevs, num)
		}
	}
	if !alreadyOrdered {
		return nums[len(nums)/2]
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	order := make(map[int][]int)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		numsStr := strings.Split(scanner.Text(), "|")
		nums, err := parseInts(numsStr)
		if err != nil {
			log.Fatal("parsing f")
		}
		order[int(nums[1])] = append(order[int(nums[1])], int(nums[0]))
	}
	rightOrderSum := 0
	wrongOrderSum := 0
	for scanner.Scan() {
		rightOrderSum += medianRightUpdate(scanner.Text(), order)
		wrongOrderSum += medianWrongUpdate(scanner.Text(), order)
	}

	fmt.Println("part 1:", rightOrderSum)
	fmt.Println("part 2:", wrongOrderSum)
}
