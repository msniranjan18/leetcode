// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	//postorder = postorder[:len(postorder)-1] // trim the last element
	root := &TreeNode{
		Val: rootVal,
	}
	mid := slices.Index(inorder, rootVal)
	root.Left = buildTree(inorder[:mid], postorder[:mid])
	root.Right = buildTree(inorder[mid+1:], postorder[mid:len(postorder)-1])
	return root
}
