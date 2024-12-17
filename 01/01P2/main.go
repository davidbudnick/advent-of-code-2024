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

	var listOne []int
	var listTwo []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")

		value1, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}

		value2, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}

		listOne = append(listOne, value1)
		listTwo = append(listTwo, value2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	var countMap = make(map[int]int)
	for i := 0; i < len(listTwo); i++ {
		countMap[listTwo[i]] = countMap[listTwo[i]] + 1
	}

	var totalCount int
	for _, v := range listOne {
		totalCount += (v * countMap[v])
	}

	fmt.Println(totalCount)
}
