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

type Dir struct {
	X int
	Y int
}

func makeDirKey(d Dir) string {
	return fmt.Sprintf("%d,%d", d.X, d.Y)
}

type PipeDirs struct {
	Prev Dir
	Next Dir
}

var tileDirs = map[rune][]PipeDirs{
	'|': {
		PipeDirs{
			Prev: Dir{X: 0, Y: 1},
			Next: Dir{X: 0, Y: 1},
		},
		PipeDirs{
			Prev: Dir{X: 0, Y: -1},
			Next: Dir{X: 0, Y: -1},
		},
	},
	'-': {
		PipeDirs{
			Prev: Dir{X: 1, Y: 0},
			Next: Dir{X: 1, Y: 0},
		},
		PipeDirs{
			Prev: Dir{X: -1, Y: 0},
			Next: Dir{X: -1, Y: 0},
		},
	},
	'L': {
		PipeDirs{
			Prev: Dir{X: -1, Y: 0},
			Next: Dir{X: 0, Y: -1},
		},
		PipeDirs{
			Prev: Dir{X: 0, Y: 1},
			Next: Dir{X: 1, Y: 0},
		},
	},
	'J': {
		PipeDirs{
			Prev: Dir{X: 1, Y: 0},
			Next: Dir{X: 0, Y: -1},
		},
		PipeDirs{
			Prev: Dir{X: 0, Y: 1},
			Next: Dir{X: -1, Y: 0},
		},
	},
	'7': {
		PipeDirs{
			Prev: Dir{X: 1, Y: 0},
			Next: Dir{X: 0, Y: 1},
		},
		PipeDirs{
			Prev: Dir{X: 0, Y: -1},
			Next: Dir{X: -1, Y: 0},
		},
	},
	'F': {
		PipeDirs{
			Prev: Dir{X: 0, Y: -1},
			Next: Dir{X: 1, Y: 0},
		},
		PipeDirs{
			Prev: Dir{X: -1, Y: 0},
			Next: Dir{X: 0, Y: 1},
		},
	},
	'.': {},
}

var allowedPipesByDir = map[string][]rune{
	"0,1":  {'L', 'J', '|'},
	"0,-1": {'|', '7', 'F'},
	"1,0":  {'-', 'J', '7'},
	"-1,0": {'-', 'L', 'F'},
}

var AllDirs = []Dir{
	{X: 1, Y: 0},
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}

var AllSearchDirs = []Dir{
	{X: 1, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 1},
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: 0, Y: -1},
	{X: 1, Y: -1},
}

func findPath(pipesMap [][]rune) []Dir {
	startDir := Dir{}
	for i := 0; i < len(pipesMap); i++ {
		for j := 0; j < len(pipesMap[i]); j++ {
			r := pipesMap[i][j]
			if r == 'S' {
				startDir.X = j
				startDir.Y = i
			}
		}
	}
	mapLenX := len(pipesMap[0])
	mapLenY := len(pipesMap)

	firstPipe := Dir{}
	for _, d := range AllDirs {
		newY := startDir.Y + d.Y
		newX := startDir.X + d.X
		if newY < 0 || newY >= mapLenY || newX < 0 || newX >= mapLenX {
			continue
		}

		dirK := makeDirKey(d)
		nextR := pipesMap[startDir.Y+d.Y][startDir.X+d.X]

		allowed := false
		for _, r := range allowedPipesByDir[dirK] {
			if nextR == r {
				allowed = true
				break
			}
		}

		if allowed {
			firstPipe = d
			break
		}
	}
	/*
		fmt.Printf(
			"%c: %s\n",
			pipesMap[startDir.Y+firstPipe.Y][startDir.X+firstPipe.X],
			makeDirKey(firstPipe),
		)
	*/
	path := []Dir{startDir}

	nextDir := firstPipe
	currentPos := startDir
	// currentRune := pipesMap[currentPos.Y][currentPos.X]
	count := 0
	for {
		count++
		// fmt.Println(fmt.Sprintf("%c", currentRune), currentPos, makeDirKey(nextDir))
		newPos := Dir{
			Y: currentPos.Y + nextDir.Y,
			X: currentPos.X + nextDir.X,
		}
		newR := pipesMap[newPos.Y][newPos.X]

		if newR == 'S' {
			break
		}
		// fmt.Println(fmt.Sprintf("%c", newR), newPos)

		possibleDirs := tileDirs[newR]
		for _, d := range possibleDirs {
			if makeDirKey(nextDir) == makeDirKey(d.Prev) {
				nextDir = d.Next
				break
			}
		}
		path = append(path, newPos)
		currentPos = newPos
		// currentRune = newR
	}

	return path
}

