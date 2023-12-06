package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	b, err := os.ReadFile("in2.txt")
	checkErr(err)

	txt := string(b)
	sum := 0
	sumGearRatios := 0
	rowLen := len([]rune(strings.Split(txt, "\n")[0]))
	directions := []int{
		-rowLen - 1, // topleft
		-rowLen,     // topCenter
		-rowLen + 1, // topCenter
		-1,          // left
		1,           // right
		rowLen - 1,  // bottomLeft
		rowLen,      // bottomCenter
		rowLen + 1,  // bottomRight
	}

	txt = strings.ReplaceAll(txt, "\n", "")
	runeTxt := []rune(txt)

	for i, c := range runeTxt {
		if unicode.IsDigit(c) || c == '.' {
			continue
		}

		found := map[string]bool{}
		// we hit a symbol
		for _, dir := range directions {
			searchStartIdx := i + dir
			if searchStartIdx < 0 || searchStartIdx >= len(txt) {
				continue
			}

			if runeTxt[searchStartIdx] == '.' {
				continue
			}

			if !unicode.IsDigit(runeTxt[searchStartIdx]) {
				fmt.Println("ERROR", string(runeTxt[searchStartIdx]))
				panic("something went terribly wrong")
			}

			// search left
			numStartIdx := searchStartIdx
			for numStartIdx != 0 {
				if numStartIdx%rowLen == 0 {
					break
				}
				newRune := runeTxt[numStartIdx-1]
				if !unicode.IsDigit(newRune) {
					break
				}
				numStartIdx--
			}

			// search right
			numEndIdx := searchStartIdx
			for numEndIdx < len(runeTxt) {
				if (numEndIdx+1)%rowLen == 0 {
					break
				}
				newRune := runeTxt[numEndIdx+1]
				if !unicode.IsDigit(newRune) {
					break
				}
				numEndIdx++
			}

			numberStr := string(runeTxt[numStartIdx : numEndIdx+1])
			if _, ok := found[numberStr]; ok {
				continue // if number is already added -> skip it| this is a hack
			}
			found[numberStr] = true
			number, err := strconv.Atoi(numberStr)
			checkErr(err)
			sum += number
		}

		if c == '*' && len(found) == 2 {
			ratio := 1
			for k := range found {
				number, err := strconv.Atoi(k)
				checkErr(err)
				ratio *= number
			}

			sumGearRatios += ratio
		}
	}

	fmt.Println(sum)
	fmt.Println(sumGearRatios)
}
