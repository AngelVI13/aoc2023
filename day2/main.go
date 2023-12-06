package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID    int
	Red   int
	Green int
	Blue  int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func processLine(line string) Game {
	g := Game{}

	contents := strings.Split(line, ": ")
	id, err := strconv.Atoi(strings.ReplaceAll(contents[0], "Game ", ""))
	checkErr(err)

	g.ID = id

	for _, set := range strings.Split(contents[1], "; ") {
		colors := strings.Split(set, ", ")

		for _, color := range colors {
			if strings.Contains(color, "red") {
				colorNum, err := strconv.Atoi(strings.ReplaceAll(color, " red", ""))
				checkErr(err)
				if colorNum > g.Red {
					g.Red = colorNum
				}
			} else if strings.Contains(color, "green") {
				colorNum, err := strconv.Atoi(strings.ReplaceAll(color, " green", ""))
				checkErr(err)
				if colorNum > g.Green {
					g.Green = colorNum
				}
			} else if strings.Contains(color, "blue") {
				colorNum, err := strconv.Atoi(strings.ReplaceAll(color, " blue", ""))
				checkErr(err)
				if colorNum > g.Blue {
					g.Blue = colorNum
				}
			} else {
				panic(color)
			}
		}
	}

	return g
}

func main() {
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	maxGame := Game{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	txt := string(b)
	sum := 0
	sumPower := 0
	for _, line := range strings.Split(txt, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		game := processLine(line)
		if game.Red <= maxGame.Red && game.Green <= maxGame.Green &&
			game.Blue <= maxGame.Blue {
			sum += game.ID
		}

		sumPower += game.Red * game.Green * game.Blue
	}

	fmt.Println(sum)
	fmt.Println(sumPower)
}
