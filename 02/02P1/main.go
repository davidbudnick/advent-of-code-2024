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
		if isAscending(v) {
			count++
		}

		if isDescending(v) {
			count++
		}
	}

	fmt.Println(count)
}

func isAscending(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] == numbers[i-1] {
			return false
		}

		if absInt(numbers[i]-numbers[i-1]) > 3 {
			return false
		}

		if numbers[i] < numbers[i-1] {
			return false
		}
	}
	return true
}

func isDescending(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] == numbers[i-1] {
			return false
		}

		if absInt(numbers[i]-numbers[i-1]) > 3 {
			return false
		}

		if numbers[i] > numbers[i-1] {
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
