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

type Loc struct {
	X int
	Y int
}

func findGalaxies(galaxyMap [][]rune) []Loc {
	var galaxies []Loc
	for i, row := range galaxyMap {
		for j, col := range row {
			if col == '#' {
				galaxies = append(galaxies, Loc{X: j, Y: i})
			}
		}
	}
	return galaxies
}

func findExpansionIdxs(galaxyMap [][]rune) [2][]int {
	out := [2][]int{
		{},
		{},
	}

	for i, row := range galaxyMap {
		empty := true
		for _, col := range row {
			if col == '#' {
				empty = false
			}
		}

		if empty {
			out[0] = append(out[0], i)
		}
	}

	// horizontal expansion

	for i := 0; i < len(galaxyMap[0]); i++ {
		empty := true
		for j := 0; j < len(galaxyMap); j++ {
			if galaxyMap[j][i] == '#' {
				empty = false
			}
		}

		if empty {
			out[1] = append(out[1], i)
		}

	}

	return out
}

func printGalaxy(g [][]rune) {
	for _, row := range g {
		for _, col := range row {
			fmt.Printf("%c", col)
		}
		fmt.Println()
	}
}

func makePairKey(l1, l2 int) string {
	first := l1
	second := l2
	if l2 < l1 {
		first = l2
		second = l1
	}

	return fmt.Sprintf("%d,%d", first, second)
}

func findGalaxyPairs(galaxies []Loc) map[string][2]Loc {
	pairs := map[string][2]Loc{}
	for i := range galaxies {
		for j := range galaxies {
			if i == j {
				continue
			}

			key := makePairKey(i, j)

			if _, exists := pairs[key]; exists {
				continue
			}
			pairs[key] = [2]Loc{galaxies[i], galaxies[j]}
		}
	}
	return pairs
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func solutionPartOne(galaxyMap [][]rune, expandNum int) int {
	expIdxs := findExpansionIdxs(galaxyMap)
	galaxies := findGalaxies(galaxyMap)
	pairs := findGalaxyPairs(galaxies)

	if expandNum > 1 {
		expandNum--
	}

	sum := 0
	for _, pair := range pairs {
		first := pair[0]
		second := pair[1]

		// vertical expansions
		numExpansions := 0
		for _, idx := range expIdxs[0] {
			if (first.Y < idx && idx < second.Y) || (second.Y < idx && idx < first.Y) {
				numExpansions++
			}
		}

		// horizontal expansions
		for _, idx := range expIdxs[1] {
			if (first.X < idx && idx < second.X) || (second.X < idx && idx < first.X) {
				numExpansions++
			}
		}

		steps := abs(second.X - first.X)
		steps += abs(second.Y - first.Y)
		steps += numExpansions * expandNum
		// fmt.Println(first, second, steps)

		sum += steps
	}
	return sum
}

func solutionPartTwo() int {
	return 0
}

func main() {
	start := time.Now()
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	txt := string(b)

	lines := strings.Split(txt, "\n")
	var galaxyMap [][]rune

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var m []rune
		for _, char := range line {
			m = append(m, char)
		}
		galaxyMap = append(galaxyMap, m)
	}

	answer := solutionPartOne(galaxyMap, 1)
	fmt.Println("Answer PART1:", answer)
	answer = solutionPartOne(galaxyMap, 1000000)
	fmt.Println("Answer PART2:", answer)
	fmt.Println(time.Since(start))
}
