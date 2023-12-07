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
	Number  int // Number of cards with this ID (includes copies)
}

func (c Card) Value() int {
	value := 0
	for _, num := range c.Have {
		if _, found := c.Winning[num]; found {
			value++
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
	card.Number = 1

	return card
}

func main() {
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	txt := string(b)
	sum := 0

	cards := []Card{}
	for _, line := range strings.Split(txt, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		card := processLine(line)
		cards = append(cards, card)
	}

	for cardIdx, card := range cards {
		numCopiesWon := card.Value()
		for j := 1; j <= card.Number; j++ {
			for i := cardIdx + 1; i <= cardIdx+numCopiesWon; i++ {
				cards[i].Number++
			}
		}
	}

	for _, card := range cards {
		// fmt.Println(card.Id, card.Number, card.Value())
		sum += card.Number
	}

	fmt.Println(sum)
}
