package LeetCode

import (
	"fmt"
	"testing"
)

//auxiliary
//************************************************************************************************************************
type TreeNode struct {
	Val    	 int
	Left 	 *TreeNode
	Right    *TreeNode
}

//************************************************************************************************************************


//1st//145.后续遍历
//************************************************************************************************************************
func Test_postorderTraversal(t *testing.T) {
	n3 := &TreeNode{
		Val:3,
		Left:nil,
		Right:nil,
	}
	n2 := &TreeNode{
		Val:2,
		Left:n3,
		Right:nil,
	}
	n1 := &TreeNode{
		Val:1,
		Right:n2,
	}
	fmt.Println(postorderTraversal(n1))
	fmt.Println(postorderTraversal_1(n1))
	fmt.Println(postorderTraversal_2(n1))
}

//递归
func postorderTraversal(root *TreeNode) (ans []int) {
	if root != nil {
		ans = append(ans, postorderTraversal(root.Left)...)
		ans = append(ans, postorderTraversal(root.Right)...)
		ans = append(ans, root.Val)
	} else {
		return nil
	}
	return
}
//栈
func postorderTraversal_1(root *TreeNode) (ans []int) {
	stack := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}
//Morris遍历  利用树中的空闲指针 在线性时间内只占用常数空间来实现后序遍历
func postorderTraversal_2(root *TreeNode) (ans []int) {
	addPath := func(node *TreeNode) {
		path := []int{}
		for ; node != nil ; node = node.Right {
			path = append(path, node.Val)
		}
		for i := len(path) - 1; i >= 0; i-- {
			ans = append(ans, path[i])
		}
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
}
//************************************************************************************************************************
