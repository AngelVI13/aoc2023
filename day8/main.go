package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func solution(directions []Direction, pathMap map[string][]string) int {
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

	answer := solution(directions, pathMap)
	fmt.Println("Answer:", answer)
	fmt.Println(time.Since(start))
}
