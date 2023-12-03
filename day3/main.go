package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := make([]string, 3)
	sum := 0
	for scanner.Scan() {
		lines[2] = scanner.Text()
		// Do something with line
		// fmt.Printf("%s\n> %s\n%s\n\n", lines[0], lines[1], lines[2])
		sum += hasAdjecentSymbol(lines)

		lines[0], lines[1] = lines[1], lines[2]
	}
	// For the last line
	lines[2] = ""
	sum += hasAdjecentSymbol(lines)
	// fmt.Printf("%s\n> %s\n%s\n\n", lines[0], lines[1], lines[2])

	fmt.Println("\n", sum)
}

type NumLoc struct {
	Start  int
	End    int
	Value  int
	tmpVal string
}

func (n *NumLoc) AddTemp(val int) {
	n.tmpVal += strconv.Itoa(val)
}
func (n *NumLoc) HasTemp() bool {
	return n.tmpVal != ""
}

func (n *NumLoc) SetVal() {
	n.Value, _ = strconv.Atoi(n.tmpVal)
}

func hasAdjecentSymbol(lines []string) (sum int) {
	if lines[1] == "" {
		return 0
	}

	coords := make([]NumLoc, 0)

	temp := NumLoc{}

	for i := 0; i < len(lines[1]); i++ {
		// Get all the numbers in the line with coords
		if num, err := strconv.Atoi(string(lines[1][i])); err == nil {
			if !temp.HasTemp() {
				temp.Start = i
			}
			if i == len(lines[1])-1 {
				temp.End = i
				temp.AddTemp(num)
				temp.SetVal()
				coords = append(coords, temp)
				temp = NumLoc{}
				continue
			}
			temp.AddTemp(num)
		} else if temp.HasTemp() {
			temp.End = i - 1
			temp.SetVal()
			coords = append(coords, temp)
			temp = NumLoc{}
		}
	}
	normalChars := "1234567890."
	adjecents := make([]int, 0)
	// fmt.Println(coords)
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if !strings.Contains(normalChars, string(line[i])) {
				// Found special character
				// fmt.Printf("Found special character at %d\n", i)
				// Check if it's adjacent to a number
				for _, coord := range coords {
					if math.Abs(float64(coord.Start-i)) <= 1 || math.Abs(float64(coord.End-i)) <= 1 || (i >= coord.Start && i <= coord.End){
						// print(coord.Value)
						// print(", ")
						adjecents = append(adjecents, coord.Value)
					} else {
						// fmt.Println(coord)
					}
				}
			}
		}
	}
	for _, adj := range adjecents {
		sum += adj
	}
	return sum
}
