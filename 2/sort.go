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
	//冒泡
	wg.Add(1)
	go func() {
		arr1 := make([]int, len(arr))
		copy(arr1, arr)
		t1 := time.Now()
		bubbleSort(arr1)
		fmt.Println("冒泡排序:t = ", time.Since(t1))
		wg.Done()
	}()
	//插入
	wg.Add(1)
	go func() {
		arr2 := make([]int, len(arr))
		copy(arr2, arr)
		t2 := time.Now()
		insertionSort(arr2)
		fmt.Println("插入排序:t = ", time.Since(t2))
		wg.Done()
	}()
	//选择
	wg.Add(1)
	go func() {
		arr3 := make([]int, len(arr))
		copy(arr3, arr)
		t3 := time.Now()
		selectionSort(arr3)
		fmt.Println("选择排序:t = ", time.Since(t3))
		wg.Done()
	}()
	//归并
	wg.Add(1)
	go func() {
		arr4 := make([]int, len(arr))
		copy(arr4, arr)
		t4 := time.Now()
		mergeSort(arr4)
		fmt.Println("归并排序:t = ", time.Since(t4))
		wg.Done()
	}()

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
//是否稳定：是
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
//是否稳定：是
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
//是否稳定：否
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
//时间复杂度：o(log(n))
//空间复杂度：o(n)
//是否稳定：是
func mergeSort(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}

	mid := l / 2
	arr1 := arr[0:mid]
	arr2 := arr[mid:]
	mergeSort(arr1)
	mergeSort(arr2)

	tempArr := make([]int, 0)
	i, j := 0, 0
	for {
		if i >= len(arr1) {
			tempArr = append(tempArr, arr2[j:]...)
			break
		}
		if j >= len(arr2) {
			tempArr = append(tempArr, arr1[i:]...)
			break
		}
		if arr1[i] < arr2[j] {
			tempArr = append(tempArr, arr1[i])
			i++
		} else {
			tempArr = append(tempArr, arr2[j])
			j++
		}
	}
	copy(arr, tempArr)
}

//快速排序
//时间复杂度：o(log(n))
//空间复杂度：-
//是否稳定：否
func quickSort(arr []int) {
	// pivot := 0
}
