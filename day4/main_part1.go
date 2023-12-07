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

type Card struct {
	Id      int
	Winning map[int]bool
	Have    []int
}

func (c Card) Value() int {
	value := 0
	for _, num := range c.Have {
		if _, found := c.Winning[num]; found {
			if value == 0 {
				value++
			} else {
				value *= 2
			}
		}
	}
	return value
}

func processLine(line string) Card {
	card := Card{}
	contents := strings.Split(line, ": ")

	cardIdStr := strings.ReplaceAll(contents[0], "Card ", "")
	cardId, err := strconv.Atoi(strings.TrimSpace(cardIdStr))
	checkErr(err)

	card.Id = cardId

	numberParts := strings.Split(contents[1], " | ")
	winningTxt := strings.TrimSpace(numberParts[0])
	winning := map[int]bool{}

	for _, num := range strings.Split(winningTxt, " ") {
		num = strings.TrimSpace(num)
		if num == "" {
			continue
		}
		numInt, err := strconv.Atoi(num)
		checkErr(err)
		winning[numInt] = true
	}
	card.Winning = winning

	haveTxt := strings.TrimSpace(numberParts[1])
	have := []int{}

	for _, num := range strings.Split(haveTxt, " ") {
		num = strings.TrimSpace(num)
		if num == "" {
			continue
		}
		numInt, err := strconv.Atoi(num)
		checkErr(err)
		have = append(have, numInt)
	}
	card.Have = have

	return card
}

func main() {
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	txt := string(b)
	sum := 0

	for _, line := range strings.Split(txt, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		card := processLine(line)
		sum += card.Value()
	}

	fmt.Println(sum)
}
