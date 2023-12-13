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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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
		fmt.Print(node, " ")
		directionIdx := 0
		steps := 0
		for {
			directionIdx = directionIdx % len(directions)
			node = pathMap[node][directions[directionIdx]]
			steps++
			directionIdx++
			if node[2] == 'Z' {
				fmt.Print(steps)
				stepsTillZ = append(stepsTillZ, steps)
				break
			}
		}
		fmt.Println()
		// fmt.Println(node, steps)
	}

	sort.Slice(stepsTillZ, func(i, j int) bool {
		return stepsTillZ[i] > stepsTillZ[j]
	})

	fmt.Println(stepsTillZ)
	return LCM(stepsTillZ[0], stepsTillZ[1], stepsTillZ[1:]...)
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
	fmt.Println(dirStr)
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

	// answer := solutionPartOne(directions, pathMap)
	answer := solutionPartTwo(directions, pathMap)
	fmt.Println("Answer:", answer)
	fmt.Println(time.Since(start))
}
