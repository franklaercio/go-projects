package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	figure "github.com/common-nighthawk/go-figure"
)

func main() {
		figure.NewColorFigure("Mega da Virada", "", "green", true).Print()

		var sortedNumbers = "21,24,33,41,48,56"

		numbers, _ := convertToIntArray(sortedNumbers)

		file, err := os.Open("bet.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
				numbersFile, _ := convertToIntArray(scanner.Text())
				diff := difference(numbers, numbersFile)
        log.Printf("From the game %v were sorted %d numbers - %s", numbersFile, len(diff), checkGame(len(diff)))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func difference(a, b []int) []int {
    m := make(map[int]bool)

    for _, item := range b {
        m[item] = true
    }

    var diff []int
    for _, item := range a {
        if _, found := m[item]; found {
            diff = append(diff, item)
        }
    }

    return diff
}

func convertToIntArray(s string) ([]int, error) {
	var numbers []int
	split := strings.Split(s, ",")
	for _, value := range split {
			num, err := strconv.Atoi(value)
			if err != nil {
					return nil, err
			}
			numbers = append(numbers, num)
	}
	return numbers, nil
}

func checkGame(number int) string {
	switch number {
	case 4:
		return "You won a block!"
	case 5:
		return "Nice try, you won a corner!"
	case 6:
		return "You are a new BILLIONAIRE!!!!!!!"
	default:
		return "Not this time"
	}
}