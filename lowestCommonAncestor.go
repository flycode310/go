package main

var ans *TreeNode
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	recurseTree(root, p, q)
	return ans
}

func recurseTree(root, p, q *TreeNode) {
	if root == nil {
		return
	}
	if root.Val > p.Val && root.Val < q.Val {
		ans = root
		return
	}
	if root.Val == p.Val || root.Val == q.Val {
		ans = root
		return
	}
	if root.Val > q.Val {
		recurseTree(root.Left, p, q)
	}
	if root.Val < p.Val {
		recurseTree(root.Right, p, q)
	}
	return
}
