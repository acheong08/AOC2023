package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	totalScore := 0
	data, _ := io.ReadAll(f)
	lines := strings.Split(string(data), "\n")
	cards := make([]int, len(lines)-1)
	for i := 0; i < len(cards); i++ {
		cards[i] = 1
	}
	for index, line := range lines {
		tmpArr := strings.Split(line, ":")
		if len(tmpArr) < 2 {
			break
		}
		tmpArr = strings.Split(strings.TrimSpace(tmpArr[1]), "|")
		numbers := strings.Split(strings.TrimSpace(tmpArr[0]), " ")
		winners := strings.Split(strings.TrimSpace(tmpArr[1]), " ")
		winNum := 0
		for _, number := range numbers {
			if number == "" {
				continue
			}
			for _, winner := range winners {
				if number == winner {
					winNum++
					break
				}
			}
		}
		if winNum > 0 {
			totalScore += int(math.Pow(2, float64(winNum)-1))
		}
		for i := 0; i < winNum; i++ {
			cards[index+i+1] += cards[index]
		}
	}
	fmt.Println(totalScore)
	fmt.Println(sum(cards))
}

func sum(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total
}
