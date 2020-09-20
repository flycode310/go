package main

var num, K, ansNum int
func kthSmallest(root *TreeNode, k int) int {
	num = 0
	K = k
	inorder(root)
	return ansNum
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	num ++
	if num == K {
		ansNum = root.Val
	}
	inorder(root.Right)
}
