func searchInsert(nums []int, target int) int {
   l := 0
   r := len(nums)-1
   mid := 0
   if target > nums[r] {
    return r+1
   }
   if target < nums[l] {
    return l
   }

   for l <= r {
    mid = ( l + r ) / 2
    //fmt.Println("mid:", mid, l,r)
    if target > nums[mid] {
        l = mid + 1
    } else {
        r = mid - 1
    }
   }
   if target > nums[mid] {
    return mid + 1
   }
   return mid
}