// | - L J 7 F . S

func solutionPartOne(pipesMap [][]rune) int {
	path := findPath(pipesMap)
	return len(path) / 2
}

func findStartSymbol(path []Dir) rune {
	first := path[1]
	last := path[len(path)-1]

	if first.X == last.X {
		return '|'
	} else if first.Y == last.Y {
		return '-'
	} else if first.X < last.X && first.Y < last.Y {
		return '7'
	} else if first.X < last.X && first.Y > last.Y {
		return 'J'
	} else if first.X > last.X && first.Y < last.Y {
		return 'F'
	} else if first.X > last.X && first.Y > last.Y {
		return 'L'
	} else {
		panic("something went terribly wrong")
	}
}

/*
HINT from this reddit comment

		Part 2 using one of my favorite facts from graphics engineering:
	    lets say you have an enclosed shape, and you want to color every pixel
	    inside of it. How do you know if a given pixel is inside the shape or not?
	    Well, it turns out: if you shoot a ray in any direction from the pixel and
	    it crosses the boundary an odd number of times, it's inside. if it crosses
	    an even number of times, it's outside. Works for all enclosed shapes, even
	    self-intersecting and non-convex ones.

	    It does, however, interact badly if your ray and one of the edges of the
	    shape is collinear, so you have to be clever about it for this problem.
*/
func solutionPartTwo(pipesMap [][]rune) int {
	path := findPath(pipesMap)
	pathMap := map[string]Dir{}
	for _, part := range path {
		pathMap[makeDirKey(part)] = part
	}

	startSymbol := findStartSymbol(path)
	// Replace 'S' with the appropriate symbol
	pipesMap[path[0].Y][path[0].X] = startSymbol

	tilesInside := 0
	for i := 0; i < len(pipesMap); i++ {
		for j := 0; j < len(pipesMap[0]); j++ {
			pieceLoc := Dir{X: j, Y: i}
			_, isPipePath := pathMap[makeDirKey(pieceLoc)]

			if isPipePath {
				continue
			}

			searchDir := Dir{X: 1, Y: -1}
			currentPlace := pieceLoc
			crosses := 0
			for {
				newPlace := Dir{
					X: currentPlace.X + searchDir.X,
					Y: currentPlace.Y + searchDir.Y,
				}
				if newPlace.X < 0 || newPlace.X >= len(pipesMap[0]) || newPlace.Y < 0 ||
					newPlace.Y >= len(pipesMap) {
					break
				}

				symbol := pipesMap[newPlace.Y][newPlace.X]
				_, isPipePath := pathMap[makeDirKey(newPlace)]
				if isPipePath && symbol != 'F' && symbol != 'J' {
					crosses++
				}
				currentPlace = newPlace
			}

			if crosses%2 != 0 {
				tilesInside++
			}
		}
	}

	return tilesInside
}

func main() {
	start := time.Now()
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	txt := string(b)

	lines := strings.Split(txt, "\n")
	var pipesMap [][]rune

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var m []rune
		for _, char := range line {
			m = append(m, char)
		}
		pipesMap = append(pipesMap, m)
	}

	answer := solutionPartOne(pipesMap)
	fmt.Println("Answer PART1:", answer)
	answer = solutionPartTwo(pipesMap)
	fmt.Println("Answer PART2:", answer)
	fmt.Println(time.Since(start))
}
