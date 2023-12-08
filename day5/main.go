package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Convertion struct {
	SrcStart int
	DstStart int
	Len      int
}

type MMap []Convertion

func (m MMap) Apply(v int) int {
	out := v
	for _, rule := range m {
		if v >= rule.SrcStart && v < rule.SrcStart+rule.Len {
			out = rule.DstStart + (v - rule.SrcStart)
		}
	}
	return out
}

func main() {
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	txt := string(b)
	sum := 0

	lines := strings.Split(txt, "\n")
	seedsStrSlice := strings.Split(strings.ReplaceAll(lines[0], "seeds: ", ""), " ")
	var seeds []int
	for _, s := range seedsStrSlice {
		seed, err := strconv.Atoi(s)
		checkErr(err)
		seeds = append(seeds, seed)
	}
	fmt.Println(seeds)

	rawMaps := [][]Convertion{}
	for _, line := range lines[1 : len(lines)-1] {
		if strings.TrimSpace(line) == "" {
			fmt.Println("empty line", line)
			continue
		}

		if strings.Contains(line, " map:") {
			fmt.Println("new map found line", line, rawMaps)
			rawMaps = append(rawMaps, []Convertion{})
		} else {
			fmt.Println("add line to map", line, rawMaps)
			idx := len(rawMaps) - 1
			lineSlice := strings.Split(strings.TrimSpace(line), " ")
			dstStart, err := strconv.Atoi(lineSlice[0])
			checkErr(err)
			srcStart, err := strconv.Atoi(lineSlice[1])
			checkErr(err)
			len_, err := strconv.Atoi(lineSlice[2])
			checkErr(err)

			rawMaps[idx] = append(rawMaps[idx], Convertion{DstStart: dstStart, SrcStart: srcStart, Len: len_})
		}
	}

	fmt.Println(sum)

	minLocation := -1
	minSeed := -1
	for _, seedNr := range seeds {
		value := seedNr
		for _, rules := range rawMaps {
			mmap := MMap(rules)
			value = mmap.Apply(value)
		}

		if minLocation == -1 || value < minLocation {
			minLocation = value
			minSeed = seedNr
		}
		fmt.Println(seedNr, "->", value)
	}

	fmt.Println("MIN:", minSeed, "->", minLocation)
}
