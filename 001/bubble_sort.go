package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 5, 6, 3, 21, 43, 5, 6, 7, 4, 3, 0, 0}
	fmt.Println(bubbleSort(a, true))
	fmt.Println(bubbleSort(a, false))
}

//冒泡 O(N2)
func bubbleSort(arr []int, sort bool) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-1-i; j++ {
			if sort {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			} else {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}

	return arr
}
