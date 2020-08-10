package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"./sort"
)

func main() {
	var n int
	fmt.Scanln(&n)

	arr := initArr(n)
	// arr := []int{3, 1, 7, 5, 8, 9, 6, 2, 1}

	wg := sync.WaitGroup{}
	//冒泡
	wg.Add(1)
	go func() {
		arr1 := make([]int, len(arr))
		copy(arr1, arr)
		t1 := time.Now()
		sort.BubbleSort(arr1)
		fmt.Println("冒泡排序:t = ", time.Since(t1))
		wg.Done()
	}()
	//插入
	wg.Add(1)
	go func() {
		arr2 := make([]int, len(arr))
		copy(arr2, arr)
		t2 := time.Now()
		sort.InsertionSort(arr2)
		fmt.Println("插入排序:t = ", time.Since(t2))
		wg.Done()
	}()
	//选择
	wg.Add(1)
	go func() {
		arr3 := make([]int, len(arr))
		copy(arr3, arr)
		t3 := time.Now()
		sort.SelectionSort(arr3)
		fmt.Println("选择排序:t = ", time.Since(t3))
		wg.Done()
	}()
	//归并
	wg.Add(1)
	go func() {
		arr4 := make([]int, len(arr))
		copy(arr4, arr)
		t4 := time.Now()
		sort.MergeSort(arr4)
		fmt.Println("归并排序:t = ", time.Since(t4))
		wg.Done()
	}()

	//快排
	wg.Add(1)
	go func() {
		arr5 := make([]int, len(arr))
		copy(arr5, arr)
		t5 := time.Now()
		sort.QuickSortByRecursive(arr5)
		fmt.Println("快速排序:t = ", time.Since(t5))
		wg.Done()
	}()
	wg.Wait()
}

//数组初始化(乱序)
func initArr(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		//随机生成0~10000000的数字
		arr = append(arr, rand.Intn(10000000))
	}

	return arr
}
