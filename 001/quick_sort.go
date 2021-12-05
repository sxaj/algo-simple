package main

import "fmt"

func main() {
	a := []int{2, 3, -4, 5, -6, 3, -21, 43, 5, 6, 7, 4, 3, 0, 0}
	QuickSort(a)
	fmt.Println(a)
}

//快排 O(nlogn) 二分的冒泡 正序
func quickSort(arr []int, left int, right int) {
	if left > right {
		return
	}
	temp := arr[left] //基准
	i, j := left, right

	for i < j { //左右相向遍历 查找 基准在左边 从右边开始
		for i < j && arr[j] >= temp { //右边开始查找比基准小的 跳出循环
			j--
		}
		for i < j && arr[i] <= temp { //左边开始查找比基准大的 跳出循环
			i++
		}
		//交换 i j 找到的位置
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	} // i==j 则跳出循环
	//重置基准
	arr[left], arr[i] = arr[i], arr[left]

	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}
