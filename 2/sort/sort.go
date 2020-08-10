package sort

//冒泡排序
//时间复杂度：o(n^2)
//空间复杂度：-
//是否稳定：是
func BubbleSort(arr []int) {
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
func InsertionSort(arr []int) {
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
func SelectionSort(arr []int) {
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
func MergeSort(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}

	mid := l / 2
	arr1 := arr[0:mid]
	arr2 := arr[mid:]
	MergeSort(arr1)
	MergeSort(arr2)

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

//快速排序-递归
//时间复杂度：o(log(n))
//空间复杂度：-
//是否稳定：否
func QuickSortByRecursive(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	i := 0
	pivot := l - 1
	for j := 0; j < l-1; j++ {
		if arr[j] < arr[pivot] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[pivot] = arr[pivot], arr[i]

	arr1 := arr[:i]
	arr2 := arr[i+1:]

	QuickSortByRecursive(arr1)
	QuickSortByRecursive(arr2)
}

//快速排序-循环
//时间复杂度：o(log(n))
//空间复杂度：-
//是否稳定：否
func QuickSortByLoop(arr []int) {

}
