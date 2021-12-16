package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func preOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.left == nil && root.right == nil {
		return []int{root.val}
	}
	var stack []*TreeNode
	var res []int
	stack = append(stack, root)
	for len(stack) != 0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, i.val)
		if i.right != nil {
			stack = append(stack, i.right)
		}
		if i.left != nil {
			stack = append(stack, i.left)
		}
	}
	return res
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.left == nil && root.right == nil {
		return []int{root.val}
	}
	res := inorder(root.left)
	res = append(res, root.val)
	res = append(res, inorder(root.right)...)
	return res
}

func postOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	if root.left != nil {
		lres := postOrder(root.left)
		if len(lres) > 0 {
			res = append(res, lres...)
		}
	}
	if root.right != nil {
		rres := postOrder(root.right)
		if len(rres) > 0 {
			res = append(res, rres...)
		}
	}
	res = append(res, root.val)
	return res
}
