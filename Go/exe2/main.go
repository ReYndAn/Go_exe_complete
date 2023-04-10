package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("Before sorting:", arr)
	selectionSort(arr)
	fmt.Println("After sorting:", arr)
}
