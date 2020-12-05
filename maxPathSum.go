package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

const (

	INT_MAX = int(^uint((0)) >> 1)
	INT_MIN = ^INT_MAX

)
var maxSum int
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	maxSum = INT_MIN
	maxGain(root)
	return maxSum
}

func maxGain(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftGain := maxFuncForPathSum(0, maxGain(root.Left))
	rightGain := maxFuncForPathSum(0, maxGain(root.Right))
	retVal := root.Val + maxFuncForPathSum(leftGain, rightGain)

	newPathVal := root.Val + leftGain + rightGain
	maxSum = maxFuncForPathSum(maxSum, newPathVal)

	return retVal
}

func maxFuncForPathSum(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
