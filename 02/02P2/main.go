package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []int
		for _, v := range strings.Split(scanner.Text(), " ") {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			line = append(line, i)
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	var count int

	for _, v := range lines {
		if isAscending(v) || isDescending(v) {
			count++
		}
	}

	fmt.Println(count)
}

func isAscending(numbers []int) bool {
	return checkOrder(numbers, true)
}

func isDescending(numbers []int) bool {
	return checkOrder(numbers, false)
}

func checkOrder(numbers []int, ascending bool) bool {
	if isValidSequence(numbers, ascending) {
		fmt.Printf("%v: SAFE (no removal needed)\n", numbers)
		return true
	}

	for i := 0; i < len(numbers); i++ {
		tempNumbers := make([]int, len(numbers))
		copy(tempNumbers, numbers)
		tempNumbers = append(tempNumbers[:i], tempNumbers[i+1:]...)
		if isValidSequence(tempNumbers, ascending) {
			fmt.Printf("%v: SAFE (removed %d at position %d)\n", numbers, numbers[i], i)
			return true
		}
	}

	fmt.Printf("%v: UNSAFE\n", numbers)
	return false
}

func isValidSequence(numbers []int, ascending bool) bool {
	for i := 1; i < len(numbers); i++ {
		// Check order
		if ascending && numbers[i] <= numbers[i-1] {
			return false
		}
		if !ascending && numbers[i] >= numbers[i-1] {
			return false
		}
		// Check jump size
		if absInt(numbers[i]-numbers[i-1]) > 3 {
			return false
		}
	}
	return true
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
