package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func solutionPartOne(directions []Direction, pathMap map[string][]string) int {
	steps := 0

	node := "AAA"
	directionIdx := 0
	for {
		node = pathMap[node][directions[directionIdx%len(directions)]]
		steps++
		directionIdx++
		if node == "ZZZ" {
			break
		}
	}
	return steps
}

var directionNames = map[Direction]string{
	Left:  "L",
	Right: "R",
}

func solutionPartTwo(directions []Direction, pathMap map[string][]string) int {
	var startingNodes []string
	for key := range pathMap {
		if key[2] == 'A' {
			startingNodes = append(startingNodes, key)
		}
	}

	var stepsTillZ []int

	for _, node := range startingNodes {
		directionIdx := 0
		steps := 0
		for {
			directionIdx = directionIdx % len(directions)
			node = pathMap[node][directions[directionIdx]]
			steps++
			directionIdx++
			if node[2] == 'Z' {
				stepsTillZ = append(stepsTillZ, steps)
				break
			}
		}
	}

	sort.Slice(stepsTillZ, func(i, j int) bool {
		return stepsTillZ[i] > stepsTillZ[j]
	})

	// lcm - lowest common multiplier
	lcm := stepsTillZ[0]
	for {
		allAreMultiple := true
		for _, num := range stepsTillZ {
			if lcm%num != 0 {
				allAreMultiple = false
				break
			}
		}

		if allAreMultiple {
			break
		}

		lcm += stepsTillZ[0]
	}
	return lcm
}

type Direction int

const (
	Left Direction = iota
	Right
)

func main() {
	start := time.Now()
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	txt := string(b)

	lines := strings.Split(txt, "\n")
	dirStr := lines[0]
	var directions []Direction

	for _, char := range strings.TrimSpace(dirStr) {
		if char == 'L' {
			directions = append(directions, Left)
		} else if char == 'R' {
			directions = append(directions, Right)
		} else {
			panic("failed to process directions")
		}
	}

	pathMap := map[string][]string{}
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		lineContents := strings.Split(line, " = ")
		node := strings.TrimSpace(lineContents[0])
		childrenStr := strings.ReplaceAll(lineContents[1], "(", "")
		childrenStr = strings.ReplaceAll(childrenStr, ")", "")

		pathMap[node] = strings.Split(childrenStr, ", ")
	}

	answer := solutionPartOne(directions, pathMap)
	fmt.Println("Answer PART1:", answer)
	answer = solutionPartTwo(directions, pathMap)
	fmt.Println("Answer PART2:", answer)
	fmt.Println(time.Since(start))
}
