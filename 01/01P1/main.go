package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	sort.Slice(listOne, func(i, j int) bool {
		return listOne[i] < listOne[j]
	})

	sort.Slice(listTwo, func(i, j int) bool {
		return listTwo[i] < listTwo[j]
	})

	var totalDistance int
	for i := 0; i < len(listOne); i++ {
		totalDistance += int(math.Abs(float64(listTwo[i] - listOne[i])))
	}

	fmt.Println(totalDistance)
}
