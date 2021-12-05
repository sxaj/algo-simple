package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 5, 6, 3, 21, 43, 5, 6, 7, 4, 3, 0, 0}
	fmt.Println(bucketSort(a, true))
	fmt.Println(bucketSort(a, false))
}

//只能大于0的自然数，因为数组索引从0开始  O(M+N)
func bucketSort(arr []int, sort bool) (ret []int) {
	var max int
	for _, v := range arr { //获取桶最大值
		if max < v {
			max = v
		}
	}
	bucket := make([]int, max+1) //初始化桶 全是0值
	for _, v := range arr {      //填充桶
		bucket[v]++
	}
	if sort { //正序
		for k, v := range bucket {
			for i := 0; i < v; i++ {
				ret = append(ret, k)
			}
		}
	} else { //倒序
		for bi := max; bi >= 0; bi-- {
			for i := 0; i < bucket[bi]; i++ {
				ret = append(ret, bi)
			}
		}
	}
	return
}
