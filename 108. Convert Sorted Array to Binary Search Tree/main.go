// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
    
    if len(nums) == 0 {
        return nil
    }

    return constructBST(nums, 0, len(nums)-1)
}

func constructBST(nums []int, left int, right int) *TreeNode {
    if left > right {
        return nil
    }

    mid := (left + right) / 2
    root := &TreeNode{Val: nums[mid]}

    root.Left = constructBST(nums, left, mid-1)
    root.Right = constructBST(nums, mid+1, right)
    return root
}
