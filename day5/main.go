package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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

func solutionPartOne(seeds []int, rawMaps [][]Convertion) (int, int) {
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
	return minSeed, minLocation
}

func solutionPartTwo(seeds []int, rawMaps [][]Convertion) (int, int) {
	minLocation := -1
	minSeed := -1
	for i := 0; i < len(seeds)-1; i += 2 {
		startSeedNr := seeds[i]
		seedRangeLen := seeds[i+1]
		fmt.Println("calculating location for seedNr", startSeedNr, seedRangeLen)

		for seedNr := startSeedNr; seedNr < startSeedNr+seedRangeLen; seedNr++ {
			value := seedNr
			for _, rules := range rawMaps {
				mmap := MMap(rules)
				value = mmap.Apply(value)
			}

			if minLocation == -1 || value < minLocation {
				minLocation = value
				minSeed = seedNr
			}
		}
	}
	return minSeed, minLocation
}

func solutionPartTwoParallel(seeds []int, rawMaps [][]Convertion) (int, int) {
	wg := sync.WaitGroup{}
	out := make(chan [2]int, len(seeds)/2)
	for i := 0; i < len(seeds)-1; i += 2 {
		startSeedNr := seeds[i]
		seedRangeLen := seeds[i+1]
		wg.Add(1)
		go func(idx, start, len_ int, out chan<- [2]int) {
			fmt.Println(idx, "calculating location for seedNr", start, len_)
			minLocation := -1
			minSeed := -1
			for seedNr := start; seedNr < len_; seedNr++ {
				value := seedNr
				for _, rules := range rawMaps {
					mmap := MMap(rules)
					value = mmap.Apply(value)
				}

				if minLocation == -1 || value < minLocation {
					minLocation = value
					minSeed = seedNr
				}
			}
			fmt.Println(idx, "found min location", minSeed, minLocation)
			wg.Done()
			out <- [2]int{minSeed, minLocation}
		}(i, startSeedNr, startSeedNr+seedRangeLen, out)
	}

	wg.Wait()
	close(out)

	minLocation := -1
	minSeed := -1
	for i := 0; i < len(seeds)/2; i++ {
		values := <-out
		if minLocation == -1 || values[1] < minLocation {
			minLocation = values[1]
			minSeed = values[0]
		}
	}

	return minSeed, minLocation
}

func main() {
	start := time.Now()
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	txt := string(b)

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
			continue
		}

		if strings.Contains(line, " map:") {
			rawMaps = append(rawMaps, []Convertion{})
		} else {
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

	minSeed, minLocation := solutionPartTwoParallel(seeds, rawMaps)
	fmt.Println("MIN:", minSeed, "->", minLocation)
	fmt.Println(time.Since(start))
}
