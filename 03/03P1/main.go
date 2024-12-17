package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line += scanner.Text()
	}

	var re = regexp.MustCompile(`(?m)mul\(\d{1,3},\d{1,3}\)`)

	total := 0
	for _, match := range re.FindAllString(line, -1) {
		numbers := strings.Split(
			strings.Trim(match, "mul()"),
			",",
		)

		n1, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		n2, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		total += n1 * n2
	}

	fmt.Println(total)
}
