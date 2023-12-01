package main

import (
	"bufio"
	"os"
	"strconv"
)


func main() {
	total := 0
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		nums := make([]int, 0)
		for i := 0; i < len(txt); i++ {
			iTxt, err := strconv.Atoi(string(txt[i]))
			if err != nil {
				continue
			}
			nums = append(nums, iTxt)
		}
		if len(nums) == 0 {
			continue
		}
		total = total + (nums[0]*10) + nums[len(nums)-1]
	}
	if scanner.Err() != nil {
		panic("Failed to scan file")
	}
	println(total)
}
