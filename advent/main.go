package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ROUND_WIN = 6
const ROUND_DRAW = 3
const ROCK = "A"
const PAPER = "B"
const SCISSORS = "C"
const FAKE_WIN = "Z"
const FAKE_DRAW = "Y"

func readNumber() int {
	i := 0

	for ok := true; ok; ok = (i < 1 || i > 24) {
		fmt.Println("-- Select day result (1-24) --")
		fmt.Print(">")
		var w1 string
		fmt.Scanln(&w1)

		var err error
		i, err = strconv.Atoi(w1)
		if err != nil {
			fmt.Println("-- Input is not a number --")
		} else if i < 1 || i > 24 {
			fmt.Println("-- Number not in range --")
		}
	}

	return i
}

func open(n string) *bufio.Scanner {

	f, _ := os.Open("../files/" + n + ".txt")
	return bufio.NewScanner(f)
}

func day1() {

	var total, acc, mElf int
	var top3 [3]int
	var top3Elves [3]int

	scanner := open("1")
	for elf := 1; scanner.Scan(); {
		value := scanner.Text()
		if value != "" {
			i, _ := strconv.Atoi(value)
			acc += i
		} else {
			if acc > total {
				total = acc
				mElf = elf
			}

			if acc > top3[0] {
				top3[2] = top3[1]
				top3[1] = top3[0]
				top3[0] = acc
			} else if acc > top3[1] {
				top3[2] = top3[1]
				top3[1] = acc
			} else if acc > top3[2] {
				top3[2] = acc
			}

			acc = 0
			elf++
		}
	}

	fmt.Printf("The elf %d is carrying the most calories with a total of %d\n", mElf, total)

	sum := 0
	for _, v := range top3 {
		sum += v
	}

	fmt.Printf("Top 3 elves carrying calories are %d(%d), %d(%d) and %d(%d), carrying a total of %d\n", top3Elves[0], top3[0], top3Elves[1], top3[1], top3Elves[2], top3[2], sum)

}

func rock(hand string) int {
	if hand == ROCK {
		return ROUND_DRAW
	}
	if hand == SCISSORS {
		return ROUND_WIN
	}
	return 0
}

func paper(hand string) int {
	if hand == PAPER {
		return ROUND_DRAW
	}
	if hand == ROCK {
		return ROUND_WIN
	}
	return 0
}

func scissors(hand string) int {
	if hand == SCISSORS {
		return ROUND_DRAW
	}
	if hand == PAPER {
		return ROUND_WIN
	}
	return 0
}

func roundScore(input string) int {
	split := strings.Split(input, " ")
	hand := split[0]

	var score int
	switch split[1] {
	case "X":
		score += rock(hand) + 1
	case "Y":
		score += paper(hand) + 2
	case "Z":
		score += scissors(hand) + 3
	}

	return score
}

func fakeRock(result string) int {
	if result == FAKE_DRAW {
		return ROUND_DRAW + 1
	}
	if result == FAKE_WIN {
		return ROUND_WIN + 2
	}

	return 3
}

func fakePaper(result string) int {
	if result == FAKE_DRAW {
		return ROUND_DRAW + 2
	}
	if result == FAKE_WIN {
		return ROUND_WIN + 3
	}

	return 1
}

func fakeScissors(result string) int {
	if result == FAKE_DRAW {
		return ROUND_DRAW + 3
	}
	if result == FAKE_WIN {
		return ROUND_WIN + 1
	}

	return 2
}

func fakeRoundScore(input string) int {
	split := strings.Split(input, " ")
	hand, result := split[0], split[1]

	switch hand {
	case ROCK:
		return fakeRock(result)
	case PAPER:
		return fakePaper(result)
	case SCISSORS:
		return fakeScissors(result)
	}

	return 0
}

func day2() {
	scanner := open("2")

	var realScore int
	var score int
	for scanner.Scan() {
		line := scanner.Text()
		score += roundScore(line)
		realScore += fakeRoundScore(line)
	}

	fmt.Printf("The figured total score is %d, but the real strategy score is %d", score, realScore)
}

func main() {
	n := readNumber()

	switch n {
	case 1:
		day1()
	case 2:
		day2()
	}
}
