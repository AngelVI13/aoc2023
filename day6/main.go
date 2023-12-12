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

type Race struct {
	Time     int
	Distance int
}

func solutionPartOne(races []Race) int {
	result := 1
	for _, race := range races {
		ways := 0

		for speed := 1; speed < race.Time; speed++ {
			distance := speed * (race.Time - speed)

			if distance > race.Distance {
				ways++
			}
		}

		result *= ways
	}
	return result
}

func solutionPartTwo() {
}

func parseLine(s, head string) []int {
	s = strings.TrimSpace(strings.ReplaceAll(s, head, ""))
	var out []int

	for _, numStr := range strings.Split(s, " ") {
		numStr = strings.TrimSpace(numStr)
		if numStr == "" {
			continue
		}

		num, err := strconv.Atoi(numStr)
		checkErr(err)
		out = append(out, num)
	}

	return out
}

func main() {
	start := time.Now()
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	txt := string(b)

	lines := strings.Split(txt, "\n")
	times := parseLine(lines[0], "Time: ")
	distances := parseLine(lines[1], "Distance: ")

	var races []Race

	for i := range times {
		races = append(races, Race{
			Time:     times[i],
			Distance: distances[i],
		})
	}

	answer := solutionPartOne(races)
	fmt.Println("Answer:", answer)
	fmt.Println(time.Since(start))
}
