package main

import "fmt"

func getRow(rowIndex int) []int {
	var res []int
	var computeNum = cachedComputeNum(rowIndex + 1)
	for i := 0; i <= rowIndex; i++ {
		res = append(res, computeNum(rowIndex, i))
	}
	return res
}

func cachedComputeNum(numRows int) func(int, int) int {
	cache := make([][]int, numRows)
	for i := 0; i < numRows; i++  {
		cache[i] = make([]int, numRows)
	}

	return func (i, j int) int {
		return computeNum(i, j, cache)
	}
}

func computeNum(i, j int, cache [][]int) int {
	var n int
	if cache[i][j] == 0 {
		if j == 0 || i == j {
			n = 1
		} else {
			n = computeNum(i - 1, j - 1, cache) + computeNum(i - 1, j, cache)
		}
		cache[i][j] = n
	}
	return cache[i][j]
}

func main() {
	//cache := make([][]int, 5)
	//for i := 0; i < 5; i++  {
	//	cache[i] = make([]int, 5)
	//}
	//t(cache)
	//fmt.Println(cache)
	//a = append(a, []int{1})
	fmt.Println(getRow(3))
	//
	//n := 5
	//var a [][]int = make([][]int, n)
	//for i := 0; i < n; i++  {
	//	a[i] = make([]int, n)
	//}
	//a[0][0] = 1
	//fmt.Println(a)
}
