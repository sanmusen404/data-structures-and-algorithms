package main

import (
	"fmt"
	"math"
	"time"

	"./search"
)

func main() {
	var n int
	fmt.Scanln(&n)

	arr := initArr()

	t := time.Now()
	index := search.BinarySearch(arr, n)
	fmt.Printf("要查找的数，下标为%d，共用时%s\n", index, time.Since(t))
}

func initArr() []int {
	arr := make([]int, 0)
	for i := 0; i < math.MaxInt8; i++ {
		arr = append(arr, i)
	}
	return arr
}
