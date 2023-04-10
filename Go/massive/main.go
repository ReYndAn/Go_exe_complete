package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numRows := 5
	numCols := 4

	arr := make([][]int, numRows)
	for i := range arr {
		arr[i] = make([]int, numCols)
	}

	used := make(map[int]bool)
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			var num int
			for {
				num = rand.Intn(numRows*numCols) + 1
				if !used[num] {
					used[num] = true
					break
				}
			}
			arr[i][j] = num
		}
	}

	maxRow := 0
	maxSum := 0
	for i := 0; i < numRows; i++ {
		sum := 0
		for j := 0; j < numCols; j++ {
			sum += arr[i][j]
		}
		if sum > maxSum {
			maxSum = sum
			maxRow = i
		}
	}

	fmt.Println("Массив:")
	for _, row := range arr {
		fmt.Println(row)
	}
	fmt.Printf("Строка с самой большой суммой: %v (сумма = %v)\n", arr[maxRow], maxSum)
}
