package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Scanln(&n)

	arr := initArr(n)
	wg := sync.WaitGroup{}
	//
	wg.Add(1)
	go func() {
		arr1 := make([]int, len(arr))
		copy(arr1, arr)
		t1 := time.Now()
		bubbleSort(arr1)
		fmt.Println("冒泡排序:t = ", time.Since(t1))
		wg.Done()
	}()
	//
	wg.Add(1)
	go func() {
		arr2 := make([]int, len(arr))
		copy(arr2, arr)
		t2 := time.Now()
		insertionSort(arr2)
		fmt.Println("插入排序:t = ", time.Since(t2))
		wg.Done()
	}()

	//
	wg.Add(1)
	go func() {
		arr3 := make([]int, len(arr))
		copy(arr3, arr)
		t3 := time.Now()
		selectionSort(arr3)
		fmt.Println("选择排序:t = ", time.Since(t3))
		wg.Done()
	}()

	// arr4 := arr
	// t4 := time.Now()
	// fmt.Println("归并排序:t = ", time.Since(t4))

	// arr5 := arr
	// t5 := time.Now()
	// fmt.Println("快速排序:t = ", time.Since(t5))

	wg.Wait()
}

//打印数组
func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
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

//冒泡排序
//时间复杂度：o(n^2)
//空间复杂度：-
//是否稳定：稳定
func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

//插入排序
//时间复杂度：o(n^2)
//空间复杂度：-
//是否稳定：稳定
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				copy(arr[j+1:i+1], arr[j:i])
				arr[j] = temp
			}
		}
	}
}

//选择排序
//时间复杂度：o(n^2)
//空间复杂度：-
//是否稳定：稳定
func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

//归并排序
func mergeSort(arr []int) {

}

//快速排序
func quickSort(arr []int) {

}
