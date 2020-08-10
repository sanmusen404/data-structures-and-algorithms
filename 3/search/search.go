package search

//二分查找
func BinarySearch(arr []int, need int) int {
	l := len(arr)
	low := 0
	high := l - 1
	mid := l / 2
	for low <= high {
		if arr[mid] > need { //左
			high = mid - 1
		} else { //右
			low = mid + 1
		}
		mid = (low + high) / 2
	}
	return mid
}
