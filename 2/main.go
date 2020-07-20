package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Scanln(&n)

	arr := initArr(n)

	arr1 := arr
	t1 := time.Now()
	bubbleSort(arr1)
	fmt.Println("冒泡排序:t = ", time.Since(t1))

	// arr2 := arr
	// t2 := time.Now()
	// fmt.Println("插入排序:t = ", time.Since(t2))

	// arr3 := arr
	// t3 := time.Now()
	// fmt.Println("选择排序:t = ", time.Since(t3))

	// arr4 := arr
	// t4 := time.Now()
	// fmt.Println("归并排序:t = ", time.Since(t4))

	// arr5 := arr
	// t5 := time.Now()
	// fmt.Println("快速排序:t = ", time.Since(t5))
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
func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

//选择排序
func selectionSort(arr []int) {

}

//归并排序
func mergeSort(arr []int) {

}

//快速排序
func quickSort(arr []int) {

}
