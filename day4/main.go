package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		tmpArr := strings.Split(line, ":")
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
	}
	fmt.Println(totalScore)
}
