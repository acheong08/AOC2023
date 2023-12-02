package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var MAX_MAP = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		tmpSplit := strings.Split(line, ":")
		gameNum, err := strconv.Atoi(strings.Split(tmpSplit[0], " ")[1])
		if err != nil {
			panic(err)
		}
		sets := strings.Split(tmpSplit[1], ";")
		gamePassed := true
		for _, set := range sets {
			// Count the number of red/green/blue balls
			balls := strings.Split(set, ",")
			ballsMap := map[string]int{}
			for _, ball := range balls {
				ball = strings.TrimSpace(ball)
				tmpSplit := strings.Split(ball, " ")
				num, err := strconv.Atoi(tmpSplit[0])
				if err != nil {
					panic(err)
				}
				if _, ok := ballsMap[tmpSplit[1]]; ok {
					ballsMap[tmpSplit[1]] += num
				} else {
					ballsMap[tmpSplit[1]] = num
				}
			}
			for k, v := range ballsMap {
				if v > MAX_MAP[k] {
					gamePassed = false
					break
				}
			}
		}
		if gamePassed {
			total += gameNum
		}
	}
	println(total)
}
