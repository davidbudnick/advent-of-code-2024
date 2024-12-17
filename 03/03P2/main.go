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

	var reDoDoNot = regexp.MustCompile(`(?m)do\(\)|don't\(\)`)

	type marker struct {
		position int
		enabled  bool
	}
	var markers []marker
	for _, matchIdx := range reDoDoNot.FindAllStringIndex(line, -1) {
		match := line[matchIdx[0]:matchIdx[1]]
		markers = append(markers, marker{
			position: matchIdx[0],
			enabled:  match == "do()",
		})
	}

	total := 0
	active := true

	for _, matchIdx := range re.FindAllStringIndex(line, -1) {
		currentState := active
		for _, m := range markers {
			if m.position < matchIdx[0] {
				currentState = m.enabled
				active = currentState
			}
		}

		if currentState {
			match := line[matchIdx[0]:matchIdx[1]]
			numbers := strings.Split(strings.Trim(match, "mul()"), ",")
			n1, _ := strconv.Atoi(numbers[0])
			n2, _ := strconv.Atoi(numbers[1])
			total += n1 * n2
		}
	}

	fmt.Println(total)
}
