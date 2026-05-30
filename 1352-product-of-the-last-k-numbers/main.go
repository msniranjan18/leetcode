/*
Design an algorithm that accepts a stream of integers and retrieves the product of the last k integers of the stream.

Implement the ProductOfNumbers class:

ProductOfNumbers() Initializes the object with an empty stream.
void add(int num) Appends the integer num to the stream.
int getProduct(int k) Returns the product of the last k numbers in the current list. You can assume that always the current list has at least k numbers.
The test cases are generated so that, at any time, the product of any contiguous sequence of numbers will fit into a single 32-bit integer without overflowing.

 

Example:

Input
["ProductOfNumbers","add","add","add","add","add","getProduct","getProduct","getProduct","add","getProduct"]
[[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]

Output
[null,null,null,null,null,null,20,40,0,null,32]

Explanation
ProductOfNumbers productOfNumbers = new ProductOfNumbers();
productOfNumbers.add(3);        // [3]
productOfNumbers.add(0);        // [3,0]
productOfNumbers.add(2);        // [3,0,2]
productOfNumbers.add(5);        // [3,0,2,5]
productOfNumbers.add(4);        // [3,0,2,5,4]
productOfNumbers.getProduct(2); // return 20. The product of the last 2 numbers is 5 * 4 = 20
productOfNumbers.getProduct(3); // return 40. The product of the last 3 numbers is 2 * 5 * 4 = 40
productOfNumbers.getProduct(4); // return 0. The product of the last 4 numbers is 0 * 2 * 5 * 4 = 0
productOfNumbers.add(8);        // [3,0,2,5,4,8]
productOfNumbers.getProduct(2); // return 32. The product of the last 2 numbers is 4 * 8 = 32 
 

Constraints:

0 <= num <= 100
1 <= k <= 4 * 104
At most 4 * 104 calls will be made to add and getProduct.
The product of the stream at any point in time will fit in a 32-bit integer.
 

Follow-up: Can you implement both GetProduct and Add to work in O(1) time complexity instead of O(k) time complexity?
*/

/*
Step 1: Choosing a Data Structure
Since we are continuously adding numbers to the stream and need to retrieve the product of the last k elements, we need a dynamic structure. Here are a few possible choices:

Option 1: Slice (Dynamic Array)
We can use a slice ([]int) to store the sequence of numbers.
Operations:
add(num): Append num to the slice (O(1) amortized).
getProduct(k): Compute the product of the last k elements (O(k) worst-case).
✔️ Pros:

Simple to implement.
Direct indexing support.
❌ Cons:

getProduct(k) could be slow if k is large (up to 40,000 calls).
Option 2: Prefix Product Array
Instead of storing raw numbers, maintain prefix products.
Operations:
add(num): Store cumulative product in a list.
getProduct(k): Compute the product using division in O(1).
✔️ Pros:

Very fast queries (O(1) instead of O(k)).
Saves computation time.
❌ Cons:

Needs extra handling for zeros (since division fails).
Memory grows with the number of additions.
*/

/*
🔹 Prefix Product Approach
✅ Steps to Implement
Maintain a prefixProduct list

The first element is 1 (helps with multiplication).
If num == 0, clear prefixProduct and restart from 1.
Otherwise, multiply the last element in prefixProduct by num and append.
Modify GetProduct(k int)

If k > len(prefixProduct) - 1, return 0 (because a zero was encountered in that range).
Otherwise, use division:
product=prefixProduct[𝑁]/prefixProduct[𝑁−𝑘]
 
💡 Example Walkthrough
Given Input:
plaintext
Copy
Edit
Stream:  [3, 0, 2, 5, 4]
Prefix:  [1] → [1, 3] → [1] (reset) → [1, 2] → [1, 2, 10] → [1, 2, 10, 40]
Query Results
getProduct(2) → 40 / 10 = 20 ✅
getProduct(3) → 40 / 2 = 40 ✅
getProduct(4) → 0 (zero in range) ✅
add(8) → [1, 2, 10, 40, 320]
getProduct(2) → 320 / 40 = 32 ✅
🚀 Time Complexity
Add(num): O(1)
GetProduct(k): O(1)
This approach is significantly faster than the brute-force method (O(k) per query).
*/

type ProductOfNumbers struct {
    PrefixProduct []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        PrefixProduct: []int{1},
    }
}

func (this *ProductOfNumbers) Add(num int)  {
    if num == 0 { // reset the product if zero encountered
        this.PrefixProduct = []int{1} // resetting it to [1]
        return
    }
    this.PrefixProduct = append(this.PrefixProduct, this.PrefixProduct[len(this.PrefixProduct)-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int { //32ms 22.65MB    
    if k >= len(this.PrefixProduct) { // zero was encountered in k length so product is zero
        return 0
    }
    return this.PrefixProduct[len(this.PrefixProduct)-1]/this.PrefixProduct[len(this.PrefixProduct)-1- k]
}


// Simple slice approach O(k)
// 537ms 16.65MB

// type ProductOfNumbers struct {
//     Nums []int
// }

// func Constructor() ProductOfNumbers {
//     nums := make([]int,0)
//     return ProductOfNumbers{
//         Nums: nums,
//     }
// }

// func (this *ProductOfNumbers) Add(num int)  {
//     this.Nums = append(this.Nums, num)
// }

// func (this *ProductOfNumbers) GetProduct(k int) int { //537ms 16.65MB
//     product := 1
//     for i:=len(this.Nums)-1;i>=len(this.Nums)-k;i-- {
//         if this.Nums[i] == 0 {
//             return 0
//         }
//         product *= this.Nums[i]
//     }
//     return product
// }


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
