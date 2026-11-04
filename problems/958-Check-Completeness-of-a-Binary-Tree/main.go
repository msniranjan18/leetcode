package main

import "fmt"

func intptr(i int) *int {
	return &i
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildBinaryTree(nums []*int) *TreeNode {
	if len(nums) == 0 || nums[0] == nil {
		return nil
	}

	root := &TreeNode{
		Val: *nums[0],
	}

	que := []*TreeNode{root}

	i := 1 //start after root
	for i < len(nums) && len(que) > 0 {
		currNode := que[0]
		que = que[1:]
		if nums[i] != nil {
			currNode.Left = &TreeNode{Val: *nums[i]}
			que = append(que, currNode.Left)
			i++
		}

		if i < len(nums) && nums[i] != nil {
			currNode.Right = &TreeNode{Val: *nums[i]}
			que = append(que, currNode.Right)
		}
		i++

	}
	return root
}

// printTree recursively prints the tree in a sideways ASCII format
func printTree(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Printf("├── L: ")
	} else {
		fmt.Printf("└── R: ")
	}
	fmt.Println(node.Val)

	// Choose the appropriate prefix for child recursive calls
	childPrefix := prefix
	if isLeft {
		childPrefix += "│   "
	} else {
		childPrefix += "    "
	}

	printTree(node.Left, childPrefix, true)
	printTree(node.Right, childPrefix, false)
}

func main() {

	t1 := []*int{
		intptr(1), intptr(2), intptr(3), intptr(4), intptr(5), intptr(6), intptr(7), // [1,2,3,4,5,6,7]
	}

	t2 := []*int{
		intptr(1), intptr(2), intptr(3), intptr(4), intptr(5), nil, intptr(7), // [1,2,3,4,5,null,7]
	}

	root := buildBinaryTree(t1)

	printTree(root, "", true)

	fmt.Println(isCompleteTree(root))

	root = buildBinaryTree(t2)

	printTree(root, "", true)

	fmt.Println(isCompleteTree(root))
}

func isCompleteTree(root *TreeNode) bool {

	// BFS

	que := []*TreeNode{root}

	for len(que) > 0 {
		node := que[0] // extract left element
		que = que[1:]  // trim the que
		if node != nil {
			que = append(que, node.Left) // Add the children of the node if not nil
			que = append(que, node.Right)
		} else { // if the node is nil then all other node should also be nil to be a complete binary tree
			for len(que) > 0 {
				node := que[0]   // extract left element
				que = que[1:]    // trim the que
				if node != nil { // node should be nil to be  acomplete binary tree after nil node occur.
					return false
				}
			}
		}
	}
	return true
}
