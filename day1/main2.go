package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var strToNum = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	total := 0
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := strings.ToLower(scanner.Text())
		nums := make([]int, 0)
		currString := ""
		for i := 0; i < len(txt); i++ {
			iTxt, err := strconv.Atoi(string(txt[i]))
			if err != nil {
				currString += string(txt[i])
				shouldBreak := false
				for k, v := range strToNum {
					if strings.HasSuffix(currString, k) {
						iTxt = v
						shouldBreak = true
						break
					}
				}
				if !shouldBreak {
					continue
				}
			}
			nums = append(nums, iTxt)
			if len(currString) != 0 {
				currString = string(currString[len(currString)-1])
			}
		}
		if len(nums) == 0 {
			continue
		}
		total = total + (nums[0] * 10) + nums[len(nums)-1]
	}
	if scanner.Err() != nil {
		panic("Failed to scan file")
	}
	println(total)
}
