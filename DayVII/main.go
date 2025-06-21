package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_OP_LEN = 15
)

var allPossibleOpsP1 [][]rune
var allPossibleOpsP2 [][]rune

func init() {
	lines := int(math.Pow(2, MAX_OP_LEN))
	for line := range lines {
		allPossibleOpsP1 = append(allPossibleOpsP1, make([]rune, MAX_OP_LEN))
		for col := range MAX_OP_LEN {
			colExp := int(math.Pow(2, float64(col)))
			if (line/colExp)%2 == 0 {
				allPossibleOpsP1[line][col] = '*'
			} else {
				allPossibleOpsP1[line][col] = '+'
			}
		}
	}
	lines = int(math.Pow(3, MAX_OP_LEN))
	for line := range lines {
		allPossibleOpsP2 = append(allPossibleOpsP2, make([]rune, MAX_OP_LEN))
		for col := range MAX_OP_LEN {
			colExp := int(math.Pow(3, float64(col)))
			part := (line / colExp) % 3
			if part == 0 {
				allPossibleOpsP2[line][col] = '*'
			} else if part == 1 {
				allPossibleOpsP2[line][col] = '+'
			} else {
				allPossibleOpsP2[line][col] = '|'
			}

		}
	}
}

func concatNums(left, right int64) int64 {
	l := strconv.FormatInt(left, 10)
	r := strconv.FormatInt(right, 10)
	ret, err := strconv.ParseInt(l+r, 10, 64)
	if err != nil {
		panic("skill issue")
	}
	return ret
}

func calcEquation(nums []int64, ops []rune) int64 {
	res := int64(nums[0])
	for i, op := range ops {
		if op == '+' {
			res += nums[i+1]
		} else if op == '*' {
			res *= nums[i+1]
		} else if op == '|' {
			res = concatNums(res, nums[i+1])
		}
	}
	return res
}

func canBeProducedP1(res int64, nums []int64) bool {
	opsLen := len(nums) - 1
	opsLenExp := int(math.Pow(2, float64(opsLen)))
	for i := range opsLenExp {
		if calcEquation(nums, allPossibleOpsP1[i][:opsLen]) == res {
			return true
		}
	}
	return false
}

func canBeProducedP2(res int64, nums []int64) bool {
	opsLen := len(nums) - 1
	opsLenExp := int(math.Pow(3, float64(opsLen)))
	for i := range opsLenExp {
		if calcEquation(nums, allPossibleOpsP2[i][:opsLen]) == res {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	calibrationSumP1 := int64(0)
	calibrationSumP2 := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Fatalf("expected one ':', line: %s", line)
		}

		res, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatalf("error parsing %s", parts[0])
		}

		nums := []int64{}
		for _, field := range strings.Fields(parts[1]) {
			num, err := strconv.ParseInt(field, 10, 64)
			if err != nil {
				log.Fatalf("error parsing number: %s", field)
			}
			nums = append(nums, num)
		}

		if canBeProducedP1(res, nums) {
			calibrationSumP1 += res
		}
		if canBeProducedP2(res, nums) {
			calibrationSumP2 += res
		}
	}

	fmt.Println("part 1:", calibrationSumP1)
	fmt.Println("part 2:", calibrationSumP2)
}
