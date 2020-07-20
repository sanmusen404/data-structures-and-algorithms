package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	fmt.Scanln(&n)

	arr := initArr(n);

	mp(arr)
	printArr(arr)
}

//打印数组
func printArr(arr []int){
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}
}

//数组初始化(乱序)
func initArr(n int)[]int{
	var arr []int
	for i:=0;i<n;i++{
		//随机生成0~10000000的数字
		arr = append(arr,rand.Intn(10000000))
	}

	return arr
}

//冒泡排序
func mp(arr []int){
	for i:=0;i<len(arr);i++{
		for j:=i;j<len(arr);j++{
			if(arr[i]>arr[j]){
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
}