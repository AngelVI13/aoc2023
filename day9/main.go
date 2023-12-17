package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func solutionPartOne(measurements [][]int) int {
	sum := 0

	for _, sensor_measurements := range measurements {
		var analysis [][]int
		analysis = append(analysis, sensor_measurements)
		for {
			var diff []int

			allZeros := true
			compare_values := analysis[len(analysis)-1]
			for i, val := range compare_values {
				if i == 0 {
					continue
				}
				valDiff := val - compare_values[i-1]
				if valDiff != 0 {
					allZeros = false
				}
				diff = append(diff, valDiff)
			}

			analysis = append(analysis, diff)

			if allZeros {
				break
			}
		}
		// fmt.Println(analysis)

		for i := len(analysis) - 1; i >= 0; i-- {
			if i == len(analysis)-1 {
				analysis[i] = append(analysis[i], analysis[i][len(analysis[i])-1])
				continue
			}

			lastElementFromNextLevel := analysis[i+1][len(analysis[i+1])-1]
			lastElementFromCurrentLevel := analysis[i][len(analysis[i])-1]
			newValue := lastElementFromCurrentLevel + lastElementFromNextLevel
			analysis[i] = append(analysis[i], newValue)

			if i == 0 {
				sum += newValue
			}
		}
		// fmt.Println(analysis)
		// fmt.Println()
	}
	return sum
}

func solutionPartTwo(measurements [][]int) int {
	sum := 0

	for _, sensor_measurements := range measurements {
		var analysis [][]int
		analysis = append(analysis, sensor_measurements)
		for {
			var diff []int

			allZeros := true
			compare_values := analysis[len(analysis)-1]
			for i, val := range compare_values {
				if i == 0 {
					continue
				}
				valDiff := val - compare_values[i-1]
				if valDiff != 0 {
					allZeros = false
				}
				diff = append(diff, valDiff)
			}

			analysis = append(analysis, diff)

			if allZeros {
				break
			}
		}
		// fmt.Println(analysis)

		for i := len(analysis) - 1; i >= 0; i-- {
			if i == len(analysis)-1 {
				analysis[i] = append(analysis[i], analysis[i][len(analysis[i])-1])
				continue
			}

			firstElementFromNextLevel := analysis[i+1][0]
			firstElementFromCurrentLevel := analysis[i][0]
			newValue := firstElementFromCurrentLevel - firstElementFromNextLevel
			analysis[i] = append([]int{newValue}, analysis[i]...)

			if i == 0 {
				sum += newValue
			}
		}
		// fmt.Println(analysis)
		// fmt.Println()
	}
	return sum
}

func main() {
	start := time.Now()
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	txt := string(b)

	lines := strings.Split(txt, "\n")
	var measurements [][]int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var m []int
		lineElems := strings.Split(line, " ")
		for _, elem := range lineElems {
			elemInt, err := strconv.Atoi(elem)
			checkErr(err)

			m = append(m, elemInt)
		}
		measurements = append(measurements, m)
	}

	answer := solutionPartOne(measurements)
	fmt.Println("Answer PART1:", answer)
	answer = solutionPartTwo(measurements)
	fmt.Println("Answer PART2:", answer)
	fmt.Println(time.Since(start))
}
