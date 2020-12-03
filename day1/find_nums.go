package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	nums := getNumList("day1/nums.txt")

	a, b := findTwoNums(nums, 2020);

	fmt.Printf("Found %v, %v, which multiply to %v\n", a, b, a*b)

	a, b, c := findThreeNums(nums, 2020);

	fmt.Printf("Found %v, %v, %v which multiply to %v\n", a, b, c, a*b*c)

}

func getNumList(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := make([]int, 0, 100)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textNum := scanner.Text()

		if len(textNum) == 0 {
			continue
		}

		num, err := strconv.Atoi(textNum)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nums
}

func findTwoNums(nums []int, target int) (int, int) {
	numCounts := make(map[int]int)

	for _, num := range nums {
		numCounts[num]++
	}

	for _, num := range nums {
		if numCounts[target - num] > 0 {
			return num, target - num
		}
	}

	return 0, 0
}

func findThreeNums(nums []int, target int) (int, int, int) {
	numIdx := make(map[int]int)

	for i, num := range nums {
		numIdx[num]  = i + 1
	}

	for i, num1 := range nums {
		for j, num2 := range nums {
			if i == j {
				continue
			}

			num3 := target - num1 - num2

			num3Idx := numIdx[num3] - 1

			if num3Idx >= 0 && num3Idx != i && num3Idx != j {
				return num1, num2, num3
			}
		}
	}

	return 0, 0, 0
}
