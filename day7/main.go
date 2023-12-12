package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var CardRank = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

type CamelHand struct {
	Hand string
	Bid  int
	Type int
}

func (c *CamelHand) IsBigger(other *CamelHand) bool {
	for i, char := range c.Hand {
		otherChar := rune(other.Hand[i])
		if CardRank[char] > CardRank[otherChar] {
			return true
		} else if CardRank[char] < CardRank[otherChar] {
			return false
		}
	}
	return false
}

const (
	HighCard int = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func getType(hand string) int {
	m := map[rune]int{}

	for _, c := range hand {
		m[c]++
	}

	values := []int{}
	for _, v := range m {
		values = append(values, v)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	handValue := HighCard
	for _, val := range values {
		switch val {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			handValue = ThreeOfAKind
		case 2:
			switch handValue {
			case ThreeOfAKind:
				handValue = FullHouse
			case OnePair:
				handValue = TwoPair
			case HighCard:
				handValue = OnePair
			default:
				panic(fmt.Sprintf("unexpected combination of cards: %v", hand))
			}
		}
	}

	return handValue
}

func solutionPartOne(cardsMap map[int][]CamelHand) int {
	answer := 0
	currentRank := 1
	for i := 0; i <= 6; i++ {
		rankCards := cardsMap[i]
		sort.Slice(rankCards, func(i, j int) bool {
			return !rankCards[i].IsBigger(&rankCards[j])
		})

		for _, card := range rankCards {
			fmt.Println(
				currentRank,
				card.Hand,
				card.Type,
				card.Bid,
				(currentRank * card.Bid),
			)
			answer += (currentRank * card.Bid)
			currentRank++
		}
	}
	return answer
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

	cardsMap := map[int][]CamelHand{}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		contents := strings.Split(line, " ")
		handStr := strings.TrimSpace(contents[0])
		bid, err := strconv.Atoi(contents[1])
		checkErr(err)

		cardType := getType(handStr)
		card := CamelHand{
			Hand: handStr,
			Bid:  bid,
			Type: cardType,
		}

		if _, found := cardsMap[cardType]; !found {
			cardsMap[cardType] = []CamelHand{card}
		} else {
			cardsMap[cardType] = append(cardsMap[cardType], card)
		}
	}

	answer := solutionPartOne(cardsMap)
	fmt.Println("Answer:", answer)
	fmt.Println(time.Since(start))
}
