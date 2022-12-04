package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func day1() {
	f, _ := os.Open("../files/1.txt")
	scanner := bufio.NewScanner(f)

	var total, acc, mElf int
	var top3 [3]int
	var top3Elves [3]int

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

func main() {
	n := readNumber()

	switch n {
	case 1:
		day1()
	}
}
