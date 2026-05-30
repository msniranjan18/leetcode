/*
You are given an integer array cost of size n. You are currently at position n (at the end of the line) in a line of n + 1 people (numbered from 0 to n).

You wish to move forward in the line, but each person in front of you charges a specific amount to swap places. The cost to swap with person i is given by cost[i].

You are allowed to swap places with people as follows:

If they are in front of you, you must pay them cost[i] to swap with them.
If they are behind you, they can swap with you for free.
Return an array answer of size n, where answer[i] is the minimum total cost to reach each position i in the line.

 

Example 1:

Input: cost = [5,3,4,1,3,2]

Output: [5,3,3,1,1,1]

Explanation:

We can get to each position in the following way:

i = 0. We can swap with person 0 for a cost of 5.
i = 1. We can swap with person 1 for a cost of 3.
i = 2. We can swap with person 1 for a cost of 3, then swap with person 2 for free.
i = 3. We can swap with person 3 for a cost of 1.
i = 4. We can swap with person 3 for a cost of 1, then swap with person 4 for free.
i = 5. We can swap with person 3 for a cost of 1, then swap with person 5 for free.
Example 2:

Input: cost = [1,2,4,6,7]

Output: [1,1,1,1,1]

Explanation:

We can swap with person 0 for a cost of 1, then we will be able to reach any position i for free.

 

Constraints:

1 <= n == cost.length <= 100
1 <= cost[i] <= 100

*/

/*
Understanding the Problem with a Visual Representation
Let's visualize the line of people and how you swap positions.

Example 1
Input:

cost = [5, 3, 4, 1, 3, 2]
🔴 You start at position n = 6 (end of the line).


Position:  0    1    2    3    4    5    (You are at 6)
Cost:      5    3    4    1    3    2
You want to move forward to each position i with minimum cost.

Step-by-step Explanation
Step 1: Move to Position 0
You pay 5 rupees to swap with person at position 0.

Now, you are at position 0.


[🔵 0]  1    2    3    4    5   
       (You paid 5 rupees)
✅ Minimum cost to reach position 0: 5

Step 2: Move to Position 1
You pay 3 rupees to swap with person at position 1.

Now, you are at position 1.


  0  [🔵 1]  2    3    4    5   
       (You paid 3 rupees)
✅ Minimum cost to reach position 1: 3

Step 3: Move to Position 2
You can first swap with 1 (cost = 3)

Then swap with 2 for free (since 2 is behind 1).


  0    1  [🔵 2]  3    4    5   
          (Total cost = 3)
✅ Minimum cost to reach position 2: 3

Step 4: Move to Position 3
You pay 1 rupee to swap with person at position 3.

Now, you are at position 3.


  0    1    2  [🔵 3]  4    5   
           (You paid 1 rupee)
✅ Minimum cost to reach position 3: 1

Step 5: Move to Position 4
You swap with 3 (cost = 1)

Then swap with 4 for free (since 4 is behind 3).


  0    1    2    3  [🔵 4]  5   
              (Total cost = 1)
✅ Minimum cost to reach position 4: 1

Step 6: Move to Position 5
You swap with 3 (cost = 1)

Then swap with 5 for free (since 5 is behind 3).


  0    1    2    3    4  [🔵 5]  
                (Total cost = 1)
✅ Minimum cost to reach position 5: 1

Final Output

[5, 3, 3, 1, 1, 1]
Why Can You Swap for Free with Someone Behind You?
When you swap with someone in front, you pay them to take their position.
Once you're in that position, people behind you can swap forward for free (they take your spot while you go ahead).

It's like a "chain reaction" where you pay once to move forward, and from there, you can get ahead for free using previous swaps.

Key Takeaways
✅ The minimum cost to reach i is determined by the cheapest way to get there.
✅ Once you're in a cheap position, you can move further ahead for free.
✅ You only pay when moving forward into a new "expensive" position.
*/


func minCosts(cost []int) []int {
    minSwapAmount := 999 // since 1 <= cost[i] <= 100
    for idx, swapAmount := range cost {
        if swapAmount < minSwapAmount {
            minSwapAmount = swapAmount
        }
        cost[idx] = minSwapAmount
    }
    return cost
}
