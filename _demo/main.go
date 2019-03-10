package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {

	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	return helper(nums)
}

func helper(nums []int) []*TreeNode {
	if len(nums) == 0 {
		return []*TreeNode{
			nil,
		}
	}
	if len(nums) == 1 {
		return []*TreeNode{
			{
				Val: nums[0]+1,
				Left: nil,
				Right: nil,
			},
		}
	}
	res := []*TreeNode{}
	for idx, i := range nums {
		leftNums, rightNums := getSubArr(nums, idx)
		leftNodes := helper(leftNums)
		rightNodes := helper(rightNums)
		for _, leftRoot := range leftNodes {
			for _, rightRoot := range rightNodes {
				node := &TreeNode{
					Val: i+1,
					Left: leftRoot,
					Right: rightRoot,
				}
				res = append(res, node)
			}
		}
	}
	return res
}

func getSubArr(nums []int, i int) (leftNums, rightNums []int) {
	leftNums = nums[:i]
	rightNums = nums[i+1:]
	if len(leftNums) == 0 {
		leftNums = nil
	}
	if len(rightNums) == 0 {
		rightNums = nil
	}
	return
}

func main() {
	fmt.Printf("%v", generateTrees(0))
	//for _, row := range generateTrees(3) {
	//	fmt.Println(row)
	//}
	//a := []*int{}
	//b := make([]*int, 1)
	//b[0] = nil
	//fmt.Println(b)
	//a = append(a, b...)
	//fmt.Println(a)
}
